# Part 3 (AWS Overview)

## Intro

This probably seems pre-mature, however, we're going to need to talk about it sooner-or-later. AWS a.k.a Amazon Web Services.

You've probably used Amazon, and you've probably heard of the internet right? We'll roughly 40% of it runs on AWS [reference](https://www.theverge.com/2018/7/28/17622792/plugin-use-the-internet-without-the-amazon-cloud).

Today, right now, in a few minutes we're going to use AWS to create some resources to run our API and website. And its gonna be cheap, really cheap. And its gonna be freakin' awesome. Alright let's go!

## Create An Account & Login

It comes as no surprise, you're going to need to create an account. You can start [here](https://aws.amazon.com/premiumsupport/knowledge-center/create-and-activate-aws-account/). Make sure you turn on billing alerts, just for the peace-of-mind; I know I have. 

You're also going to want an "Admin" IAM user you can use at the CLI level that has access to tools we're going to use. Its' not the best but for the sake of time feel free to create a user called admin that has the following permissions:

* AmazonEC2FullAccess
* IAMFullAccess
* AmazonEC2ContainerRegistryFullAccess
* AmazonS3FullAccess
* AmazonVPCFullAccess
* AmazonRoute53FullAccess

After that you should have an access and secret key you can use to confugre the AWS CLI. 

## AWS CLI

We're going to use the AWS CLI more often than the console since its quick, easy and much easier to show in a markdown file than tons of screenshots. You can easily configure the AWS CLI [here](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html).

For my examples I have an `admin` user in my `~/.aws/credentials` file that has access to all the permissions I listed above. 

## What's next?

We just built our API and pushed it to ECR. That's great progress but there's much more to be done. Let's quickly go over what we have left to go over to get our stuff up-and-running live.

1. Run our API in k8s, exposing our service running on port `8080` behind an AWS ELB
2. Register our domain, create a hosted zone, and generate a TLS cert
3. Back-track, and get our website stored in an S3 bucket, hosted on a global CDN using CloudFront.
