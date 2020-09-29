# BfloBox Web

The following service is our frontend webapp. This is what our customers use and interact with, and this represents the face of our brand. This was made with create-react-app and I highly recommend using react and/or GatsbyJS depending on your usecase. You can learn more in the [Create React App documentation](https://facebook.github.io/create-react-app/docs/getting-started)

## Getting Started

### Configuration (Required)

The API uses a .env file `.env.production` which gets mounted in the Docker container during linting, testing and bundling. This also gets mounted in the container in the `docker-compose.yaml`. Feel free to add any additional configurations to suit your needs from API URLs to anything configurable per environment. With a real website, you'd want at least a development, staging and production environment. 

### Running Locally

The following Make directives will build, run and lint your current state of the API. Make sure linting and testing pass
before making a PR, this will save both your time and the reviewers.

* `make build` - Build the Docker image
* `make run` - Run the Docker container in dev-mode housing the webapp
* `make lint` - Lint all code
* `make test` - Run all necessary unit and integration tests

# Contributing

You may contribute in several ways by creating/proposing new features, fixing bugs, improving documentation and adding more to the README! Find more information in [CONTRIBUTING.md](docs/CONTRIBUTING.md).
