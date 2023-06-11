// Code generated by MockGen. DO NOT EDIT.
// Source: back-end1-mini-project/repositories (interfaces: AdminRepositoryInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "back-end1-mini-project/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAdminRepositoryInterface is a mock of AdminRepositoryInterface interface.
type MockAdminRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAdminRepositoryInterfaceMockRecorder
}

// MockAdminRepositoryInterfaceMockRecorder is the mock recorder for MockAdminRepositoryInterface.
type MockAdminRepositoryInterfaceMockRecorder struct {
	mock *MockAdminRepositoryInterface
}

// NewMockAdminRepositoryInterface creates a new mock instance.
func NewMockAdminRepositoryInterface(ctrl *gomock.Controller) *MockAdminRepositoryInterface {
	mock := &MockAdminRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockAdminRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminRepositoryInterface) EXPECT() *MockAdminRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateApproval mocks base method.
func (m *MockAdminRepositoryInterface) CreateApproval(arg0 uint) (*entities.Approval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApproval", arg0)
	ret0, _ := ret[0].(*entities.Approval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateApproval indicates an expected call of CreateApproval.
func (mr *MockAdminRepositoryInterfaceMockRecorder) CreateApproval(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApproval", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).CreateApproval), arg0)
}

// CreateCustomer mocks base method.
func (m *MockAdminRepositoryInterface) CreateCustomer(arg0 *entities.Customer) (*entities.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCustomer", arg0)
	ret0, _ := ret[0].(*entities.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCustomer indicates an expected call of CreateCustomer.
func (mr *MockAdminRepositoryInterfaceMockRecorder) CreateCustomer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCustomer", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).CreateCustomer), arg0)
}

// DeleteCustomerById mocks base method.
func (m *MockAdminRepositoryInterface) DeleteCustomerById(arg0 uint, arg1 *entities.Customer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCustomerById", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCustomerById indicates an expected call of DeleteCustomerById.
func (mr *MockAdminRepositoryInterfaceMockRecorder) DeleteCustomerById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCustomerById", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).DeleteCustomerById), arg0, arg1)
}

// FetchCustomersFromAPI mocks base method.
func (m *MockAdminRepositoryInterface) FetchCustomersFromAPI() ([]*entities.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCustomersFromAPI")
	ret0, _ := ret[0].([]*entities.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCustomersFromAPI indicates an expected call of FetchCustomersFromAPI.
func (mr *MockAdminRepositoryInterfaceMockRecorder) FetchCustomersFromAPI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCustomersFromAPI", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).FetchCustomersFromAPI))
}

// GetAllCustomers mocks base method.
func (m *MockAdminRepositoryInterface) GetAllCustomers(arg0, arg1, arg2 string, arg3, arg4 int) ([]*entities.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCustomers", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*entities.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCustomers indicates an expected call of GetAllCustomers.
func (mr *MockAdminRepositoryInterfaceMockRecorder) GetAllCustomers(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCustomers", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).GetAllCustomers), arg0, arg1, arg2, arg3, arg4)
}

// GetApprovalById mocks base method.
func (m *MockAdminRepositoryInterface) GetApprovalById(arg0 uint) (*entities.Approval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApprovalById", arg0)
	ret0, _ := ret[0].(*entities.Approval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApprovalById indicates an expected call of GetApprovalById.
func (mr *MockAdminRepositoryInterfaceMockRecorder) GetApprovalById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApprovalById", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).GetApprovalById), arg0)
}

// GetCustomerByEmail mocks base method.
func (m *MockAdminRepositoryInterface) GetCustomerByEmail(arg0 string) (*entities.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomerByEmail", arg0)
	ret0, _ := ret[0].(*entities.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomerByEmail indicates an expected call of GetCustomerByEmail.
func (mr *MockAdminRepositoryInterfaceMockRecorder) GetCustomerByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomerByEmail", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).GetCustomerByEmail), arg0)
}

// GetCustomerById mocks base method.
func (m *MockAdminRepositoryInterface) GetCustomerById(arg0 uint) (*entities.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomerById", arg0)
	ret0, _ := ret[0].(*entities.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomerById indicates an expected call of GetCustomerById.
func (mr *MockAdminRepositoryInterfaceMockRecorder) GetCustomerById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomerById", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).GetCustomerById), arg0)
}

// LoginAdmin mocks base method.
func (m *MockAdminRepositoryInterface) LoginAdmin(arg0 string) (*entities.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginAdmin", arg0)
	ret0, _ := ret[0].(*entities.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginAdmin indicates an expected call of LoginAdmin.
func (mr *MockAdminRepositoryInterfaceMockRecorder) LoginAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginAdmin", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).LoginAdmin), arg0)
}

// RegisterAdmin mocks base method.
func (m *MockAdminRepositoryInterface) RegisterAdmin(arg0 *entities.Account) (*entities.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterAdmin", arg0)
	ret0, _ := ret[0].(*entities.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterAdmin indicates an expected call of RegisterAdmin.
func (mr *MockAdminRepositoryInterfaceMockRecorder) RegisterAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAdmin", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).RegisterAdmin), arg0)
}

// SaveCustomersFromAPI mocks base method.
func (m *MockAdminRepositoryInterface) SaveCustomersFromAPI(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCustomersFromAPI", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveCustomersFromAPI indicates an expected call of SaveCustomersFromAPI.
func (mr *MockAdminRepositoryInterfaceMockRecorder) SaveCustomersFromAPI(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCustomersFromAPI", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).SaveCustomersFromAPI), arg0)
}
