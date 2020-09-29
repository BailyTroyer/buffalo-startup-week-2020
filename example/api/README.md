# BfloBox API

The following service is our backend API that hosts the meat and potatoes of all business functionality. Aside from our website this is the main heart of any company, where all functionality and SEVs originate. 

## Getting Started

### Configuration (Required)

The API uses a config yaml file `config.yaml` which gets parsed and unmarshalled in [config.go](internal/util/config.go). This also gets mounted in the container in the `docker-compose.yaml`. Currently the service name and port are configured, however feel free to add any additional configurations to suit your needs. In its current setup, the following values are required to setup the app name, port and new relic license. This would normally get mounted in your Deployment from a barebones Kubernetes secret and/or Vault. 

```
meta:
  name: "CHANGE_ME"
  port: 8080
  env: "CHANGE_ME"
```

### Running Locally

The following Make directives will build, run and lint your current state of the API. Make sure linting and testing pass
before making a PR, this will save both your time and the reviewers.

* `make build` - Build the Docker image
* `make run` - Run the Docker container housing the API
* `make lint` - Lint all code
* `make test` - Run all necessary unit and integration tests
* `make coverage-html` - generate codecoverage html file
* `make coverage-percent` - generate codecoverage percent

# Structure

The following map outlines the generic structure of the API. With all "business logic" nested within `internal/handler/`, models in `internal/model/`, and all utils (HTTP mock client, logging, config) inside `internal/util`.

```
├── internal
│   ├── handler
│   │   ├── app.go              // Core API handlers
│   │   ├── health.go           // Generic Server health
│   ├── model                   // Our API core handlers
│   │   └── model.go            // Models for application
│   └── util
│       └── response.go         // Common response functions
│       └── logging.go          // Zap logging config
│       └── mock.go             // All purpose HTTP client mock
│       └── config.go           // Config loader
└── cmd
    └── main.go                 // Main API entrypoint
```

# API Endpoints

The following routes are currently in-use. You can also build and view the swagger docs locally by running `make swagger`.

HTTP request | Description
------------- | -------------
**GET** /health | Server Health | 

# Contributing

You may contribute in several ways by creating/proposing new features, fixing bugs, improving documentation and adding more to the README! Find more information in [CONTRIBUTING.md](docs/CONTRIBUTING.md).
