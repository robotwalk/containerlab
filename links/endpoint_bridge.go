package links

import (
	"context"
	"errors"
	"fmt"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"
)

type EndpointBridge struct {
	EndpointGeneric
	isMgmtBridgeEndpoint bool
}

func NewEndpointBridge(eg *EndpointGeneric, isMgmtBridgeEndpoint bool) *EndpointBridge {
	return &EndpointBridge{
		isMgmtBridgeEndpoint: isMgmtBridgeEndpoint,
		EndpointGeneric:      *eg,
	}
}

func (e *EndpointBridge) Verify(p *VerifyLinkParams) error {
	var errs []error
	err := CheckEndpointUniqueness(e)
	if err != nil {
		errs = append(errs, err)
	}
	if p.RunBridgeExistsCheck {
		err = CheckBridgeExists(e.GetNode())
		if err != nil {
			errs = append(errs, err)
		}
	}
	err = CheckEndpointDoesNotExistYet(e)
	if err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func (e *EndpointBridge) Deploy(ctx context.Context) error {
	return e.GetLink().Deploy(ctx, e)
}

func (e *EndpointBridge) IsNodeless() bool {
	// the mgmt bridge is nodeless.
	// If this is a regular bridge, then it should trigger BEnd deployment.
	return e.isMgmtBridgeEndpoint
}

// CheckBridgeExists verifies that the given bridge is present in the
// network namespace referenced via the provided nspath handle.
func CheckBridgeExists(n Node) error {
	return n.ExecFunction(func(_ ns.NetNS) error {
		br, err := netlink.LinkByName(n.GetShortName())
		_, notfound := err.(netlink.LinkNotFoundError)
		switch {
		case notfound:
			return fmt.Errorf("bridge %q referenced in topology but does not exist", n.GetShortName())
		case err != nil:
			return err
		case br.Type() != "bridge" && br.Type() != "openvswitch":
			return fmt.Errorf("interface %s found. expected type \"bridge\" or \"openvswitch\", actual is %q", n.GetShortName(), br.Type())
		}
		return nil
	})
}
