# nuon-go

An interface for working with our API from go.

## Overview

The Nuon API allows you to configure your Nuon apps, release to them and create new installs.

Full documentation is available at https://docs.nuon.co.

All endpoints in the API follow REST conventions and standard HTTP methods. You can find the OpenAPI Spec [here](https://ctl.prod.nuon.co/docs/doc.json)

## Installation

In your project, you can install the package directly using `go get`:

```bash
go get github.com/nuonco/nuon-go
```

In your code, add the following import:

```
import nuon "github.com/nuonco/nuon-go"
```

## Create a client

Create a new api client, using an API key set in the environment.
```go

apiURL := "https://ctl.prod.nuon.co"
apiToken := os.Getenv("NUON_API_TOKEN")
orgID := os.Getenv("NUON_ORG_ID")

apiClient, err := client.New(s.v,
  client.WithAuthToken(apiToken),
  client.WithURL(apiURL),
  client.WithOrgID(orgID),
)
if err != nil {
  return fmt.Errorf("unable to get api client: %w", err)
}
```

## Example usage

### List an app

```go
apps, err := apiClient.ListApps(ctx)
```

### Get an app

```go
app, err := apiClient.GetApp(ctx, appID)
```

## Contributing

This repo is currently not setup for local contributions. However, if you would like to contribute you can open an issue or contact us on our [shared slack](https://join.slack.com/t/nuoncommunity/shared_invite/zt-1q323vw9z-C8ztRP~HfWjZx6AXi50VRA).

