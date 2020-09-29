# Part 5 (KOPS++)

## Overview

Now that we have our API running in a Docker environment, with an image stored in ECR we need to run that somewhere. For sake of time and monies, we're going to evaluate [KOPS, or, Kubernetes Operations](https://github.com/kubernetes/kops). Its a very very helpful tool used to spinup Kubernetes clusters without using a managed service like GKE or EKS. Why would we not just use EKS? Mainly becuase `You pay $0.10 per hour for each Amazon EKS cluster that you create` which is absolutely disappointing since that would be roughly $70/month for just operating a cluster with no nodes... that's the price of two t2.medium nodes :(

So to save time and money we can easily create a k8s cluster with a few commands and get up and running in no time. 

## Install & Configure our KOPS Cluster

More information on what we're about to do can be found [here](https://kubernetes.io/docs/setup/production-environment/tools/kops/).

You can easily install kops using brew:

```
$ brew update && brew install kops
```

Once installed let's make a Route53 hosted zone for our cluster. I used `dev.bflobox.com` but feel free to use whatever you prefer. You can easily create a hosted zone in Route53 by selecting **Create Hosted Zone** and entering the required fields. 

With our hosted zone configured, lets create an S3 bucket to store our cluster configurations generated and managed by KOPS. I used the domain name `clusters.dev.bflobox.com`, but, it's a free country use whatever you prefer!

Create an S3 bucket by going into the S3 console, selecting **Create Bucket** and entering in your bucket name. Feel free to use all the default configurtions since we want this bucket to be private. 

With our S3 bucket created export the following environment variable for KOPS and create the cluster like below.

It will by default create one master node based on a t3.medium and two worker nodes that are also t3.mediums. These nodes cost roughly $30/month, so feel free to use small or micros if you prefer. I used a small since this example service won't get tons of traffic. 

```
$ export KOPS_STATE_STORE=s3://clusters.dev.bflobox.com
$ kops create cluster --zones=us-east-1c useast1.bflobox.example.com

// run edit if you want to change the node type to something cheaper :)
$ kops edit ig --name=useast1.dev.example.com nodes

$ kops update cluster useast1.bflobox.example.com --yes
```

Once created you can run the following and go grab a coffee while everything spins up:

```
$ kops validate cluster --wait 10m
...
NODE STATUS
NAME				ROLE	READY
ip-{HIDDEN}.ec2.internal	node	True
ip-{HIDDEN}.ec2.internal	node	True
ip-{HIDDEN}.ec2.internal	master	True

Your cluster useast1.dev.bflobox.com is ready
```

Sick! We have a Kubernetes cluster now!!!
