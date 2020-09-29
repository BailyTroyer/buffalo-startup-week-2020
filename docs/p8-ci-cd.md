# Part 8 (CI/CD ft. Github Actions)

## Overview

I won't sell CI/CD to you because I have faith you understand its benefits and what it really provides for not only a startup but any development team. What I'm going to briefly touch upon is creating a pipeline in Github Actions.

Note you can find the example actions in `.github/workflows`; they run on PRs and merges to `master`

## Github Actions

I've broken this repo's CI/CD into two pipeline files: `master-ci-cd.yaml` and `pull-request.yaml`. One runs on merge to master, which updates the S3 buckets exposed through CloudFront, and our helm chart which is running live on a kubernetes cluster. 

If you take a look at the [master-ci-cd.yaml](../.github/workflows/master-ci-cd.yaml) you'll see a common theme in the pipeline:

* deploy presentation webapp
    * build static assets
    * push static assets
* deploy bflobox webapp
    * build image
    * build statoc assets
    * push static assets
* deploy API
    * build image
    * push image to ECR
    * helm upgrade

If you take a look at the [pull-request.yaml](../.github/workflows/pull-request.yaml) you'll see a common theme in the pipeline:

* lint & test web
    * build image
    * lint codebase
    * unit test codebase
* lint and test API
    * build image
    * lint codebase
    * unit test codebase

This is a common theme, and is often best practice when creating a CI/CD pipeline, where you build, lint, test and then once merged, push, deploy and upgrade the service. In production grade pipelines, you would also want to include monitoring to maybe a slack channel and add some post-deploy smoketests.

**Note:** More information and documentation on interacting with and building Github actions can be found [here](https://docs.github.com/en/free-pro-team@latest/actions).
