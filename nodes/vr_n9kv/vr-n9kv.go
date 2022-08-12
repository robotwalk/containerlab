// Copyright 2020 Nokia
// Licensed under the BSD 3-Clause License.
// SPDX-License-Identifier: BSD-3-Clause

package vr_n9kv

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/cloudflare/cfssl/log"
	"github.com/srl-labs/containerlab/nodes"
	"github.com/srl-labs/containerlab/runtime"
	"github.com/srl-labs/containerlab/types"
	"github.com/srl-labs/containerlab/utils"
)

var kindnames = []string{"vr-n9kv", "vr-cisco_n9kv"}

const (
	configDirName   = "config"
	startupCfgFName = "startup-config.cfg"

	defaultUser     = "admin"
	defaultPassword = "admin"
)

func init() {
	nodes.Register(kindnames, func() nodes.Node {
		return new(vrN9kv)
	})
	err := nodes.SetDefaultCredentials(kindnames, defaultUser, defaultPassword)
	if err != nil {
		log.Error(err)
	}
}

type vrN9kv struct {
	cfg     *types.NodeConfig
	mgmt    *types.MgmtNet
	runtime runtime.ContainerRuntime
}

func (s *vrN9kv) Init(cfg *types.NodeConfig, opts ...nodes.NodeOption) error {
	s.cfg = cfg
	for _, o := range opts {
		o(s)
	}
	// env vars are used to set launch.py arguments in vrnetlab container
	defEnv := map[string]string{
		"CONNECTION_MODE":    nodes.VrDefConnMode,
		"USERNAME":           defaultUser,
		"PASSWORD":           defaultPassword,
		"DOCKER_NET_V4_ADDR": s.mgmt.IPv4Subnet,
		"DOCKER_NET_V6_ADDR": s.mgmt.IPv6Subnet,
	}
	s.cfg.Env = utils.MergeStringMaps(defEnv, s.cfg.Env)

	// mount config dir to support startup-config functionality
	s.cfg.Binds = append(s.cfg.Binds, fmt.Sprint(path.Join(s.cfg.LabDir, configDirName), ":/config"))

	if s.cfg.Env["CONNECTION_MODE"] == "macvtap" {
		// mount dev dir to enable macvtap
		s.cfg.Binds = append(s.cfg.Binds, "/dev:/dev")
	}

	s.cfg.Cmd = fmt.Sprintf("--username %s --password %s --hostname %s --connection-mode %s --trace",
		s.cfg.Env["USERNAME"], s.cfg.Env["PASSWORD"], s.cfg.ShortName, s.cfg.Env["CONNECTION_MODE"])

	// set virtualization requirement
	s.cfg.HostRequirements.VirtRequired = true

	return nil
}
func (s *vrN9kv) Config() *types.NodeConfig { return s.cfg }
func (s *vrN9kv) PreDeploy(_, _, _ string) error {
	utils.CreateDirectory(s.cfg.LabDir, 0777)
	return loadStartupConfigFile(s.cfg)
}

func (s *vrN9kv) Deploy(ctx context.Context) error {
	cID, err := s.runtime.CreateContainer(ctx, s.cfg)
	if err != nil {
		return err
	}
	_, err = s.runtime.StartContainer(ctx, cID, s.cfg)
	return err
}

func (*vrN9kv) PostDeploy(_ context.Context, _ map[string]nodes.Node) error {
	return nil
}

func (s *vrN9kv) GetImages() map[string]string {
	return map[string]string{
		nodes.ImageKey: s.cfg.Image,
	}
}

func (*vrN9kv) Destroy(_ context.Context) error          { return nil }
func (s *vrN9kv) WithMgmtNet(mgmt *types.MgmtNet)        { s.mgmt = mgmt }
func (s *vrN9kv) WithRuntime(r runtime.ContainerRuntime) { s.runtime = r }
func (s *vrN9kv) GetRuntime() runtime.ContainerRuntime   { return s.runtime }

func (s *vrN9kv) Delete(ctx context.Context) error {
	return s.runtime.DeleteContainer(ctx, s.cfg.LongName)
}

func (*vrN9kv) SaveConfig(_ context.Context) error {
	return nil
}

func loadStartupConfigFile(node *types.NodeConfig) error {
	// create config directory that will be bind mounted to vrnetlab container at / path
	utils.CreateDirectory(path.Join(node.LabDir, configDirName), 0777)

	if node.StartupConfig != "" {
		// dstCfg is a path to a file on the clab host that will have rendered configuration
		dstCfg := filepath.Join(node.LabDir, configDirName, startupCfgFName)

		c, err := os.ReadFile(node.StartupConfig)
		if err != nil {
			return err
		}

		cfgTemplate := string(c)

		err = node.GenerateConfig(dstCfg, cfgTemplate)
		if err != nil {
			log.Errorf("node=%s, failed to generate config: %v", node.ShortName, err)
		}
	}
	return nil
}
