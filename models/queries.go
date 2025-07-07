package models

type GetActionWorkflowsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetInstallActionWorkflowRecentRunsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetAppConfigsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetAppInputConfigsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetAppRunnerConfigsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetAppSandboxConfigsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetAppSecretsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetAppsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetComponentBuildsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetAllComponentsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetAppComponentsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetComponentConfigsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetInstallComponentsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetInstallComponentDeploysQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetInstallDeploysQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetInstallInputsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetInstallSandboxRunsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetInstallersQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetOrgInvitesQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetOrgsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetReleaseStepsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetAppReleasesQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetComponentReleasesQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetVCSConnectionsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetAllVCSConnectedReposQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetAppInstallsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetAllInstallsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetInstallWorkflowsQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}

type GetPaginatedQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}
