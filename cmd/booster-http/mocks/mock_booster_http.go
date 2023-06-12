// Code generated by MockGen. DO NOT EDIT.
// Source: server.go

// Package mocks_booster_http is a generated GoMock package.
package mocks_booster_http

import (
	context "context"
	reflect "reflect"

	model "github.com/filecoin-project/boostd-data/model"
	mount "github.com/filecoin-project/dagstore/mount"
	abi "github.com/filecoin-project/go-state-types/abi"
	gomock "github.com/golang/mock/gomock"
	cid "github.com/ipfs/go-cid"
)

// MockHttpServerApi is a mock of HttpServerApi interface.
type MockHttpServerApi struct {
	ctrl     *gomock.Controller
	recorder *MockHttpServerApiMockRecorder
}

// MockHttpServerApiMockRecorder is the mock recorder for MockHttpServerApi.
type MockHttpServerApiMockRecorder struct {
	mock *MockHttpServerApi
}

// NewMockHttpServerApi creates a new mock instance.
func NewMockHttpServerApi(ctrl *gomock.Controller) *MockHttpServerApi {
	mock := &MockHttpServerApi{ctrl: ctrl}
	mock.recorder = &MockHttpServerApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpServerApi) EXPECT() *MockHttpServerApiMockRecorder {
	return m.recorder
}

// GetPieceDeals mocks base method.
func (m *MockHttpServerApi) GetPieceDeals(ctx context.Context, pieceCID cid.Cid) ([]model.DealInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPieceDeals", ctx, pieceCID)
	ret0, _ := ret[0].([]model.DealInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPieceDeals indicates an expected call of GetPieceDeals.
func (mr *MockHttpServerApiMockRecorder) GetPieceDeals(ctx, pieceCID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPieceDeals", reflect.TypeOf((*MockHttpServerApi)(nil).GetPieceDeals), ctx, pieceCID)
}

// IsUnsealed mocks base method.
func (m *MockHttpServerApi) IsUnsealed(ctx context.Context, sectorID abi.SectorNumber, offset, length abi.UnpaddedPieceSize) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUnsealed", ctx, sectorID, offset, length)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUnsealed indicates an expected call of IsUnsealed.
func (mr *MockHttpServerApiMockRecorder) IsUnsealed(ctx, sectorID, offset, length interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUnsealed", reflect.TypeOf((*MockHttpServerApi)(nil).IsUnsealed), ctx, sectorID, offset, length)
}

// UnsealSectorAt mocks base method.
func (m *MockHttpServerApi) UnsealSectorAt(ctx context.Context, sectorID abi.SectorNumber, pieceOffset, length abi.UnpaddedPieceSize) (mount.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsealSectorAt", ctx, sectorID, pieceOffset, length)
	ret0, _ := ret[0].(mount.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnsealSectorAt indicates an expected call of UnsealSectorAt.
func (mr *MockHttpServerApiMockRecorder) UnsealSectorAt(ctx, sectorID, pieceOffset, length interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsealSectorAt", reflect.TypeOf((*MockHttpServerApi)(nil).UnsealSectorAt), ctx, sectorID, pieceOffset, length)
}
