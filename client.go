package nuon

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-playground/validator/v10"
	genclient "github.com/nuonco/nuon-go/client"
	"github.com/nuonco/nuon-go/models"
)

//
//go:generate -command swagger go run github.com/go-swagger/go-swagger/cmd/swagger
//go:generate swagger generate client --skip-tag-packages -f $NUON_API_URL/docs/doc.json
//go:generate -command mockgen go run github.com/golang/mock/mockgen
//go:generate mockgen -destination=mock.go -source=client.go -package=nuon
type Client interface {
	SetOrgID(orgID string)

	//get / create org
	GetOrgs(ctx context.Context) ([]*models.AppOrg, error)
	CreateOrg(ctx context.Context, req *models.ServiceCreateOrgRequest) (*models.AppOrg, error)

	//current org
	GetOrg(ctx context.Context) (*models.AppOrg, error)
	UpdateOrg(ctx context.Context, req *models.ServiceUpdateOrgRequest) (*models.AppOrg, error)
	DeleteOrg(ctx context.Context) (bool, error)
	CreateOrgUser(ctx context.Context, req *models.ServiceCreateOrgUserRequest) (*models.AppUserOrg, error)

	// internal methods
	GetApp(ctx context.Context, appID string) (*models.AppApp, error)
	GetApps(ctx context.Context) ([]*models.AppApp, error)
	CreateApp(ctx context.Context, req *models.ServiceCreateAppRequest) (*models.AppApp, error)
	UpdateApp(ctx context.Context, appID string, req *models.ServiceUpdateAppRequest) (*models.AppApp, error)
	DeleteApp(ctx context.Context, appID string) (bool, error)
	UpdateAppSandbox(ctx context.Context, appID string, req *models.ServiceUpdateAppSandboxRequest) (*models.AppApp, error)

	// general methods
	GetCurrentUser(ctx context.Context) (*models.AppUserToken, error)
	PublishMetrics(ctx context.Context, req []*models.ServicePublishMetricInput) error

	// sandbox methods
	GetSandboxes(ctx context.Context) ([]*models.AppSandbox, error)
	GetSandbox(ctx context.Context, sandboxID string) (*models.AppSandbox, error)
	GetSandboxReleases(ctx context.Context, sandboxID string) ([]*models.AppSandboxRelease, error)

	// vcs connections
	CreateVCSConnection(ctx context.Context, req *models.ServiceCreateConnectionRequest) (*models.AppVCSConnection, error)
	GetVCSConnections(ctx context.Context) ([]*models.AppVCSConnection, error)
	GetVCSConnection(ctx context.Context, connID string) (*models.AppVCSConnection, error)
	GetAllVCSConnectedRepos(ctx context.Context) ([]*models.ServiceRepository, error)

	// installs
	CreateInstall(ctx context.Context, appID string, req *models.ServiceCreateInstallRequest) (*models.AppInstall, error)
	GetAppInstalls(ctx context.Context, appID string) ([]*models.AppInstall, error)
	GetAllInstalls(ctx context.Context) ([]*models.AppInstall, error)

	GetInstall(ctx context.Context, installID string) (*models.AppInstall, error)
	UpdateInstall(ctx context.Context, installID string, req *models.ServiceUpdateInstallRequest) (*models.AppInstall, error)
	DeleteInstall(ctx context.Context, installID string) (bool, error)

	// install deploys
	GetInstallDeploys(ctx context.Context, installID string) ([]*models.AppInstallDeploy, error)
	CreateInstallDeploy(ctx context.Context, installID string, req *models.ServiceCreateInstallDeployRequest) (*models.AppInstallDeploy, error)
	GetInstallDeploy(ctx context.Context, installID, deployID string) (*models.AppInstallDeploy, error)
	GetInstallLatestDeploy(ctx context.Context, installID string) (*models.AppInstallDeploy, error)
	GetInstallDeployLogs(ctx context.Context, installID, deployID string) ([]models.ServiceDeployLog, error)

	// install components
	GetInstallComponents(ctx context.Context, installID string) ([]*models.AppInstallComponent, error)
	GetInstallComponentDeploys(ctx context.Context, installID, componentID string) ([]*models.AppInstallDeploy, error)
	GetInstallComponentLatestDeploy(ctx context.Context, installID, componentID string) (*models.AppInstallDeploy, error)

	// components
	GetAllComponents(ctx context.Context) ([]*models.AppComponent, error)
	GetAppComponents(ctx context.Context, appID string) ([]*models.AppComponent, error)
	CreateComponent(ctx context.Context, appID string, req *models.ServiceCreateComponentRequest) (*models.AppComponent, error)

	GetComponent(ctx context.Context, componentID string) (*models.AppComponent, error)
	UpdateComponent(ctx context.Context, componentID string, req *models.ServiceUpdateComponentRequest) (*models.AppComponent, error)
	DeleteComponent(ctx context.Context, componentID string) (bool, error)

	// component configs
	CreateTerraformModuleComponentConfig(ctx context.Context, componentID string, req *models.ServiceCreateTerraformModuleComponentConfigRequest) (*models.AppTerraformModuleComponentConfig, error)
	CreateHelmComponentConfig(ctx context.Context, componentID string, req *models.ServiceCreateHelmComponentConfigRequest) (*models.AppHelmComponentConfig, error)
	CreateDockerBuildComponentConfig(ctx context.Context, componentID string, req *models.ServiceCreateDockerBuildComponentConfigRequest) (*models.AppDockerBuildComponentConfig, error)
	CreateExternalImageComponentConfig(ctx context.Context, componentID string, req *models.ServiceCreateExternalImageComponentConfigRequest) (*models.AppExternalImageComponentConfig, error)
	GetComponentConfigs(ctx context.Context, componentID string) ([]*models.AppComponentConfigConnection, error)
	GetComponentLatestConfig(ctx context.Context, componentID string) (*models.AppComponentConfigConnection, error)

	// builds
	CreateComponentBuild(ctx context.Context, componentID string, req *models.ServiceCreateComponentBuildRequest) (*models.AppComponentBuild, error)
	GetComponentBuilds(ctx context.Context, componentID string) ([]*models.AppComponentBuild, error)
	GetComponentLatestBuild(ctx context.Context, componentID string) (*models.AppComponentBuild, error)
	GetComponentBuild(ctx context.Context, componentID, buildID string) (*models.AppComponentBuild, error)
	GetComponentBuildLogs(ctx context.Context, componentID, buildID string) ([]models.ServiceBuildLog, error)

	// releases
	GetAppReleases(ctx context.Context, appID string) ([]*models.AppComponentRelease, error)
	GetComponentReleases(ctx context.Context, componentID string) ([]*models.AppComponentRelease, error)
	CreateComponentRelease(ctx context.Context, componentID string, req *models.ServiceCreateComponentReleaseRequest) (*models.AppComponentRelease, error)

	GetRelease(ctx context.Context, releaseID string) (*models.AppComponentRelease, error)
	GetReleaseSteps(ctx context.Context, releaseID string) ([]*models.AppComponentReleaseStep, error)
}

var _ Client = (*client)(nil)

type client struct {
	v *validator.Validate

	APIToken string `validate:"required"`
	APIURL   string `validate:"required"`
	OrgID    string

	genClient    *genclient.NuonAPI
	appTransport *appTransport
}

type clientOption func(*client) error

func New(v *validator.Validate, opts ...clientOption) (*client, error) {
	c := &client{v: v}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	if err := c.v.Struct(c); err != nil {
		return nil, err
	}

	apiURL, err := url.Parse(c.APIURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse api url: %w", err)
	}

	transport := httptransport.New(apiURL.Host, "", []string{apiURL.Scheme})
	appTransport := &appTransport{
		authToken: c.APIToken,
		orgID:     c.OrgID,
		transport: http.DefaultTransport,
	}
	c.appTransport = appTransport
	transport.Transport = appTransport
	genClient := genclient.New(transport, nil)
	c.genClient = genClient

	return c, nil
}

// WithAuthToken specifies the auth token to use
func WithAuthToken(token string) clientOption {
	return func(c *client) error {
		c.APIToken = token
		return nil
	}
}

// WithURL specifies the url to use
func WithURL(url string) clientOption {
	return func(c *client) error {
		c.APIURL = url
		return nil
	}
}

// WithOrgID specifies the org id to use
func WithOrgID(orgID string) clientOption {
	return func(c *client) error {
		c.OrgID = orgID
		return nil
	}
}
