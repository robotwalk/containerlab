// Code generated by MockGen. DO NOT EDIT.
// Source: types/exec.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/srl-labs/containerlab/types"
)

// MockExecExecutor is a mock of ExecExecutor interface.
type MockExecExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockExecExecutorMockRecorder
}

// MockExecExecutorMockRecorder is the mock recorder for MockExecExecutor.
type MockExecExecutorMockRecorder struct {
	mock *MockExecExecutor
}

// NewMockExecExecutor creates a new mock instance.
func NewMockExecExecutor(ctrl *gomock.Controller) *MockExecExecutor {
	mock := &MockExecExecutor{ctrl: ctrl}
	mock.recorder = &MockExecExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecExecutor) EXPECT() *MockExecExecutorMockRecorder {
	return m.recorder
}

// GetCmd mocks base method.
func (m *MockExecExecutor) GetCmd() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCmd")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetCmd indicates an expected call of GetCmd.
func (mr *MockExecExecutorMockRecorder) GetCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCmd", reflect.TypeOf((*MockExecExecutor)(nil).GetCmd))
}

// GetCmdString mocks base method.
func (m *MockExecExecutor) GetCmdString() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCmdString")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCmdString indicates an expected call of GetCmdString.
func (mr *MockExecExecutorMockRecorder) GetCmdString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCmdString", reflect.TypeOf((*MockExecExecutor)(nil).GetCmdString))
}

// SetReturnCode mocks base method.
func (m *MockExecExecutor) SetReturnCode(rc int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetReturnCode", rc)
}

// SetReturnCode indicates an expected call of SetReturnCode.
func (mr *MockExecExecutorMockRecorder) SetReturnCode(rc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReturnCode", reflect.TypeOf((*MockExecExecutor)(nil).SetReturnCode), rc)
}

// SetStdErr mocks base method.
func (m *MockExecExecutor) SetStdErr(stderr []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStdErr", stderr)
}

// SetStdErr indicates an expected call of SetStdErr.
func (mr *MockExecExecutorMockRecorder) SetStdErr(stderr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStdErr", reflect.TypeOf((*MockExecExecutor)(nil).SetStdErr), stderr)
}

// SetStdOut mocks base method.
func (m *MockExecExecutor) SetStdOut(stdout []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStdOut", stdout)
}

// SetStdOut indicates an expected call of SetStdOut.
func (mr *MockExecExecutorMockRecorder) SetStdOut(stdout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStdOut", reflect.TypeOf((*MockExecExecutor)(nil).SetStdOut), stdout)
}

// MockExecReader is a mock of ExecReader interface.
type MockExecReader struct {
	ctrl     *gomock.Controller
	recorder *MockExecReaderMockRecorder
}

// MockExecReaderMockRecorder is the mock recorder for MockExecReader.
type MockExecReaderMockRecorder struct {
	mock *MockExecReader
}

// NewMockExecReader creates a new mock instance.
func NewMockExecReader(ctrl *gomock.Controller) *MockExecReader {
	mock := &MockExecReader{ctrl: ctrl}
	mock.recorder = &MockExecReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecReader) EXPECT() *MockExecReaderMockRecorder {
	return m.recorder
}

// GetCmdString mocks base method.
func (m *MockExecReader) GetCmdString() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCmdString")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCmdString indicates an expected call of GetCmdString.
func (mr *MockExecReaderMockRecorder) GetCmdString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCmdString", reflect.TypeOf((*MockExecReader)(nil).GetCmdString))
}

// GetEntryInFormat mocks base method.
func (m *MockExecReader) GetEntryInFormat(format types.ExecOutputFormat) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntryInFormat", format)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntryInFormat indicates an expected call of GetEntryInFormat.
func (mr *MockExecReaderMockRecorder) GetEntryInFormat(format interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntryInFormat", reflect.TypeOf((*MockExecReader)(nil).GetEntryInFormat), format)
}

// GetReturnCode mocks base method.
func (m *MockExecReader) GetReturnCode() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReturnCode")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetReturnCode indicates an expected call of GetReturnCode.
func (mr *MockExecReaderMockRecorder) GetReturnCode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReturnCode", reflect.TypeOf((*MockExecReader)(nil).GetReturnCode))
}

// GetStdErrByteSlice mocks base method.
func (m *MockExecReader) GetStdErrByteSlice() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStdErrByteSlice")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetStdErrByteSlice indicates an expected call of GetStdErrByteSlice.
func (mr *MockExecReaderMockRecorder) GetStdErrByteSlice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStdErrByteSlice", reflect.TypeOf((*MockExecReader)(nil).GetStdErrByteSlice))
}

// GetStdErrString mocks base method.
func (m *MockExecReader) GetStdErrString() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStdErrString")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetStdErrString indicates an expected call of GetStdErrString.
func (mr *MockExecReaderMockRecorder) GetStdErrString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStdErrString", reflect.TypeOf((*MockExecReader)(nil).GetStdErrString))
}

// GetStdOutByteSlice mocks base method.
func (m *MockExecReader) GetStdOutByteSlice() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStdOutByteSlice")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetStdOutByteSlice indicates an expected call of GetStdOutByteSlice.
func (mr *MockExecReaderMockRecorder) GetStdOutByteSlice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStdOutByteSlice", reflect.TypeOf((*MockExecReader)(nil).GetStdOutByteSlice))
}

// GetStdOutString mocks base method.
func (m *MockExecReader) GetStdOutString() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStdOutString")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetStdOutString indicates an expected call of GetStdOutString.
func (mr *MockExecReaderMockRecorder) GetStdOutString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStdOutString", reflect.TypeOf((*MockExecReader)(nil).GetStdOutString))
}

// SetCmd mocks base method.
func (m *MockExecReader) SetCmd(s string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCmd", s)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCmd indicates an expected call of SetCmd.
func (mr *MockExecReaderMockRecorder) SetCmd(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCmd", reflect.TypeOf((*MockExecReader)(nil).SetCmd), s)
}

// String mocks base method.
func (m *MockExecReader) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockExecReaderMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockExecReader)(nil).String))
}
