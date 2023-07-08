// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vinhluan/graphql (interfaces: GraphQL)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGraphQL is a mock of GraphQL interface.
type MockGraphQL struct {
	ctrl     *gomock.Controller
	recorder *MockGraphQLMockRecorder
}

// MockGraphQLMockRecorder is the mock recorder for MockGraphQL.
type MockGraphQLMockRecorder struct {
	mock *MockGraphQL
}

// NewMockGraphQL creates a new mock instance.
func NewMockGraphQL(ctrl *gomock.Controller) *MockGraphQL {
	mock := &MockGraphQL{ctrl: ctrl}
	mock.recorder = &MockGraphQLMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGraphQL) EXPECT() *MockGraphQLMockRecorder {
	return m.recorder
}

// Mutate mocks base method.
func (m *MockGraphQL) Mutate(arg0 context.Context, arg1 interface{}, arg2 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mutate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mutate indicates an expected call of Mutate.
func (mr *MockGraphQLMockRecorder) Mutate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mutate", reflect.TypeOf((*MockGraphQL)(nil).Mutate), arg0, arg1, arg2)
}

// MutateString mocks base method.
func (m *MockGraphQL) MutateString(arg0 context.Context, arg1 string, arg2 map[string]interface{}, arg3 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MutateString", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// MutateString indicates an expected call of MutateString.
func (mr *MockGraphQLMockRecorder) MutateString(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MutateString", reflect.TypeOf((*MockGraphQL)(nil).MutateString), arg0, arg1, arg2, arg3)
}

// Query mocks base method.
func (m *MockGraphQL) Query(arg0 context.Context, arg1 interface{}, arg2 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockGraphQLMockRecorder) Query(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockGraphQL)(nil).Query), arg0, arg1, arg2)
}

// QueryString mocks base method.
func (m *MockGraphQL) QueryString(arg0 context.Context, arg1 string, arg2 map[string]interface{}, arg3 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryString", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueryString indicates an expected call of QueryString.
func (mr *MockGraphQLMockRecorder) QueryString(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryString", reflect.TypeOf((*MockGraphQL)(nil).QueryString), arg0, arg1, arg2, arg3)
}
