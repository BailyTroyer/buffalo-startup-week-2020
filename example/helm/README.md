# Bflobox API

## TL;DR;

```console
$ helm install ./example/helm
```

## Introduction

This chart bootstraps a Bflobox API deployment on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.3+ with Beta APIs enabled

## Installing the Chart

To install the chart with the release name `my-release`:

```console
helm install --name my-release ./example/helm
```

The command deploys the Bflobox API on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list --all-namespaces`

## Uninstalling the Chart

To uninstall/delete the `my-release` deployment:

```console
helm delelete my-release
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the Bflobox API chart and their default values.

Parameter | Description | Default
--------- | ----------- | -------
`meta.provider` | cloud provider the cluster is running on | `aws`
`meta.region` | kubernetes cluster region | `us-east-1`
`meta.clusterName` | kubernetes cluster name | `useast1.dev.bflobox.com`
`meta.name` | main application name | `bflobox-api`
`meta.namespace` | deployed namespace | `bflobox-api`
`meta.environment` | deployed namespace | `bflobox-api`
`containers.bfloBoxApi.name` | main container name | `bflobox-api`
`containers.bfloBoxApi.image.repository` | main container image repo | `ECR_REPO_HERE`
`containers.bfloBoxApi.image.tag` | main container image tag | `CHANGE_ME`
`containers.bfloBoxApi.port` | main pod's port | `8080`
`containers.bfloBoxApi.replicas` | number of main pod replicas | `1`
`containers.bfloBoxApi.image.resources` | main pod resource requests & limits | `{}`
`loadBalancer.cert` | AWS ELB Cert | `arn:aws:acm:us-east-1:CHANGE_ME:certificate/CHANGE_ME`