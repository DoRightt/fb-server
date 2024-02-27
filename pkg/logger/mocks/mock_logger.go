// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/logger/logger.go
//
// Generated by this command:
//
//	mockgen -source=pkg/logger/logger.go -destination=pkg/logger/mocks/mock_logger.go
//

// Package mock_logger is a generated GoMock package.
package mock_logger

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	zap "go.uber.org/zap"
)

// MockFbLogger is a mock of FbLogger interface.
type MockFbLogger struct {
	ctrl     *gomock.Controller
	recorder *MockFbLoggerMockRecorder
}

// MockFbLoggerMockRecorder is the mock recorder for MockFbLogger.
type MockFbLoggerMockRecorder struct {
	mock *MockFbLogger
}

// NewMockFbLogger creates a new mock instance.
func NewMockFbLogger(ctrl *gomock.Controller) *MockFbLogger {
	mock := &MockFbLogger{ctrl: ctrl}
	mock.recorder = &MockFbLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFbLogger) EXPECT() *MockFbLoggerMockRecorder {
	return m.recorder
}

// Debug mocks base method.
func (m *MockFbLogger) Debug(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug.
func (mr *MockFbLoggerMockRecorder) Debug(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockFbLogger)(nil).Debug), args...)
}

// Debugf mocks base method.
func (m *MockFbLogger) Debugf(template string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{template}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockFbLoggerMockRecorder) Debugf(template any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{template}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockFbLogger)(nil).Debugf), varargs...)
}

// Debugw mocks base method.
func (m *MockFbLogger) Debugw(msg string, keysAndValues ...any) {
	m.ctrl.T.Helper()
	varargs := []any{msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugw", varargs...)
}

// Debugw indicates an expected call of Debugw.
func (mr *MockFbLoggerMockRecorder) Debugw(msg any, keysAndValues ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugw", reflect.TypeOf((*MockFbLogger)(nil).Debugw), varargs...)
}

// Error mocks base method.
func (m *MockFbLogger) Error(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockFbLoggerMockRecorder) Error(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockFbLogger)(nil).Error), args...)
}

// Errorf mocks base method.
func (m *MockFbLogger) Errorf(template string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{template}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockFbLoggerMockRecorder) Errorf(template any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{template}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockFbLogger)(nil).Errorf), varargs...)
}

// Fatal mocks base method.
func (m *MockFbLogger) Fatal(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatal", varargs...)
}

// Fatal indicates an expected call of Fatal.
func (mr *MockFbLoggerMockRecorder) Fatal(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockFbLogger)(nil).Fatal), args...)
}

// Fatalf mocks base method.
func (m *MockFbLogger) Fatalf(template string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{template}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalf", varargs...)
}

// Fatalf indicates an expected call of Fatalf.
func (mr *MockFbLoggerMockRecorder) Fatalf(template any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{template}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalf", reflect.TypeOf((*MockFbLogger)(nil).Fatalf), varargs...)
}

// Info mocks base method.
func (m *MockFbLogger) Info(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockFbLoggerMockRecorder) Info(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockFbLogger)(nil).Info), args...)
}

// Infof mocks base method.
func (m *MockFbLogger) Infof(template string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{template}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockFbLoggerMockRecorder) Infof(template any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{template}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockFbLogger)(nil).Infof), varargs...)
}

// Infow mocks base method.
func (m *MockFbLogger) Infow(msg string, keysAndValues ...any) {
	m.ctrl.T.Helper()
	varargs := []any{msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infow", varargs...)
}

// Infow indicates an expected call of Infow.
func (mr *MockFbLoggerMockRecorder) Infow(msg any, keysAndValues ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infow", reflect.TypeOf((*MockFbLogger)(nil).Infow), varargs...)
}

// Named mocks base method.
func (m *MockFbLogger) Named(name string) *zap.SugaredLogger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Named", name)
	ret0, _ := ret[0].(*zap.SugaredLogger)
	return ret0
}

// Named indicates an expected call of Named.
func (mr *MockFbLoggerMockRecorder) Named(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Named", reflect.TypeOf((*MockFbLogger)(nil).Named), name)
}

// Warn mocks base method.
func (m *MockFbLogger) Warn(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warn", varargs...)
}

// Warn indicates an expected call of Warn.
func (mr *MockFbLoggerMockRecorder) Warn(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockFbLogger)(nil).Warn), args...)
}

// Warnf mocks base method.
func (m *MockFbLogger) Warnf(template string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{template}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnf", varargs...)
}

// Warnf indicates an expected call of Warnf.
func (mr *MockFbLoggerMockRecorder) Warnf(template any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{template}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnf", reflect.TypeOf((*MockFbLogger)(nil).Warnf), varargs...)
}
