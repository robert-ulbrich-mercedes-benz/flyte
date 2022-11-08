// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	admin "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin"

	filters "github.com/flyteorg/flytectl/pkg/filters"

	mock "github.com/stretchr/testify/mock"

	service "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/service"
)

// AdminFetcherExtInterface is an autogenerated mock type for the AdminFetcherExtInterface type
type AdminFetcherExtInterface struct {
	mock.Mock
}

type AdminFetcherExtInterface_AdminServiceClient struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_AdminServiceClient) Return(_a0 service.AdminServiceClient) *AdminFetcherExtInterface_AdminServiceClient {
	return &AdminFetcherExtInterface_AdminServiceClient{Call: _m.Call.Return(_a0)}
}

func (_m *AdminFetcherExtInterface) OnAdminServiceClient() *AdminFetcherExtInterface_AdminServiceClient {
	c_call := _m.On("AdminServiceClient")
	return &AdminFetcherExtInterface_AdminServiceClient{Call: c_call}
}

func (_m *AdminFetcherExtInterface) OnAdminServiceClientMatch(matchers ...interface{}) *AdminFetcherExtInterface_AdminServiceClient {
	c_call := _m.On("AdminServiceClient", matchers...)
	return &AdminFetcherExtInterface_AdminServiceClient{Call: c_call}
}

// AdminServiceClient provides a mock function with given fields:
func (_m *AdminFetcherExtInterface) AdminServiceClient() service.AdminServiceClient {
	ret := _m.Called()

	var r0 service.AdminServiceClient
	if rf, ok := ret.Get(0).(func() service.AdminServiceClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(service.AdminServiceClient)
		}
	}

	return r0
}

type AdminFetcherExtInterface_FetchAllVerOfLP struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchAllVerOfLP) Return(_a0 []*admin.LaunchPlan, _a1 error) *AdminFetcherExtInterface_FetchAllVerOfLP {
	return &AdminFetcherExtInterface_FetchAllVerOfLP{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchAllVerOfLP(ctx context.Context, lpName string, project string, domain string, filter filters.Filters) *AdminFetcherExtInterface_FetchAllVerOfLP {
	c_call := _m.On("FetchAllVerOfLP", ctx, lpName, project, domain, filter)
	return &AdminFetcherExtInterface_FetchAllVerOfLP{Call: c_call}
}

func (_m *AdminFetcherExtInterface) OnFetchAllVerOfLPMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchAllVerOfLP {
	c_call := _m.On("FetchAllVerOfLP", matchers...)
	return &AdminFetcherExtInterface_FetchAllVerOfLP{Call: c_call}
}

// FetchAllVerOfLP provides a mock function with given fields: ctx, lpName, project, domain, filter
func (_m *AdminFetcherExtInterface) FetchAllVerOfLP(ctx context.Context, lpName string, project string, domain string, filter filters.Filters) ([]*admin.LaunchPlan, error) {
	ret := _m.Called(ctx, lpName, project, domain, filter)

	var r0 []*admin.LaunchPlan
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, filters.Filters) []*admin.LaunchPlan); ok {
		r0 = rf(ctx, lpName, project, domain, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*admin.LaunchPlan)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, filters.Filters) error); ok {
		r1 = rf(ctx, lpName, project, domain, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchAllVerOfTask struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchAllVerOfTask) Return(_a0 []*admin.Task, _a1 error) *AdminFetcherExtInterface_FetchAllVerOfTask {
	return &AdminFetcherExtInterface_FetchAllVerOfTask{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchAllVerOfTask(ctx context.Context, name string, project string, domain string, filter filters.Filters) *AdminFetcherExtInterface_FetchAllVerOfTask {
	c_call := _m.On("FetchAllVerOfTask", ctx, name, project, domain, filter)
	return &AdminFetcherExtInterface_FetchAllVerOfTask{Call: c_call}
}

func (_m *AdminFetcherExtInterface) OnFetchAllVerOfTaskMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchAllVerOfTask {
	c_call := _m.On("FetchAllVerOfTask", matchers...)
	return &AdminFetcherExtInterface_FetchAllVerOfTask{Call: c_call}
}

// FetchAllVerOfTask provides a mock function with given fields: ctx, name, project, domain, filter
func (_m *AdminFetcherExtInterface) FetchAllVerOfTask(ctx context.Context, name string, project string, domain string, filter filters.Filters) ([]*admin.Task, error) {
	ret := _m.Called(ctx, name, project, domain, filter)

	var r0 []*admin.Task
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, filters.Filters) []*admin.Task); ok {
		r0 = rf(ctx, name, project, domain, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*admin.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, filters.Filters) error); ok {
		r1 = rf(ctx, name, project, domain, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchAllVerOfWorkflow struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchAllVerOfWorkflow) Return(_a0 []*admin.Workflow, _a1 error) *AdminFetcherExtInterface_FetchAllVerOfWorkflow {
	return &AdminFetcherExtInterface_FetchAllVerOfWorkflow{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchAllVerOfWorkflow(ctx context.Context, name string, project string, domain string, filter filters.Filters) *AdminFetcherExtInterface_FetchAllVerOfWorkflow {
	c_call := _m.On("FetchAllVerOfWorkflow", ctx, name, project, domain, filter)
	return &AdminFetcherExtInterface_FetchAllVerOfWorkflow{Call: c_call}
}

func (_m *AdminFetcherExtInterface) OnFetchAllVerOfWorkflowMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchAllVerOfWorkflow {
	c_call := _m.On("FetchAllVerOfWorkflow", matchers...)
	return &AdminFetcherExtInterface_FetchAllVerOfWorkflow{Call: c_call}
}

// FetchAllVerOfWorkflow provides a mock function with given fields: ctx, name, project, domain, filter
func (_m *AdminFetcherExtInterface) FetchAllVerOfWorkflow(ctx context.Context, name string, project string, domain string, filter filters.Filters) ([]*admin.Workflow, error) {
	ret := _m.Called(ctx, name, project, domain, filter)

	var r0 []*admin.Workflow
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, filters.Filters) []*admin.Workflow); ok {
		r0 = rf(ctx, name, project, domain, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*admin.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, filters.Filters) error); ok {
		r1 = rf(ctx, name, project, domain, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchAllWorkflows struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchAllWorkflows) Return(_a0 []*admin.NamedEntity, _a1 error) *AdminFetcherExtInterface_FetchAllWorkflows {
	return &AdminFetcherExtInterface_FetchAllWorkflows{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchAllWorkflows(ctx context.Context, project string, domain string, filter filters.Filters) *AdminFetcherExtInterface_FetchAllWorkflows {
	c_call := _m.On("FetchAllWorkflows", ctx, project, domain, filter)
	return &AdminFetcherExtInterface_FetchAllWorkflows{Call: c_call}
}

func (_m *AdminFetcherExtInterface) OnFetchAllWorkflowsMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchAllWorkflows {
	c_call := _m.On("FetchAllWorkflows", matchers...)
	return &AdminFetcherExtInterface_FetchAllWorkflows{Call: c_call}
}

// FetchAllWorkflows provides a mock function with given fields: ctx, project, domain, filter
func (_m *AdminFetcherExtInterface) FetchAllWorkflows(ctx context.Context, project string, domain string, filter filters.Filters) ([]*admin.NamedEntity, error) {
	ret := _m.Called(ctx, project, domain, filter)

	var r0 []*admin.NamedEntity
	if rf, ok := ret.Get(0).(func(context.Context, string, string, filters.Filters) []*admin.NamedEntity); ok {
		r0 = rf(ctx, project, domain, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*admin.NamedEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, filters.Filters) error); ok {
		r1 = rf(ctx, project, domain, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchExecution struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchExecution) Return(_a0 *admin.Execution, _a1 error) *AdminFetcherExtInterface_FetchExecution {
	return &AdminFetcherExtInterface_FetchExecution{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchExecution(ctx context.Context, name string, project string, domain string) *AdminFetcherExtInterface_FetchExecution {
	c_call := _m.On("FetchExecution", ctx, name, project, domain)
	return &AdminFetcherExtInterface_FetchExecution{Call: c_call}
}

func (_m *AdminFetcherExtInterface) OnFetchExecutionMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchExecution {
	c_call := _m.On("FetchExecution", matchers...)
	return &AdminFetcherExtInterface_FetchExecution{Call: c_call}
}

// FetchExecution provides a mock function with given fields: ctx, name, project, domain
func (_m *AdminFetcherExtInterface) FetchExecution(ctx context.Context, name string, project string, domain string) (*admin.Execution, error) {
	ret := _m.Called(ctx, name, project, domain)

	var r0 *admin.Execution
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *admin.Execution); ok {
		r0 = rf(ctx, name, project, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Execution)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, name, project, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchLPLatestVersion struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchLPLatestVersion) Return(_a0 *admin.LaunchPlan, _a1 error) *AdminFetcherExtInterface_FetchLPLatestVersion {
	return &AdminFetcherExtInterface_FetchLPLatestVersion{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchLPLatestVersion(ctx context.Context, name string, project string, domain string, filter filters.Filters) *AdminFetcherExtInterface_FetchLPLatestVersion {
	c_call := _m.On("FetchLPLatestVersion", ctx, name, project, domain, filter)
	return &AdminFetcherExtInterface_FetchLPLatestVersion{Call: c_call}
}

func (_m *AdminFetcherExtInterface) OnFetchLPLatestVersionMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchLPLatestVersion {
	c_call := _m.On("FetchLPLatestVersion", matchers...)
	return &AdminFetcherExtInterface_FetchLPLatestVersion{Call: c_call}
}

// FetchLPLatestVersion provides a mock function with given fields: ctx, name, project, domain, filter
func (_m *AdminFetcherExtInterface) FetchLPLatestVersion(ctx context.Context, name string, project string, domain string, filter filters.Filters) (*admin.LaunchPlan, error) {
	ret := _m.Called(ctx, name, project, domain, filter)

	var r0 *admin.LaunchPlan
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, filters.Filters) *admin.LaunchPlan); ok {
		r0 = rf(ctx, name, project, domain, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.LaunchPlan)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, filters.Filters) error); ok {
		r1 = rf(ctx, name, project, domain, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchLPVersion struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchLPVersion) Return(_a0 *admin.LaunchPlan, _a1 error) *AdminFetcherExtInterface_FetchLPVersion {
	return &AdminFetcherExtInterface_FetchLPVersion{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchLPVersion(ctx context.Context, name string, version string, project string, domain string) *AdminFetcherExtInterface_FetchLPVersion {
	c_call := _m.On("FetchLPVersion", ctx, name, version, project, domain)
	return &AdminFetcherExtInterface_FetchLPVersion{Call: c_call}
}

func (_m *AdminFetcherExtInterface) OnFetchLPVersionMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchLPVersion {
	c_call := _m.On("FetchLPVersion", matchers...)
	return &AdminFetcherExtInterface_FetchLPVersion{Call: c_call}
}

// FetchLPVersion provides a mock function with given fields: ctx, name, version, project, domain
func (_m *AdminFetcherExtInterface) FetchLPVersion(ctx context.Context, name string, version string, project string, domain string) (*admin.LaunchPlan, error) {
	ret := _m.Called(ctx, name, version, project, domain)

	var r0 *admin.LaunchPlan
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) *admin.LaunchPlan); ok {
		r0 = rf(ctx, name, version, project, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.LaunchPlan)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, name, version, project, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchNodeExecutionData struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchNodeExecutionData) Return(_a0 *admin.NodeExecutionGetDataResponse, _a1 error) *AdminFetcherExtInterface_FetchNodeExecutionData {
	return &AdminFetcherExtInterface_FetchNodeExecutionData{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchNodeExecutionData(ctx context.Context, nodeID string, execName string, project string, domain string) *AdminFetcherExtInterface_FetchNodeExecutionData {
	c_call := _m.On("FetchNodeExecutionData", ctx, nodeID, execName, project, domain)
	return &AdminFetcherExtInterface_FetchNodeExecutionData{Call: c_call}
}

func (_m *AdminFetcherExtInterface) OnFetchNodeExecutionDataMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchNodeExecutionData {
	c_call := _m.On("FetchNodeExecutionData", matchers...)
	return &AdminFetcherExtInterface_FetchNodeExecutionData{Call: c_call}
}

// FetchNodeExecutionData provides a mock function with given fields: ctx, nodeID, execName, project, domain
func (_m *AdminFetcherExtInterface) FetchNodeExecutionData(ctx context.Context, nodeID string, execName string, project string, domain string) (*admin.NodeExecutionGetDataResponse, error) {
	ret := _m.Called(ctx, nodeID, execName, project, domain)

	var r0 *admin.NodeExecutionGetDataResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) *admin.NodeExecutionGetDataResponse); ok {
		r0 = rf(ctx, nodeID, execName, project, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.NodeExecutionGetDataResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, nodeID, execName, project, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchNodeExecutionDetails struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchNodeExecutionDetails) Return(_a0 *admin.NodeExecutionList, _a1 error) *AdminFetcherExtInterface_FetchNodeExecutionDetails {
	return &AdminFetcherExtInterface_FetchNodeExecutionDetails{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchNodeExecutionDetails(ctx context.Context, name string, project string, domain string, uniqueParentID string) *AdminFetcherExtInterface_FetchNodeExecutionDetails {
	c_call := _m.On("FetchNodeExecutionDetails", ctx, name, project, domain, uniqueParentID)
	return &AdminFetcherExtInterface_FetchNodeExecutionDetails{Call: c_call}
}

func (_m *AdminFetcherExtInterface) OnFetchNodeExecutionDetailsMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchNodeExecutionDetails {
	c_call := _m.On("FetchNodeExecutionDetails", matchers...)
	return &AdminFetcherExtInterface_FetchNodeExecutionDetails{Call: c_call}
}

// FetchNodeExecutionDetails provides a mock function with given fields: ctx, name, project, domain, uniqueParentID
func (_m *AdminFetcherExtInterface) FetchNodeExecutionDetails(ctx context.Context, name string, project string, domain string, uniqueParentID string) (*admin.NodeExecutionList, error) {
	ret := _m.Called(ctx, name, project, domain, uniqueParentID)

	var r0 *admin.NodeExecutionList
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) *admin.NodeExecutionList); ok {
		r0 = rf(ctx, name, project, domain, uniqueParentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.NodeExecutionList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, name, project, domain, uniqueParentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchProjectAttributes struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchProjectAttributes) Return(_a0 *admin.ProjectAttributesGetResponse, _a1 error) *AdminFetcherExtInterface_FetchProjectAttributes {
	return &AdminFetcherExtInterface_FetchProjectAttributes{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchProjectAttributes(ctx context.Context, project string, rsType admin.MatchableResource) *AdminFetcherExtInterface_FetchProjectAttributes {
	c_call := _m.On("FetchProjectAttributes", ctx, project, rsType)
	return &AdminFetcherExtInterface_FetchProjectAttributes{Call: c_call}
}

func (_m *AdminFetcherExtInterface) OnFetchProjectAttributesMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchProjectAttributes {
	c_call := _m.On("FetchProjectAttributes", matchers...)
	return &AdminFetcherExtInterface_FetchProjectAttributes{Call: c_call}
}

// FetchProjectAttributes provides a mock function with given fields: ctx, project, rsType
func (_m *AdminFetcherExtInterface) FetchProjectAttributes(ctx context.Context, project string, rsType admin.MatchableResource) (*admin.ProjectAttributesGetResponse, error) {
	ret := _m.Called(ctx, project, rsType)

	var r0 *admin.ProjectAttributesGetResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, admin.MatchableResource) *admin.ProjectAttributesGetResponse); ok {
		r0 = rf(ctx, project, rsType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.ProjectAttributesGetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, admin.MatchableResource) error); ok {
		r1 = rf(ctx, project, rsType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchProjectDomainAttributes struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchProjectDomainAttributes) Return(_a0 *admin.ProjectDomainAttributesGetResponse, _a1 error) *AdminFetcherExtInterface_FetchProjectDomainAttributes {
	return &AdminFetcherExtInterface_FetchProjectDomainAttributes{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchProjectDomainAttributes(ctx context.Context, project string, domain string, rsType admin.MatchableResource) *AdminFetcherExtInterface_FetchProjectDomainAttributes {
	c_call := _m.On("FetchProjectDomainAttributes", ctx, project, domain, rsType)
	return &AdminFetcherExtInterface_FetchProjectDomainAttributes{Call: c_call}
}

func (_m *AdminFetcherExtInterface) OnFetchProjectDomainAttributesMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchProjectDomainAttributes {
	c_call := _m.On("FetchProjectDomainAttributes", matchers...)
	return &AdminFetcherExtInterface_FetchProjectDomainAttributes{Call: c_call}
}

// FetchProjectDomainAttributes provides a mock function with given fields: ctx, project, domain, rsType
func (_m *AdminFetcherExtInterface) FetchProjectDomainAttributes(ctx context.Context, project string, domain string, rsType admin.MatchableResource) (*admin.ProjectDomainAttributesGetResponse, error) {
	ret := _m.Called(ctx, project, domain, rsType)

	var r0 *admin.ProjectDomainAttributesGetResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, string, admin.MatchableResource) *admin.ProjectDomainAttributesGetResponse); ok {
		r0 = rf(ctx, project, domain, rsType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.ProjectDomainAttributesGetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, admin.MatchableResource) error); ok {
		r1 = rf(ctx, project, domain, rsType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchTaskExecutionsOnNode struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchTaskExecutionsOnNode) Return(_a0 *admin.TaskExecutionList, _a1 error) *AdminFetcherExtInterface_FetchTaskExecutionsOnNode {
	return &AdminFetcherExtInterface_FetchTaskExecutionsOnNode{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchTaskExecutionsOnNode(ctx context.Context, nodeID string, execName string, project string, domain string) *AdminFetcherExtInterface_FetchTaskExecutionsOnNode {
	c_call := _m.On("FetchTaskExecutionsOnNode", ctx, nodeID, execName, project, domain)
	return &AdminFetcherExtInterface_FetchTaskExecutionsOnNode{Call: c_call}
}

func (_m *AdminFetcherExtInterface) OnFetchTaskExecutionsOnNodeMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchTaskExecutionsOnNode {
	c_call := _m.On("FetchTaskExecutionsOnNode", matchers...)
	return &AdminFetcherExtInterface_FetchTaskExecutionsOnNode{Call: c_call}
}

// FetchTaskExecutionsOnNode provides a mock function with given fields: ctx, nodeID, execName, project, domain
func (_m *AdminFetcherExtInterface) FetchTaskExecutionsOnNode(ctx context.Context, nodeID string, execName string, project string, domain string) (*admin.TaskExecutionList, error) {
	ret := _m.Called(ctx, nodeID, execName, project, domain)

	var r0 *admin.TaskExecutionList
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) *admin.TaskExecutionList); ok {
		r0 = rf(ctx, nodeID, execName, project, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.TaskExecutionList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, nodeID, execName, project, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchTaskLatestVersion struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchTaskLatestVersion) Return(_a0 *admin.Task, _a1 error) *AdminFetcherExtInterface_FetchTaskLatestVersion {
	return &AdminFetcherExtInterface_FetchTaskLatestVersion{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchTaskLatestVersion(ctx context.Context, name string, project string, domain string, filter filters.Filters) *AdminFetcherExtInterface_FetchTaskLatestVersion {
	c_call := _m.On("FetchTaskLatestVersion", ctx, name, project, domain, filter)
	return &AdminFetcherExtInterface_FetchTaskLatestVersion{Call: c_call}
}

func (_m *AdminFetcherExtInterface) OnFetchTaskLatestVersionMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchTaskLatestVersion {
	c_call := _m.On("FetchTaskLatestVersion", matchers...)
	return &AdminFetcherExtInterface_FetchTaskLatestVersion{Call: c_call}
}

// FetchTaskLatestVersion provides a mock function with given fields: ctx, name, project, domain, filter
func (_m *AdminFetcherExtInterface) FetchTaskLatestVersion(ctx context.Context, name string, project string, domain string, filter filters.Filters) (*admin.Task, error) {
	ret := _m.Called(ctx, name, project, domain, filter)

	var r0 *admin.Task
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, filters.Filters) *admin.Task); ok {
		r0 = rf(ctx, name, project, domain, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, filters.Filters) error); ok {
		r1 = rf(ctx, name, project, domain, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchTaskVersion struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchTaskVersion) Return(_a0 *admin.Task, _a1 error) *AdminFetcherExtInterface_FetchTaskVersion {
	return &AdminFetcherExtInterface_FetchTaskVersion{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchTaskVersion(ctx context.Context, name string, version string, project string, domain string) *AdminFetcherExtInterface_FetchTaskVersion {
	c_call := _m.On("FetchTaskVersion", ctx, name, version, project, domain)
	return &AdminFetcherExtInterface_FetchTaskVersion{Call: c_call}
}

func (_m *AdminFetcherExtInterface) OnFetchTaskVersionMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchTaskVersion {
	c_call := _m.On("FetchTaskVersion", matchers...)
	return &AdminFetcherExtInterface_FetchTaskVersion{Call: c_call}
}

// FetchTaskVersion provides a mock function with given fields: ctx, name, version, project, domain
func (_m *AdminFetcherExtInterface) FetchTaskVersion(ctx context.Context, name string, version string, project string, domain string) (*admin.Task, error) {
	ret := _m.Called(ctx, name, version, project, domain)

	var r0 *admin.Task
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) *admin.Task); ok {
		r0 = rf(ctx, name, version, project, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, name, version, project, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchWorkflowAttributes struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchWorkflowAttributes) Return(_a0 *admin.WorkflowAttributesGetResponse, _a1 error) *AdminFetcherExtInterface_FetchWorkflowAttributes {
	return &AdminFetcherExtInterface_FetchWorkflowAttributes{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchWorkflowAttributes(ctx context.Context, project string, domain string, name string, rsType admin.MatchableResource) *AdminFetcherExtInterface_FetchWorkflowAttributes {
	c_call := _m.On("FetchWorkflowAttributes", ctx, project, domain, name, rsType)
	return &AdminFetcherExtInterface_FetchWorkflowAttributes{Call: c_call}
}

func (_m *AdminFetcherExtInterface) OnFetchWorkflowAttributesMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchWorkflowAttributes {
	c_call := _m.On("FetchWorkflowAttributes", matchers...)
	return &AdminFetcherExtInterface_FetchWorkflowAttributes{Call: c_call}
}

// FetchWorkflowAttributes provides a mock function with given fields: ctx, project, domain, name, rsType
func (_m *AdminFetcherExtInterface) FetchWorkflowAttributes(ctx context.Context, project string, domain string, name string, rsType admin.MatchableResource) (*admin.WorkflowAttributesGetResponse, error) {
	ret := _m.Called(ctx, project, domain, name, rsType)

	var r0 *admin.WorkflowAttributesGetResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, admin.MatchableResource) *admin.WorkflowAttributesGetResponse); ok {
		r0 = rf(ctx, project, domain, name, rsType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.WorkflowAttributesGetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, admin.MatchableResource) error); ok {
		r1 = rf(ctx, project, domain, name, rsType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchWorkflowLatestVersion struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchWorkflowLatestVersion) Return(_a0 *admin.Workflow, _a1 error) *AdminFetcherExtInterface_FetchWorkflowLatestVersion {
	return &AdminFetcherExtInterface_FetchWorkflowLatestVersion{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchWorkflowLatestVersion(ctx context.Context, name string, project string, domain string, filter filters.Filters) *AdminFetcherExtInterface_FetchWorkflowLatestVersion {
	c_call := _m.On("FetchWorkflowLatestVersion", ctx, name, project, domain, filter)
	return &AdminFetcherExtInterface_FetchWorkflowLatestVersion{Call: c_call}
}

func (_m *AdminFetcherExtInterface) OnFetchWorkflowLatestVersionMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchWorkflowLatestVersion {
	c_call := _m.On("FetchWorkflowLatestVersion", matchers...)
	return &AdminFetcherExtInterface_FetchWorkflowLatestVersion{Call: c_call}
}

// FetchWorkflowLatestVersion provides a mock function with given fields: ctx, name, project, domain, filter
func (_m *AdminFetcherExtInterface) FetchWorkflowLatestVersion(ctx context.Context, name string, project string, domain string, filter filters.Filters) (*admin.Workflow, error) {
	ret := _m.Called(ctx, name, project, domain, filter)

	var r0 *admin.Workflow
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, filters.Filters) *admin.Workflow); ok {
		r0 = rf(ctx, name, project, domain, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, filters.Filters) error); ok {
		r1 = rf(ctx, name, project, domain, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_FetchWorkflowVersion struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_FetchWorkflowVersion) Return(_a0 *admin.Workflow, _a1 error) *AdminFetcherExtInterface_FetchWorkflowVersion {
	return &AdminFetcherExtInterface_FetchWorkflowVersion{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnFetchWorkflowVersion(ctx context.Context, name string, version string, project string, domain string) *AdminFetcherExtInterface_FetchWorkflowVersion {
	c_call := _m.On("FetchWorkflowVersion", ctx, name, version, project, domain)
	return &AdminFetcherExtInterface_FetchWorkflowVersion{Call: c_call}
}

func (_m *AdminFetcherExtInterface) OnFetchWorkflowVersionMatch(matchers ...interface{}) *AdminFetcherExtInterface_FetchWorkflowVersion {
	c_call := _m.On("FetchWorkflowVersion", matchers...)
	return &AdminFetcherExtInterface_FetchWorkflowVersion{Call: c_call}
}

// FetchWorkflowVersion provides a mock function with given fields: ctx, name, version, project, domain
func (_m *AdminFetcherExtInterface) FetchWorkflowVersion(ctx context.Context, name string, version string, project string, domain string) (*admin.Workflow, error) {
	ret := _m.Called(ctx, name, version, project, domain)

	var r0 *admin.Workflow
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) *admin.Workflow); ok {
		r0 = rf(ctx, name, version, project, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, name, version, project, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_ListExecution struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_ListExecution) Return(_a0 *admin.ExecutionList, _a1 error) *AdminFetcherExtInterface_ListExecution {
	return &AdminFetcherExtInterface_ListExecution{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnListExecution(ctx context.Context, project string, domain string, filter filters.Filters) *AdminFetcherExtInterface_ListExecution {
	c_call := _m.On("ListExecution", ctx, project, domain, filter)
	return &AdminFetcherExtInterface_ListExecution{Call: c_call}
}

func (_m *AdminFetcherExtInterface) OnListExecutionMatch(matchers ...interface{}) *AdminFetcherExtInterface_ListExecution {
	c_call := _m.On("ListExecution", matchers...)
	return &AdminFetcherExtInterface_ListExecution{Call: c_call}
}

// ListExecution provides a mock function with given fields: ctx, project, domain, filter
func (_m *AdminFetcherExtInterface) ListExecution(ctx context.Context, project string, domain string, filter filters.Filters) (*admin.ExecutionList, error) {
	ret := _m.Called(ctx, project, domain, filter)

	var r0 *admin.ExecutionList
	if rf, ok := ret.Get(0).(func(context.Context, string, string, filters.Filters) *admin.ExecutionList); ok {
		r0 = rf(ctx, project, domain, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.ExecutionList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, filters.Filters) error); ok {
		r1 = rf(ctx, project, domain, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AdminFetcherExtInterface_ListProjects struct {
	*mock.Call
}

func (_m AdminFetcherExtInterface_ListProjects) Return(_a0 *admin.Projects, _a1 error) *AdminFetcherExtInterface_ListProjects {
	return &AdminFetcherExtInterface_ListProjects{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AdminFetcherExtInterface) OnListProjects(ctx context.Context, filter filters.Filters) *AdminFetcherExtInterface_ListProjects {
	c_call := _m.On("ListProjects", ctx, filter)
	return &AdminFetcherExtInterface_ListProjects{Call: c_call}
}

func (_m *AdminFetcherExtInterface) OnListProjectsMatch(matchers ...interface{}) *AdminFetcherExtInterface_ListProjects {
	c_call := _m.On("ListProjects", matchers...)
	return &AdminFetcherExtInterface_ListProjects{Call: c_call}
}

// ListProjects provides a mock function with given fields: ctx, filter
func (_m *AdminFetcherExtInterface) ListProjects(ctx context.Context, filter filters.Filters) (*admin.Projects, error) {
	ret := _m.Called(ctx, filter)

	var r0 *admin.Projects
	if rf, ok := ret.Get(0).(func(context.Context, filters.Filters) *admin.Projects); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Projects)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, filters.Filters) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
