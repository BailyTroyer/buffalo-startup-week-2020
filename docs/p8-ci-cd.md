# Part 7 (CI/CD ft. Github Actions)

## Overview

Talk about CI/CD and how its so important. Talk about merging to master and updating the stuff we just deployed automatically....

## Github Actions

Talk about creating a pipeline in `.github/pipelines` directory, link out to examples and such

## API Pipeline

1. build (PR)
2. lint & test (PR)
3. push (master)
4. Update ECS (master)
5. (Optional) update slack channel, etc

## Webapp Pipeline

1. build (PR)
2. lint & test (PR)
3. build static assets (master)
4. push to ECR (master)
5. (Optional) update slack channel, etc
