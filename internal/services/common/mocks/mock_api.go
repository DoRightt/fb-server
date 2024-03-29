// Code generated by MockGen. DO NOT EDIT.
// Source: internal/services/common/api.go
//
// Generated by this command:
//
//	mockgen -source=internal/services/common/api.go -destination=internal/services/common/mocks/mock_api.go
//

// Package mock_common is a generated GoMock package.
package mock_common

import (
	context "context"
	http "net/http"
	reflect "reflect"

	pgx "github.com/jackc/pgx/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockCommonService is a mock of CommonService interface.
type MockCommonService struct {
	ctrl     *gomock.Controller
	recorder *MockCommonServiceMockRecorder
}

// MockCommonServiceMockRecorder is the mock recorder for MockCommonService.
type MockCommonServiceMockRecorder struct {
	mock *MockCommonService
}

// NewMockCommonService creates a new mock instance.
func NewMockCommonService(ctrl *gomock.Controller) *MockCommonService {
	mock := &MockCommonService{ctrl: ctrl}
	mock.recorder = &MockCommonServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommonService) EXPECT() *MockCommonServiceMockRecorder {
	return m.recorder
}

// ApplyRoutes mocks base method.
func (m *MockCommonService) ApplyRoutes() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ApplyRoutes")
}

// ApplyRoutes indicates an expected call of ApplyRoutes.
func (mr *MockCommonServiceMockRecorder) ApplyRoutes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyRoutes", reflect.TypeOf((*MockCommonService)(nil).ApplyRoutes))
}

// CheckEventIsDone mocks base method.
func (m *MockCommonService) CheckEventIsDone(ctx context.Context, tx pgx.Tx, fightId int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckEventIsDone", ctx, tx, fightId)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckEventIsDone indicates an expected call of CheckEventIsDone.
func (mr *MockCommonServiceMockRecorder) CheckEventIsDone(ctx, tx, fightId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckEventIsDone", reflect.TypeOf((*MockCommonService)(nil).CheckEventIsDone), ctx, tx, fightId)
}

// HandleNewEvent mocks base method.
func (m *MockCommonService) HandleNewEvent(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleNewEvent", w, r)
}

// HandleNewEvent indicates an expected call of HandleNewEvent.
func (mr *MockCommonServiceMockRecorder) HandleNewEvent(w, r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleNewEvent", reflect.TypeOf((*MockCommonService)(nil).HandleNewEvent), w, r)
}

// Init mocks base method.
func (m *MockCommonService) Init(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockCommonServiceMockRecorder) Init(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockCommonService)(nil).Init), ctx)
}

// Shutdown mocks base method.
func (m *MockCommonService) Shutdown(ctx context.Context, sig string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Shutdown", ctx, sig)
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockCommonServiceMockRecorder) Shutdown(ctx, sig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockCommonService)(nil).Shutdown), ctx, sig)
}
