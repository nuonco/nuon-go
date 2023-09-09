// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package nuon is a generated GoMock package.
package nuon

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	models "github.com/nuonco/nuon-go/models"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// SetOrgID mocks base method
func (m *MockClient) SetOrgID(orgID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetOrgID", orgID)
}

// SetOrgID indicates an expected call of SetOrgID
func (mr *MockClientMockRecorder) SetOrgID(orgID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOrgID", reflect.TypeOf((*MockClient)(nil).SetOrgID), orgID)
}

// GetOrgs mocks base method
func (m *MockClient) GetOrgs(ctx context.Context) ([]*models.AppOrg, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgs", ctx)
	ret0, _ := ret[0].([]*models.AppOrg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgs indicates an expected call of GetOrgs
func (mr *MockClientMockRecorder) GetOrgs(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgs", reflect.TypeOf((*MockClient)(nil).GetOrgs), ctx)
}

// CreateOrg mocks base method
func (m *MockClient) CreateOrg(ctx context.Context, req *models.ServiceCreateOrgRequest) (*models.AppOrg, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrg", ctx, req)
	ret0, _ := ret[0].(*models.AppOrg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrg indicates an expected call of CreateOrg
func (mr *MockClientMockRecorder) CreateOrg(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrg", reflect.TypeOf((*MockClient)(nil).CreateOrg), ctx, req)
}

// GetOrg mocks base method
func (m *MockClient) GetOrg(ctx context.Context) (*models.AppOrg, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrg", ctx)
	ret0, _ := ret[0].(*models.AppOrg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrg indicates an expected call of GetOrg
func (mr *MockClientMockRecorder) GetOrg(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrg", reflect.TypeOf((*MockClient)(nil).GetOrg), ctx)
}

// UpdateOrg mocks base method
func (m *MockClient) UpdateOrg(ctx context.Context, req *models.ServiceUpdateOrgRequest) (*models.AppOrg, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrg", ctx, req)
	ret0, _ := ret[0].(*models.AppOrg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrg indicates an expected call of UpdateOrg
func (mr *MockClientMockRecorder) UpdateOrg(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrg", reflect.TypeOf((*MockClient)(nil).UpdateOrg), ctx, req)
}

// DeleteOrg mocks base method
func (m *MockClient) DeleteOrg(ctx context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrg", ctx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOrg indicates an expected call of DeleteOrg
func (mr *MockClientMockRecorder) DeleteOrg(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrg", reflect.TypeOf((*MockClient)(nil).DeleteOrg), ctx)
}

// CreateOrgUser mocks base method
func (m *MockClient) CreateOrgUser(ctx context.Context, req *models.ServiceCreateOrgUserRequest) (*models.AppUserOrg, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrgUser", ctx, req)
	ret0, _ := ret[0].(*models.AppUserOrg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrgUser indicates an expected call of CreateOrgUser
func (mr *MockClientMockRecorder) CreateOrgUser(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrgUser", reflect.TypeOf((*MockClient)(nil).CreateOrgUser), ctx, req)
}

// GetApp mocks base method
func (m *MockClient) GetApp(ctx context.Context, appID string) (*models.AppApp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApp", ctx, appID)
	ret0, _ := ret[0].(*models.AppApp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApp indicates an expected call of GetApp
func (mr *MockClientMockRecorder) GetApp(ctx, appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApp", reflect.TypeOf((*MockClient)(nil).GetApp), ctx, appID)
}

// GetApps mocks base method
func (m *MockClient) GetApps(ctx context.Context) ([]*models.AppApp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApps", ctx)
	ret0, _ := ret[0].([]*models.AppApp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApps indicates an expected call of GetApps
func (mr *MockClientMockRecorder) GetApps(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApps", reflect.TypeOf((*MockClient)(nil).GetApps), ctx)
}

// CreateApp mocks base method
func (m *MockClient) CreateApp(ctx context.Context, req *models.ServiceCreateAppRequest) (*models.AppApp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApp", ctx, req)
	ret0, _ := ret[0].(*models.AppApp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateApp indicates an expected call of CreateApp
func (mr *MockClientMockRecorder) CreateApp(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApp", reflect.TypeOf((*MockClient)(nil).CreateApp), ctx, req)
}

// UpdateApp mocks base method
func (m *MockClient) UpdateApp(ctx context.Context, appID string, req *models.ServiceUpdateAppRequest) (*models.AppApp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApp", ctx, appID, req)
	ret0, _ := ret[0].(*models.AppApp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateApp indicates an expected call of UpdateApp
func (mr *MockClientMockRecorder) UpdateApp(ctx, appID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApp", reflect.TypeOf((*MockClient)(nil).UpdateApp), ctx, appID, req)
}

// DeleteApp mocks base method
func (m *MockClient) DeleteApp(ctx context.Context, appID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteApp", ctx, appID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteApp indicates an expected call of DeleteApp
func (mr *MockClientMockRecorder) DeleteApp(ctx, appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApp", reflect.TypeOf((*MockClient)(nil).DeleteApp), ctx, appID)
}

// UpdateAppSandbox mocks base method
func (m *MockClient) UpdateAppSandbox(ctx context.Context, appID string, req *models.ServiceUpdateAppSandboxRequest) (*models.AppApp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAppSandbox", ctx, appID, req)
	ret0, _ := ret[0].(*models.AppApp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAppSandbox indicates an expected call of UpdateAppSandbox
func (mr *MockClientMockRecorder) UpdateAppSandbox(ctx, appID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAppSandbox", reflect.TypeOf((*MockClient)(nil).UpdateAppSandbox), ctx, appID, req)
}

// GetCurrentUser mocks base method
func (m *MockClient) GetCurrentUser(ctx context.Context) (*models.AppUserToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentUser", ctx)
	ret0, _ := ret[0].(*models.AppUserToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentUser indicates an expected call of GetCurrentUser
func (mr *MockClientMockRecorder) GetCurrentUser(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentUser", reflect.TypeOf((*MockClient)(nil).GetCurrentUser), ctx)
}

// PublishMetrics mocks base method
func (m *MockClient) PublishMetrics(ctx context.Context, req []*models.ServicePublishMetricInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishMetrics", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishMetrics indicates an expected call of PublishMetrics
func (mr *MockClientMockRecorder) PublishMetrics(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishMetrics", reflect.TypeOf((*MockClient)(nil).PublishMetrics), ctx, req)
}

// GetSandboxes mocks base method
func (m *MockClient) GetSandboxes(ctx context.Context) ([]*models.AppSandbox, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSandboxes", ctx)
	ret0, _ := ret[0].([]*models.AppSandbox)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSandboxes indicates an expected call of GetSandboxes
func (mr *MockClientMockRecorder) GetSandboxes(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSandboxes", reflect.TypeOf((*MockClient)(nil).GetSandboxes), ctx)
}

// GetSandbox mocks base method
func (m *MockClient) GetSandbox(ctx context.Context, sandboxID string) (*models.AppSandbox, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSandbox", ctx, sandboxID)
	ret0, _ := ret[0].(*models.AppSandbox)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSandbox indicates an expected call of GetSandbox
func (mr *MockClientMockRecorder) GetSandbox(ctx, sandboxID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSandbox", reflect.TypeOf((*MockClient)(nil).GetSandbox), ctx, sandboxID)
}

// GetSandboxReleases mocks base method
func (m *MockClient) GetSandboxReleases(ctx context.Context, sandboxID string) ([]*models.AppSandboxRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSandboxReleases", ctx, sandboxID)
	ret0, _ := ret[0].([]*models.AppSandboxRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSandboxReleases indicates an expected call of GetSandboxReleases
func (mr *MockClientMockRecorder) GetSandboxReleases(ctx, sandboxID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSandboxReleases", reflect.TypeOf((*MockClient)(nil).GetSandboxReleases), ctx, sandboxID)
}

// CreateVCSConnection mocks base method
func (m *MockClient) CreateVCSConnection(ctx context.Context, req *models.ServiceCreateConnectionRequest) (*models.AppVCSConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVCSConnection", ctx, req)
	ret0, _ := ret[0].(*models.AppVCSConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVCSConnection indicates an expected call of CreateVCSConnection
func (mr *MockClientMockRecorder) CreateVCSConnection(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVCSConnection", reflect.TypeOf((*MockClient)(nil).CreateVCSConnection), ctx, req)
}

// GetVCSConnections mocks base method
func (m *MockClient) GetVCSConnections(ctx context.Context) ([]*models.AppVCSConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVCSConnections", ctx)
	ret0, _ := ret[0].([]*models.AppVCSConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVCSConnections indicates an expected call of GetVCSConnections
func (mr *MockClientMockRecorder) GetVCSConnections(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVCSConnections", reflect.TypeOf((*MockClient)(nil).GetVCSConnections), ctx)
}

// GetVCSConnection mocks base method
func (m *MockClient) GetVCSConnection(ctx context.Context, connID string) (*models.AppVCSConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVCSConnection", ctx, connID)
	ret0, _ := ret[0].(*models.AppVCSConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVCSConnection indicates an expected call of GetVCSConnection
func (mr *MockClientMockRecorder) GetVCSConnection(ctx, connID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVCSConnection", reflect.TypeOf((*MockClient)(nil).GetVCSConnection), ctx, connID)
}

// GetAllVCSConnectedRepos mocks base method
func (m *MockClient) GetAllVCSConnectedRepos(ctx context.Context) ([]*models.ServiceRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllVCSConnectedRepos", ctx)
	ret0, _ := ret[0].([]*models.ServiceRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllVCSConnectedRepos indicates an expected call of GetAllVCSConnectedRepos
func (mr *MockClientMockRecorder) GetAllVCSConnectedRepos(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllVCSConnectedRepos", reflect.TypeOf((*MockClient)(nil).GetAllVCSConnectedRepos), ctx)
}

// CreateInstall mocks base method
func (m *MockClient) CreateInstall(ctx context.Context, appID string, req *models.ServiceCreateInstallRequest) (*models.AppInstall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInstall", ctx, appID, req)
	ret0, _ := ret[0].(*models.AppInstall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInstall indicates an expected call of CreateInstall
func (mr *MockClientMockRecorder) CreateInstall(ctx, appID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInstall", reflect.TypeOf((*MockClient)(nil).CreateInstall), ctx, appID, req)
}

// GetAppInstalls mocks base method
func (m *MockClient) GetAppInstalls(ctx context.Context, appID string) ([]*models.AppInstall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppInstalls", ctx, appID)
	ret0, _ := ret[0].([]*models.AppInstall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppInstalls indicates an expected call of GetAppInstalls
func (mr *MockClientMockRecorder) GetAppInstalls(ctx, appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppInstalls", reflect.TypeOf((*MockClient)(nil).GetAppInstalls), ctx, appID)
}

// GetAllInstalls mocks base method
func (m *MockClient) GetAllInstalls(ctx context.Context) ([]*models.AppInstall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllInstalls", ctx)
	ret0, _ := ret[0].([]*models.AppInstall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllInstalls indicates an expected call of GetAllInstalls
func (mr *MockClientMockRecorder) GetAllInstalls(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllInstalls", reflect.TypeOf((*MockClient)(nil).GetAllInstalls), ctx)
}

// GetInstall mocks base method
func (m *MockClient) GetInstall(ctx context.Context, installID string) (*models.AppInstall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstall", ctx, installID)
	ret0, _ := ret[0].(*models.AppInstall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstall indicates an expected call of GetInstall
func (mr *MockClientMockRecorder) GetInstall(ctx, installID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstall", reflect.TypeOf((*MockClient)(nil).GetInstall), ctx, installID)
}

// UpdateInstall mocks base method
func (m *MockClient) UpdateInstall(ctx context.Context, installID string, req *models.ServiceUpdateInstallRequest) (*models.AppInstall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInstall", ctx, installID, req)
	ret0, _ := ret[0].(*models.AppInstall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateInstall indicates an expected call of UpdateInstall
func (mr *MockClientMockRecorder) UpdateInstall(ctx, installID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInstall", reflect.TypeOf((*MockClient)(nil).UpdateInstall), ctx, installID, req)
}

// DeleteInstall mocks base method
func (m *MockClient) DeleteInstall(ctx context.Context, installID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInstall", ctx, installID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteInstall indicates an expected call of DeleteInstall
func (mr *MockClientMockRecorder) DeleteInstall(ctx, installID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInstall", reflect.TypeOf((*MockClient)(nil).DeleteInstall), ctx, installID)
}

// GetInstallDeploys mocks base method
func (m *MockClient) GetInstallDeploys(ctx context.Context, installID string) ([]*models.AppInstallDeploy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstallDeploys", ctx, installID)
	ret0, _ := ret[0].([]*models.AppInstallDeploy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstallDeploys indicates an expected call of GetInstallDeploys
func (mr *MockClientMockRecorder) GetInstallDeploys(ctx, installID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstallDeploys", reflect.TypeOf((*MockClient)(nil).GetInstallDeploys), ctx, installID)
}

// CreateInstallDeploy mocks base method
func (m *MockClient) CreateInstallDeploy(ctx context.Context, installID string, req *models.ServiceCreateInstallDeployRequest) (*models.AppInstallDeploy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInstallDeploy", ctx, installID, req)
	ret0, _ := ret[0].(*models.AppInstallDeploy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInstallDeploy indicates an expected call of CreateInstallDeploy
func (mr *MockClientMockRecorder) CreateInstallDeploy(ctx, installID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInstallDeploy", reflect.TypeOf((*MockClient)(nil).CreateInstallDeploy), ctx, installID, req)
}

// GetInstallDeploy mocks base method
func (m *MockClient) GetInstallDeploy(ctx context.Context, installID, deployID string) (*models.AppInstallDeploy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstallDeploy", ctx, installID, deployID)
	ret0, _ := ret[0].(*models.AppInstallDeploy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstallDeploy indicates an expected call of GetInstallDeploy
func (mr *MockClientMockRecorder) GetInstallDeploy(ctx, installID, deployID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstallDeploy", reflect.TypeOf((*MockClient)(nil).GetInstallDeploy), ctx, installID, deployID)
}

// GetInstallLatestDeploy mocks base method
func (m *MockClient) GetInstallLatestDeploy(ctx context.Context, installID string) (*models.AppInstallDeploy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstallLatestDeploy", ctx, installID)
	ret0, _ := ret[0].(*models.AppInstallDeploy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstallLatestDeploy indicates an expected call of GetInstallLatestDeploy
func (mr *MockClientMockRecorder) GetInstallLatestDeploy(ctx, installID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstallLatestDeploy", reflect.TypeOf((*MockClient)(nil).GetInstallLatestDeploy), ctx, installID)
}

// GetInstallDeployLogs mocks base method
func (m *MockClient) GetInstallDeployLogs(ctx context.Context, installID, deployID string) ([]models.ServiceDeployLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstallDeployLogs", ctx, installID, deployID)
	ret0, _ := ret[0].([]models.ServiceDeployLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstallDeployLogs indicates an expected call of GetInstallDeployLogs
func (mr *MockClientMockRecorder) GetInstallDeployLogs(ctx, installID, deployID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstallDeployLogs", reflect.TypeOf((*MockClient)(nil).GetInstallDeployLogs), ctx, installID, deployID)
}

// GetInstallComponents mocks base method
func (m *MockClient) GetInstallComponents(ctx context.Context, installID string) ([]*models.AppInstallComponent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstallComponents", ctx, installID)
	ret0, _ := ret[0].([]*models.AppInstallComponent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstallComponents indicates an expected call of GetInstallComponents
func (mr *MockClientMockRecorder) GetInstallComponents(ctx, installID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstallComponents", reflect.TypeOf((*MockClient)(nil).GetInstallComponents), ctx, installID)
}

// GetInstallComponentDeploys mocks base method
func (m *MockClient) GetInstallComponentDeploys(ctx context.Context, installID, componentID string) ([]*models.AppInstallDeploy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstallComponentDeploys", ctx, installID, componentID)
	ret0, _ := ret[0].([]*models.AppInstallDeploy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstallComponentDeploys indicates an expected call of GetInstallComponentDeploys
func (mr *MockClientMockRecorder) GetInstallComponentDeploys(ctx, installID, componentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstallComponentDeploys", reflect.TypeOf((*MockClient)(nil).GetInstallComponentDeploys), ctx, installID, componentID)
}

// GetInstallComponentLatestDeploy mocks base method
func (m *MockClient) GetInstallComponentLatestDeploy(ctx context.Context, installID, componentID string) (*models.AppInstallDeploy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstallComponentLatestDeploy", ctx, installID, componentID)
	ret0, _ := ret[0].(*models.AppInstallDeploy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstallComponentLatestDeploy indicates an expected call of GetInstallComponentLatestDeploy
func (mr *MockClientMockRecorder) GetInstallComponentLatestDeploy(ctx, installID, componentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstallComponentLatestDeploy", reflect.TypeOf((*MockClient)(nil).GetInstallComponentLatestDeploy), ctx, installID, componentID)
}

// GetAllComponents mocks base method
func (m *MockClient) GetAllComponents(ctx context.Context) ([]*models.AppComponent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllComponents", ctx)
	ret0, _ := ret[0].([]*models.AppComponent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllComponents indicates an expected call of GetAllComponents
func (mr *MockClientMockRecorder) GetAllComponents(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllComponents", reflect.TypeOf((*MockClient)(nil).GetAllComponents), ctx)
}

// GetAppComponents mocks base method
func (m *MockClient) GetAppComponents(ctx context.Context, appID string) ([]*models.AppComponent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppComponents", ctx, appID)
	ret0, _ := ret[0].([]*models.AppComponent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppComponents indicates an expected call of GetAppComponents
func (mr *MockClientMockRecorder) GetAppComponents(ctx, appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppComponents", reflect.TypeOf((*MockClient)(nil).GetAppComponents), ctx, appID)
}

// CreateComponent mocks base method
func (m *MockClient) CreateComponent(ctx context.Context, appID string, req *models.ServiceCreateComponentRequest) (*models.AppComponent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComponent", ctx, appID, req)
	ret0, _ := ret[0].(*models.AppComponent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComponent indicates an expected call of CreateComponent
func (mr *MockClientMockRecorder) CreateComponent(ctx, appID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComponent", reflect.TypeOf((*MockClient)(nil).CreateComponent), ctx, appID, req)
}

// GetComponent mocks base method
func (m *MockClient) GetComponent(ctx context.Context, componentID string) (*models.AppComponent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComponent", ctx, componentID)
	ret0, _ := ret[0].(*models.AppComponent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComponent indicates an expected call of GetComponent
func (mr *MockClientMockRecorder) GetComponent(ctx, componentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComponent", reflect.TypeOf((*MockClient)(nil).GetComponent), ctx, componentID)
}

// UpdateComponent mocks base method
func (m *MockClient) UpdateComponent(ctx context.Context, componentID string, req *models.ServiceUpdateComponentRequest) (*models.AppComponent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComponent", ctx, componentID, req)
	ret0, _ := ret[0].(*models.AppComponent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateComponent indicates an expected call of UpdateComponent
func (mr *MockClientMockRecorder) UpdateComponent(ctx, componentID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComponent", reflect.TypeOf((*MockClient)(nil).UpdateComponent), ctx, componentID, req)
}

// DeleteComponent mocks base method
func (m *MockClient) DeleteComponent(ctx context.Context, componentID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComponent", ctx, componentID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteComponent indicates an expected call of DeleteComponent
func (mr *MockClientMockRecorder) DeleteComponent(ctx, componentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComponent", reflect.TypeOf((*MockClient)(nil).DeleteComponent), ctx, componentID)
}

// CreateTerraformModuleComponentConfig mocks base method
func (m *MockClient) CreateTerraformModuleComponentConfig(ctx context.Context, componentID string, req *models.ServiceCreateTerraformModuleComponentConfigRequest) (*models.AppTerraformModuleComponentConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTerraformModuleComponentConfig", ctx, componentID, req)
	ret0, _ := ret[0].(*models.AppTerraformModuleComponentConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTerraformModuleComponentConfig indicates an expected call of CreateTerraformModuleComponentConfig
func (mr *MockClientMockRecorder) CreateTerraformModuleComponentConfig(ctx, componentID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTerraformModuleComponentConfig", reflect.TypeOf((*MockClient)(nil).CreateTerraformModuleComponentConfig), ctx, componentID, req)
}

// CreateHelmComponentConfig mocks base method
func (m *MockClient) CreateHelmComponentConfig(ctx context.Context, componentID string, req *models.ServiceCreateHelmComponentConfigRequest) (*models.AppHelmComponentConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHelmComponentConfig", ctx, componentID, req)
	ret0, _ := ret[0].(*models.AppHelmComponentConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHelmComponentConfig indicates an expected call of CreateHelmComponentConfig
func (mr *MockClientMockRecorder) CreateHelmComponentConfig(ctx, componentID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHelmComponentConfig", reflect.TypeOf((*MockClient)(nil).CreateHelmComponentConfig), ctx, componentID, req)
}

// CreateDockerBuildComponentConfig mocks base method
func (m *MockClient) CreateDockerBuildComponentConfig(ctx context.Context, componentID string, req *models.ServiceCreateDockerBuildComponentConfigRequest) (*models.AppDockerBuildComponentConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDockerBuildComponentConfig", ctx, componentID, req)
	ret0, _ := ret[0].(*models.AppDockerBuildComponentConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDockerBuildComponentConfig indicates an expected call of CreateDockerBuildComponentConfig
func (mr *MockClientMockRecorder) CreateDockerBuildComponentConfig(ctx, componentID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDockerBuildComponentConfig", reflect.TypeOf((*MockClient)(nil).CreateDockerBuildComponentConfig), ctx, componentID, req)
}

// CreateExternalImageComponentConfig mocks base method
func (m *MockClient) CreateExternalImageComponentConfig(ctx context.Context, componentID string, req *models.ServiceCreateExternalImageComponentConfigRequest) (*models.AppExternalImageComponentConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExternalImageComponentConfig", ctx, componentID, req)
	ret0, _ := ret[0].(*models.AppExternalImageComponentConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateExternalImageComponentConfig indicates an expected call of CreateExternalImageComponentConfig
func (mr *MockClientMockRecorder) CreateExternalImageComponentConfig(ctx, componentID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExternalImageComponentConfig", reflect.TypeOf((*MockClient)(nil).CreateExternalImageComponentConfig), ctx, componentID, req)
}

// GetComponentConfigs mocks base method
func (m *MockClient) GetComponentConfigs(ctx context.Context, componentID string) ([]*models.AppComponentConfigConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComponentConfigs", ctx, componentID)
	ret0, _ := ret[0].([]*models.AppComponentConfigConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComponentConfigs indicates an expected call of GetComponentConfigs
func (mr *MockClientMockRecorder) GetComponentConfigs(ctx, componentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComponentConfigs", reflect.TypeOf((*MockClient)(nil).GetComponentConfigs), ctx, componentID)
}

// GetComponentLatestConfig mocks base method
func (m *MockClient) GetComponentLatestConfig(ctx context.Context, componentID string) (*models.AppComponentConfigConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComponentLatestConfig", ctx, componentID)
	ret0, _ := ret[0].(*models.AppComponentConfigConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComponentLatestConfig indicates an expected call of GetComponentLatestConfig
func (mr *MockClientMockRecorder) GetComponentLatestConfig(ctx, componentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComponentLatestConfig", reflect.TypeOf((*MockClient)(nil).GetComponentLatestConfig), ctx, componentID)
}

// CreateComponentBuild mocks base method
func (m *MockClient) CreateComponentBuild(ctx context.Context, componentID string, req *models.ServiceCreateComponentBuildRequest) (*models.AppComponentBuild, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComponentBuild", ctx, componentID, req)
	ret0, _ := ret[0].(*models.AppComponentBuild)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComponentBuild indicates an expected call of CreateComponentBuild
func (mr *MockClientMockRecorder) CreateComponentBuild(ctx, componentID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComponentBuild", reflect.TypeOf((*MockClient)(nil).CreateComponentBuild), ctx, componentID, req)
}

// GetComponentBuilds mocks base method
func (m *MockClient) GetComponentBuilds(ctx context.Context, componentID string) ([]*models.AppComponentBuild, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComponentBuilds", ctx, componentID)
	ret0, _ := ret[0].([]*models.AppComponentBuild)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComponentBuilds indicates an expected call of GetComponentBuilds
func (mr *MockClientMockRecorder) GetComponentBuilds(ctx, componentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComponentBuilds", reflect.TypeOf((*MockClient)(nil).GetComponentBuilds), ctx, componentID)
}

// GetComponentLatestBuild mocks base method
func (m *MockClient) GetComponentLatestBuild(ctx context.Context, componentID string) (*models.AppComponentBuild, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComponentLatestBuild", ctx, componentID)
	ret0, _ := ret[0].(*models.AppComponentBuild)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComponentLatestBuild indicates an expected call of GetComponentLatestBuild
func (mr *MockClientMockRecorder) GetComponentLatestBuild(ctx, componentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComponentLatestBuild", reflect.TypeOf((*MockClient)(nil).GetComponentLatestBuild), ctx, componentID)
}

// GetComponentBuild mocks base method
func (m *MockClient) GetComponentBuild(ctx context.Context, componentID, buildID string) (*models.AppComponentBuild, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComponentBuild", ctx, componentID, buildID)
	ret0, _ := ret[0].(*models.AppComponentBuild)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComponentBuild indicates an expected call of GetComponentBuild
func (mr *MockClientMockRecorder) GetComponentBuild(ctx, componentID, buildID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComponentBuild", reflect.TypeOf((*MockClient)(nil).GetComponentBuild), ctx, componentID, buildID)
}

// GetComponentBuildLogs mocks base method
func (m *MockClient) GetComponentBuildLogs(ctx context.Context, componentID, buildID string) ([]models.ServiceBuildLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComponentBuildLogs", ctx, componentID, buildID)
	ret0, _ := ret[0].([]models.ServiceBuildLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComponentBuildLogs indicates an expected call of GetComponentBuildLogs
func (mr *MockClientMockRecorder) GetComponentBuildLogs(ctx, componentID, buildID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComponentBuildLogs", reflect.TypeOf((*MockClient)(nil).GetComponentBuildLogs), ctx, componentID, buildID)
}

// GetAppReleases mocks base method
func (m *MockClient) GetAppReleases(ctx context.Context, appID string) ([]*models.AppComponentRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppReleases", ctx, appID)
	ret0, _ := ret[0].([]*models.AppComponentRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppReleases indicates an expected call of GetAppReleases
func (mr *MockClientMockRecorder) GetAppReleases(ctx, appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppReleases", reflect.TypeOf((*MockClient)(nil).GetAppReleases), ctx, appID)
}

// GetComponentReleases mocks base method
func (m *MockClient) GetComponentReleases(ctx context.Context, componentID string) ([]*models.AppComponentRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComponentReleases", ctx, componentID)
	ret0, _ := ret[0].([]*models.AppComponentRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComponentReleases indicates an expected call of GetComponentReleases
func (mr *MockClientMockRecorder) GetComponentReleases(ctx, componentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComponentReleases", reflect.TypeOf((*MockClient)(nil).GetComponentReleases), ctx, componentID)
}

// CreateComponentRelease mocks base method
func (m *MockClient) CreateComponentRelease(ctx context.Context, componentID string, req *models.ServiceCreateComponentReleaseRequest) (*models.AppComponentRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComponentRelease", ctx, componentID, req)
	ret0, _ := ret[0].(*models.AppComponentRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComponentRelease indicates an expected call of CreateComponentRelease
func (mr *MockClientMockRecorder) CreateComponentRelease(ctx, componentID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComponentRelease", reflect.TypeOf((*MockClient)(nil).CreateComponentRelease), ctx, componentID, req)
}

// GetRelease mocks base method
func (m *MockClient) GetRelease(ctx context.Context, releaseID string) (*models.AppComponentRelease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRelease", ctx, releaseID)
	ret0, _ := ret[0].(*models.AppComponentRelease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRelease indicates an expected call of GetRelease
func (mr *MockClientMockRecorder) GetRelease(ctx, releaseID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelease", reflect.TypeOf((*MockClient)(nil).GetRelease), ctx, releaseID)
}

// GetReleaseSteps mocks base method
func (m *MockClient) GetReleaseSteps(ctx context.Context, releaseID string) ([]*models.AppComponentReleaseStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReleaseSteps", ctx, releaseID)
	ret0, _ := ret[0].([]*models.AppComponentReleaseStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReleaseSteps indicates an expected call of GetReleaseSteps
func (mr *MockClientMockRecorder) GetReleaseSteps(ctx, releaseID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReleaseSteps", reflect.TypeOf((*MockClient)(nil).GetReleaseSteps), ctx, releaseID)
}
