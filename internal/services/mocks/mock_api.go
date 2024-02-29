// Code generated by MockGen. DO NOT EDIT.
// Source: internal/services/api.go
//
// Generated by this command:
//
//	mockgen -source=internal/services/api.go -destination=internal/services/mocks/mock_api.go
//

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	http "net/http"
	services "projects/fb-server/internal/services"
	model "projects/fb-server/pkg/model"
	pgxs "projects/fb-server/pkg/pgxs"
	reflect "reflect"

	jwt "github.com/lestrrat-go/jwx/v2/jwt"
	gomock "go.uber.org/mock/gomock"
)

// MockApiService is a mock of ApiService interface.
type MockApiService struct {
	ctrl     *gomock.Controller
	recorder *MockApiServiceMockRecorder
}

// MockApiServiceMockRecorder is the mock recorder for MockApiService.
type MockApiServiceMockRecorder struct {
	mock *MockApiService
}

// NewMockApiService creates a new mock instance.
func NewMockApiService(ctrl *gomock.Controller) *MockApiService {
	mock := &MockApiService{ctrl: ctrl}
	mock.recorder = &MockApiServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApiService) EXPECT() *MockApiServiceMockRecorder {
	return m.recorder
}

// ApplyRoutes mocks base method.
func (m *MockApiService) ApplyRoutes() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ApplyRoutes")
}

// ApplyRoutes indicates an expected call of ApplyRoutes.
func (mr *MockApiServiceMockRecorder) ApplyRoutes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyRoutes", reflect.TypeOf((*MockApiService)(nil).ApplyRoutes))
}

// Init mocks base method.
func (m *MockApiService) Init(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockApiServiceMockRecorder) Init(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockApiService)(nil).Init), ctx)
}

// Shutdown mocks base method.
func (m *MockApiService) Shutdown(ctx context.Context, sig string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Shutdown", ctx, sig)
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockApiServiceMockRecorder) Shutdown(ctx, sig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockApiService)(nil).Shutdown), ctx, sig)
}

// MockAppInterface is a mock of AppInterface interface.
type MockAppInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAppInterfaceMockRecorder
}

// MockAppInterfaceMockRecorder is the mock recorder for MockAppInterface.
type MockAppInterfaceMockRecorder struct {
	mock *MockAppInterface
}

// NewMockAppInterface creates a new mock instance.
func NewMockAppInterface(ctrl *gomock.Controller) *MockAppInterface {
	mock := &MockAppInterface{ctrl: ctrl}
	mock.recorder = &MockAppInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppInterface) EXPECT() *MockAppInterfaceMockRecorder {
	return m.recorder
}

// AddService mocks base method.
func (m *MockAppInterface) AddService(name string, srv services.ApiService) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddService", name, srv)
}

// AddService indicates an expected call of AddService.
func (mr *MockAppInterfaceMockRecorder) AddService(name, srv any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddService", reflect.TypeOf((*MockAppInterface)(nil).AddService), name, srv)
}

// CheckIsAdmin mocks base method.
func (m *MockAppInterface) CheckIsAdmin(next http.HandlerFunc) http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIsAdmin", next)
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// CheckIsAdmin indicates an expected call of CheckIsAdmin.
func (mr *MockAppInterfaceMockRecorder) CheckIsAdmin(next any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIsAdmin", reflect.TypeOf((*MockAppInterface)(nil).CheckIsAdmin), next)
}

// GracefulShutdown mocks base method.
func (m *MockAppInterface) GracefulShutdown(ctx context.Context, sig string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GracefulShutdown", ctx, sig)
}

// GracefulShutdown indicates an expected call of GracefulShutdown.
func (mr *MockAppInterfaceMockRecorder) GracefulShutdown(ctx, sig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GracefulShutdown", reflect.TypeOf((*MockAppInterface)(nil).GracefulShutdown), ctx, sig)
}

// HandleEmailEvent mocks base method.
func (m *MockAppInterface) HandleEmailEvent(ctx context.Context, data *model.EmailData) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleEmailEvent", ctx, data)
}

// HandleEmailEvent indicates an expected call of HandleEmailEvent.
func (mr *MockAppInterfaceMockRecorder) HandleEmailEvent(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleEmailEvent", reflect.TypeOf((*MockAppInterface)(nil).HandleEmailEvent), ctx, data)
}

// HealthCheck mocks base method.
func (m *MockAppInterface) HealthCheck(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HealthCheck", w, r)
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockAppInterfaceMockRecorder) HealthCheck(w, r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockAppInterface)(nil).HealthCheck), w, r)
}

// IfLoggedIn mocks base method.
func (m *MockAppInterface) IfLoggedIn(fn http.HandlerFunc) http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IfLoggedIn", fn)
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// IfLoggedIn indicates an expected call of IfLoggedIn.
func (mr *MockAppInterfaceMockRecorder) IfLoggedIn(fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IfLoggedIn", reflect.TypeOf((*MockAppInterface)(nil).IfLoggedIn), fn)
}

// Init mocks base method.
func (m *MockAppInterface) Init(repo pgxs.FbRepo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", repo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockAppInterfaceMockRecorder) Init(repo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockAppInterface)(nil).Init), repo)
}

// Run mocks base method.
func (m *MockAppInterface) Run(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockAppInterfaceMockRecorder) Run(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockAppInterface)(nil).Run), ctx)
}

// RunHTTPServer mocks base method.
func (m *MockAppInterface) RunHTTPServer(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunHTTPServer", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunHTTPServer indicates an expected call of RunHTTPServer.
func (mr *MockAppInterfaceMockRecorder) RunHTTPServer(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunHTTPServer", reflect.TypeOf((*MockAppInterface)(nil).RunHTTPServer), ctx)
}

// ServeHTTP mocks base method.
func (m *MockAppInterface) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ServeHTTP", w, r)
}

// ServeHTTP indicates an expected call of ServeHTTP.
func (mr *MockAppInterfaceMockRecorder) ServeHTTP(w, r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServeHTTP", reflect.TypeOf((*MockAppInterface)(nil).ServeHTTP), w, r)
}

// loadJwtCerts mocks base method.
func (m *MockAppInterface) loadJwtCerts() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "loadJwtCerts")
	ret0, _ := ret[0].(error)
	return ret0
}

// loadJwtCerts indicates an expected call of loadJwtCerts.
func (mr *MockAppInterfaceMockRecorder) loadJwtCerts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "loadJwtCerts", reflect.TypeOf((*MockAppInterface)(nil).loadJwtCerts))
}

// verifyJWT mocks base method.
func (m *MockAppInterface) verifyJWT(jwtRawValue string) (jwt.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "verifyJWT", jwtRawValue)
	ret0, _ := ret[0].(jwt.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// verifyJWT indicates an expected call of verifyJWT.
func (mr *MockAppInterfaceMockRecorder) verifyJWT(jwtRawValue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "verifyJWT", reflect.TypeOf((*MockAppInterface)(nil).verifyJWT), jwtRawValue)
}
