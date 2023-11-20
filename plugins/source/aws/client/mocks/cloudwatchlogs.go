// Code generated by MockGen. DO NOT EDIT.
// Source: cloudwatchlogs.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	cloudwatchlogs "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	gomock "github.com/golang/mock/gomock"
)

// MockCloudwatchlogsClient is a mock of CloudwatchlogsClient interface.
type MockCloudwatchlogsClient struct {
	ctrl     *gomock.Controller
	recorder *MockCloudwatchlogsClientMockRecorder
}

// MockCloudwatchlogsClientMockRecorder is the mock recorder for MockCloudwatchlogsClient.
type MockCloudwatchlogsClientMockRecorder struct {
	mock *MockCloudwatchlogsClient
}

// NewMockCloudwatchlogsClient creates a new mock instance.
func NewMockCloudwatchlogsClient(ctrl *gomock.Controller) *MockCloudwatchlogsClient {
	mock := &MockCloudwatchlogsClient{ctrl: ctrl}
	mock.recorder = &MockCloudwatchlogsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudwatchlogsClient) EXPECT() *MockCloudwatchlogsClientMockRecorder {
	return m.recorder
}

// DescribeAccountPolicies mocks base method.
func (m *MockCloudwatchlogsClient) DescribeAccountPolicies(arg0 context.Context, arg1 *cloudwatchlogs.DescribeAccountPoliciesInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeAccountPoliciesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAccountPolicies")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAccountPolicies", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeAccountPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAccountPolicies indicates an expected call of DescribeAccountPolicies.
func (mr *MockCloudwatchlogsClientMockRecorder) DescribeAccountPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAccountPolicies", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeAccountPolicies), varargs...)
}

// DescribeDeliveries mocks base method.
func (m *MockCloudwatchlogsClient) DescribeDeliveries(arg0 context.Context, arg1 *cloudwatchlogs.DescribeDeliveriesInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeDeliveriesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeDeliveries")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDeliveries", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeDeliveriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDeliveries indicates an expected call of DescribeDeliveries.
func (mr *MockCloudwatchlogsClientMockRecorder) DescribeDeliveries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDeliveries", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeDeliveries), varargs...)
}

// DescribeDeliveryDestinations mocks base method.
func (m *MockCloudwatchlogsClient) DescribeDeliveryDestinations(arg0 context.Context, arg1 *cloudwatchlogs.DescribeDeliveryDestinationsInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeDeliveryDestinationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeDeliveryDestinations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDeliveryDestinations", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeDeliveryDestinationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDeliveryDestinations indicates an expected call of DescribeDeliveryDestinations.
func (mr *MockCloudwatchlogsClientMockRecorder) DescribeDeliveryDestinations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDeliveryDestinations", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeDeliveryDestinations), varargs...)
}

// DescribeDeliverySources mocks base method.
func (m *MockCloudwatchlogsClient) DescribeDeliverySources(arg0 context.Context, arg1 *cloudwatchlogs.DescribeDeliverySourcesInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeDeliverySourcesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeDeliverySources")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDeliverySources", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeDeliverySourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDeliverySources indicates an expected call of DescribeDeliverySources.
func (mr *MockCloudwatchlogsClientMockRecorder) DescribeDeliverySources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDeliverySources", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeDeliverySources), varargs...)
}

// DescribeDestinations mocks base method.
func (m *MockCloudwatchlogsClient) DescribeDestinations(arg0 context.Context, arg1 *cloudwatchlogs.DescribeDestinationsInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeDestinationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeDestinations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDestinations", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeDestinationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDestinations indicates an expected call of DescribeDestinations.
func (mr *MockCloudwatchlogsClientMockRecorder) DescribeDestinations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDestinations", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeDestinations), varargs...)
}

// DescribeExportTasks mocks base method.
func (m *MockCloudwatchlogsClient) DescribeExportTasks(arg0 context.Context, arg1 *cloudwatchlogs.DescribeExportTasksInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeExportTasksOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeExportTasks")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeExportTasks", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeExportTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeExportTasks indicates an expected call of DescribeExportTasks.
func (mr *MockCloudwatchlogsClientMockRecorder) DescribeExportTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeExportTasks", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeExportTasks), varargs...)
}

// DescribeLogGroups mocks base method.
func (m *MockCloudwatchlogsClient) DescribeLogGroups(arg0 context.Context, arg1 *cloudwatchlogs.DescribeLogGroupsInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeLogGroupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeLogGroups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLogGroups", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeLogGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLogGroups indicates an expected call of DescribeLogGroups.
func (mr *MockCloudwatchlogsClientMockRecorder) DescribeLogGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLogGroups", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeLogGroups), varargs...)
}

// DescribeLogStreams mocks base method.
func (m *MockCloudwatchlogsClient) DescribeLogStreams(arg0 context.Context, arg1 *cloudwatchlogs.DescribeLogStreamsInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeLogStreamsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeLogStreams")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLogStreams", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeLogStreamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLogStreams indicates an expected call of DescribeLogStreams.
func (mr *MockCloudwatchlogsClientMockRecorder) DescribeLogStreams(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLogStreams", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeLogStreams), varargs...)
}

// DescribeMetricFilters mocks base method.
func (m *MockCloudwatchlogsClient) DescribeMetricFilters(arg0 context.Context, arg1 *cloudwatchlogs.DescribeMetricFiltersInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeMetricFiltersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeMetricFilters")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeMetricFilters", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeMetricFiltersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeMetricFilters indicates an expected call of DescribeMetricFilters.
func (mr *MockCloudwatchlogsClientMockRecorder) DescribeMetricFilters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMetricFilters", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeMetricFilters), varargs...)
}

// DescribeQueries mocks base method.
func (m *MockCloudwatchlogsClient) DescribeQueries(arg0 context.Context, arg1 *cloudwatchlogs.DescribeQueriesInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeQueriesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeQueries")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeQueries", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeQueriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeQueries indicates an expected call of DescribeQueries.
func (mr *MockCloudwatchlogsClientMockRecorder) DescribeQueries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeQueries", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeQueries), varargs...)
}

// DescribeQueryDefinitions mocks base method.
func (m *MockCloudwatchlogsClient) DescribeQueryDefinitions(arg0 context.Context, arg1 *cloudwatchlogs.DescribeQueryDefinitionsInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeQueryDefinitionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeQueryDefinitions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeQueryDefinitions", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeQueryDefinitionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeQueryDefinitions indicates an expected call of DescribeQueryDefinitions.
func (mr *MockCloudwatchlogsClientMockRecorder) DescribeQueryDefinitions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeQueryDefinitions", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeQueryDefinitions), varargs...)
}

// DescribeResourcePolicies mocks base method.
func (m *MockCloudwatchlogsClient) DescribeResourcePolicies(arg0 context.Context, arg1 *cloudwatchlogs.DescribeResourcePoliciesInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeResourcePoliciesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeResourcePolicies")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeResourcePolicies", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeResourcePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeResourcePolicies indicates an expected call of DescribeResourcePolicies.
func (mr *MockCloudwatchlogsClientMockRecorder) DescribeResourcePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeResourcePolicies", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeResourcePolicies), varargs...)
}

// DescribeSubscriptionFilters mocks base method.
func (m *MockCloudwatchlogsClient) DescribeSubscriptionFilters(arg0 context.Context, arg1 *cloudwatchlogs.DescribeSubscriptionFiltersInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeSubscriptionFiltersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeSubscriptionFilters")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSubscriptionFilters", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeSubscriptionFiltersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSubscriptionFilters indicates an expected call of DescribeSubscriptionFilters.
func (mr *MockCloudwatchlogsClientMockRecorder) DescribeSubscriptionFilters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSubscriptionFilters", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).DescribeSubscriptionFilters), varargs...)
}

// GetDataProtectionPolicy mocks base method.
func (m *MockCloudwatchlogsClient) GetDataProtectionPolicy(arg0 context.Context, arg1 *cloudwatchlogs.GetDataProtectionPolicyInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetDataProtectionPolicyOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetDataProtectionPolicy")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDataProtectionPolicy", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.GetDataProtectionPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDataProtectionPolicy indicates an expected call of GetDataProtectionPolicy.
func (mr *MockCloudwatchlogsClientMockRecorder) GetDataProtectionPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataProtectionPolicy", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).GetDataProtectionPolicy), varargs...)
}

// GetDelivery mocks base method.
func (m *MockCloudwatchlogsClient) GetDelivery(arg0 context.Context, arg1 *cloudwatchlogs.GetDeliveryInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetDeliveryOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetDelivery")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDelivery", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.GetDeliveryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDelivery indicates an expected call of GetDelivery.
func (mr *MockCloudwatchlogsClientMockRecorder) GetDelivery(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDelivery", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).GetDelivery), varargs...)
}

// GetDeliveryDestination mocks base method.
func (m *MockCloudwatchlogsClient) GetDeliveryDestination(arg0 context.Context, arg1 *cloudwatchlogs.GetDeliveryDestinationInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetDeliveryDestinationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetDeliveryDestination")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDeliveryDestination", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.GetDeliveryDestinationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeliveryDestination indicates an expected call of GetDeliveryDestination.
func (mr *MockCloudwatchlogsClientMockRecorder) GetDeliveryDestination(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeliveryDestination", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).GetDeliveryDestination), varargs...)
}

// GetDeliveryDestinationPolicy mocks base method.
func (m *MockCloudwatchlogsClient) GetDeliveryDestinationPolicy(arg0 context.Context, arg1 *cloudwatchlogs.GetDeliveryDestinationPolicyInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetDeliveryDestinationPolicyOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetDeliveryDestinationPolicy")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDeliveryDestinationPolicy", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.GetDeliveryDestinationPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeliveryDestinationPolicy indicates an expected call of GetDeliveryDestinationPolicy.
func (mr *MockCloudwatchlogsClientMockRecorder) GetDeliveryDestinationPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeliveryDestinationPolicy", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).GetDeliveryDestinationPolicy), varargs...)
}

// GetDeliverySource mocks base method.
func (m *MockCloudwatchlogsClient) GetDeliverySource(arg0 context.Context, arg1 *cloudwatchlogs.GetDeliverySourceInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetDeliverySourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetDeliverySource")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDeliverySource", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.GetDeliverySourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeliverySource indicates an expected call of GetDeliverySource.
func (mr *MockCloudwatchlogsClientMockRecorder) GetDeliverySource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeliverySource", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).GetDeliverySource), varargs...)
}

// GetLogEvents mocks base method.
func (m *MockCloudwatchlogsClient) GetLogEvents(arg0 context.Context, arg1 *cloudwatchlogs.GetLogEventsInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetLogEventsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetLogEvents")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLogEvents", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.GetLogEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogEvents indicates an expected call of GetLogEvents.
func (mr *MockCloudwatchlogsClientMockRecorder) GetLogEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogEvents", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).GetLogEvents), varargs...)
}

// GetLogGroupFields mocks base method.
func (m *MockCloudwatchlogsClient) GetLogGroupFields(arg0 context.Context, arg1 *cloudwatchlogs.GetLogGroupFieldsInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetLogGroupFieldsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetLogGroupFields")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLogGroupFields", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.GetLogGroupFieldsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogGroupFields indicates an expected call of GetLogGroupFields.
func (mr *MockCloudwatchlogsClientMockRecorder) GetLogGroupFields(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogGroupFields", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).GetLogGroupFields), varargs...)
}

// GetLogRecord mocks base method.
func (m *MockCloudwatchlogsClient) GetLogRecord(arg0 context.Context, arg1 *cloudwatchlogs.GetLogRecordInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetLogRecordOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetLogRecord")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLogRecord", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.GetLogRecordOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogRecord indicates an expected call of GetLogRecord.
func (mr *MockCloudwatchlogsClientMockRecorder) GetLogRecord(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogRecord", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).GetLogRecord), varargs...)
}

// GetQueryResults mocks base method.
func (m *MockCloudwatchlogsClient) GetQueryResults(arg0 context.Context, arg1 *cloudwatchlogs.GetQueryResultsInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetQueryResultsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetQueryResults")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQueryResults", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.GetQueryResultsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueryResults indicates an expected call of GetQueryResults.
func (mr *MockCloudwatchlogsClientMockRecorder) GetQueryResults(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryResults", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).GetQueryResults), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockCloudwatchlogsClient) ListTagsForResource(arg0 context.Context, arg1 *cloudwatchlogs.ListTagsForResourceInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListTagsForResource")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockCloudwatchlogsClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).ListTagsForResource), varargs...)
}

// ListTagsLogGroup mocks base method.
func (m *MockCloudwatchlogsClient) ListTagsLogGroup(arg0 context.Context, arg1 *cloudwatchlogs.ListTagsLogGroupInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.ListTagsLogGroupOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudwatchlogs.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListTagsLogGroup")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsLogGroup", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.ListTagsLogGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsLogGroup indicates an expected call of ListTagsLogGroup.
func (mr *MockCloudwatchlogsClientMockRecorder) ListTagsLogGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsLogGroup", reflect.TypeOf((*MockCloudwatchlogsClient)(nil).ListTagsLogGroup), varargs...)
}
