// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/filecoin-project/boost/storagemarket/types (interfaces: PieceAdder,CommpCalculator,DealPublisher,ChainDealManager,IndexProvider)

// Package mock_types is a generated GoMock package.
package mock_types

import (
	context "context"
	io "io"
	reflect "reflect"

	storagemarket "github.com/filecoin-project/boost-gfm/storagemarket"
	types "github.com/filecoin-project/boost/storagemarket/types"
	abi "github.com/filecoin-project/go-state-types/abi"
	market "github.com/filecoin-project/go-state-types/builtin/v9/market"
	api "github.com/filecoin-project/lotus/api"
	gomock "github.com/golang/mock/gomock"
	cid "github.com/ipfs/go-cid"
)

// MockPieceAdder is a mock of PieceAdder interface.
type MockPieceAdder struct {
	ctrl     *gomock.Controller
	recorder *MockPieceAdderMockRecorder
}

// MockPieceAdderMockRecorder is the mock recorder for MockPieceAdder.
type MockPieceAdderMockRecorder struct {
	mock *MockPieceAdder
}

// NewMockPieceAdder creates a new mock instance.
func NewMockPieceAdder(ctrl *gomock.Controller) *MockPieceAdder {
	mock := &MockPieceAdder{ctrl: ctrl}
	mock.recorder = &MockPieceAdderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPieceAdder) EXPECT() *MockPieceAdderMockRecorder {
	return m.recorder
}

// AddPiece mocks base method.
func (m *MockPieceAdder) AddPiece(arg0 context.Context, arg1 abi.UnpaddedPieceSize, arg2 io.Reader, arg3 api.PieceDealInfo) (abi.SectorNumber, abi.PaddedPieceSize, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPiece", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(abi.SectorNumber)
	ret1, _ := ret[1].(abi.PaddedPieceSize)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddPiece indicates an expected call of AddPiece.
func (mr *MockPieceAdderMockRecorder) AddPiece(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPiece", reflect.TypeOf((*MockPieceAdder)(nil).AddPiece), arg0, arg1, arg2, arg3)
}

// MockCommpCalculator is a mock of CommpCalculator interface.
type MockCommpCalculator struct {
	ctrl     *gomock.Controller
	recorder *MockCommpCalculatorMockRecorder
}

// MockCommpCalculatorMockRecorder is the mock recorder for MockCommpCalculator.
type MockCommpCalculatorMockRecorder struct {
	mock *MockCommpCalculator
}

// NewMockCommpCalculator creates a new mock instance.
func NewMockCommpCalculator(ctrl *gomock.Controller) *MockCommpCalculator {
	mock := &MockCommpCalculator{ctrl: ctrl}
	mock.recorder = &MockCommpCalculatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommpCalculator) EXPECT() *MockCommpCalculatorMockRecorder {
	return m.recorder
}

// ComputeDataCid mocks base method.
func (m *MockCommpCalculator) ComputeDataCid(arg0 context.Context, arg1 abi.UnpaddedPieceSize, arg2 io.Reader) (abi.PieceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeDataCid", arg0, arg1, arg2)
	ret0, _ := ret[0].(abi.PieceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComputeDataCid indicates an expected call of ComputeDataCid.
func (mr *MockCommpCalculatorMockRecorder) ComputeDataCid(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeDataCid", reflect.TypeOf((*MockCommpCalculator)(nil).ComputeDataCid), arg0, arg1, arg2)
}

// MockDealPublisher is a mock of DealPublisher interface.
type MockDealPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockDealPublisherMockRecorder
}

// MockDealPublisherMockRecorder is the mock recorder for MockDealPublisher.
type MockDealPublisherMockRecorder struct {
	mock *MockDealPublisher
}

// NewMockDealPublisher creates a new mock instance.
func NewMockDealPublisher(ctrl *gomock.Controller) *MockDealPublisher {
	mock := &MockDealPublisher{ctrl: ctrl}
	mock.recorder = &MockDealPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDealPublisher) EXPECT() *MockDealPublisherMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *MockDealPublisher) Publish(arg0 context.Context, arg1 market.ClientDealProposal) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", arg0, arg1)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Publish indicates an expected call of Publish.
func (mr *MockDealPublisherMockRecorder) Publish(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockDealPublisher)(nil).Publish), arg0, arg1)
}

// MockChainDealManager is a mock of ChainDealManager interface.
type MockChainDealManager struct {
	ctrl     *gomock.Controller
	recorder *MockChainDealManagerMockRecorder
}

// MockChainDealManagerMockRecorder is the mock recorder for MockChainDealManager.
type MockChainDealManagerMockRecorder struct {
	mock *MockChainDealManager
}

// NewMockChainDealManager creates a new mock instance.
func NewMockChainDealManager(ctrl *gomock.Controller) *MockChainDealManager {
	mock := &MockChainDealManager{ctrl: ctrl}
	mock.recorder = &MockChainDealManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainDealManager) EXPECT() *MockChainDealManagerMockRecorder {
	return m.recorder
}

// WaitForPublishDeals mocks base method.
func (m *MockChainDealManager) WaitForPublishDeals(arg0 context.Context, arg1 cid.Cid, arg2 market.DealProposal) (*storagemarket.PublishDealsWaitResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForPublishDeals", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storagemarket.PublishDealsWaitResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForPublishDeals indicates an expected call of WaitForPublishDeals.
func (mr *MockChainDealManagerMockRecorder) WaitForPublishDeals(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForPublishDeals", reflect.TypeOf((*MockChainDealManager)(nil).WaitForPublishDeals), arg0, arg1, arg2)
}

// MockIndexProvider is a mock of IndexProvider interface.
type MockIndexProvider struct {
	ctrl     *gomock.Controller
	recorder *MockIndexProviderMockRecorder
}

// MockIndexProviderMockRecorder is the mock recorder for MockIndexProvider.
type MockIndexProviderMockRecorder struct {
	mock *MockIndexProvider
}

// NewMockIndexProvider creates a new mock instance.
func NewMockIndexProvider(ctrl *gomock.Controller) *MockIndexProvider {
	mock := &MockIndexProvider{ctrl: ctrl}
	mock.recorder = &MockIndexProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIndexProvider) EXPECT() *MockIndexProviderMockRecorder {
	return m.recorder
}

// AnnounceBoostDeal mocks base method.
func (m *MockIndexProvider) AnnounceBoostDeal(arg0 context.Context, arg1 *types.ProviderDealState) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnnounceBoostDeal", arg0, arg1)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AnnounceBoostDeal indicates an expected call of AnnounceBoostDeal.
func (mr *MockIndexProviderMockRecorder) AnnounceBoostDeal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnounceBoostDeal", reflect.TypeOf((*MockIndexProvider)(nil).AnnounceBoostDeal), arg0, arg1)
}

// Enabled mocks base method.
func (m *MockIndexProvider) Enabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Enabled indicates an expected call of Enabled.
func (mr *MockIndexProviderMockRecorder) Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockIndexProvider)(nil).Enabled))
}

// Start mocks base method.
func (m *MockIndexProvider) Start(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", arg0)
}

// Start indicates an expected call of Start.
func (mr *MockIndexProviderMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockIndexProvider)(nil).Start), arg0)
}
