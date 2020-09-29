# Part 7 (Kubectl & Helm)

## Overview

As of right now, we have our website running in a global CDN, and our API bunded and store in ECR. Time to meet Kubernetes, one of the largest and biggest open source projects in the world that has blown up the industry since July 21, 2015. I'm not sure if there's a metric on % of the internet running on k8s but I bet its a sizeable amount (except for legacy boomer companies and Government related contractors/agencies).

## Kubernetes

Before in [Section 5](p5-kops.md) we created a k8s cluster with KOPS, and just left it hanging. Now we get to play and interact using a very handy CLI: `kubectl` ([link](https://kubernetes.io/docs/tasks/tools/install-kubectl/) to install).

I won't go into **too** much depth on Kubernetes, kubectl and its resources since we are limited for time however here we go:

So previously we created a Docker container that runs our application bundled inside its own "virtual environment" with its own operating system and dependencies pre-packaged and ready-to-use. Now that's great and all, but docker alone doesn't really solve our problem in wanting to host a highly available, fault tolerant service. With Kubernetes we get several benefits on top of just running locally:

* A higher layer of abstraction to represent large meshes of microservices running in user-defined spaces, or, namespaces
* Automatic rolling deployments and straight forward configurations that allow us to specify a plethora of settings from number of replicas, health probes, secrets information, volume mounted data, and much much more
* A very large living open source community with thousands of articles, documentation and even companies that specialize in handling Kubernetes for you if you can foot the $$

Now aside from the business pitch in Kuberenetes, let's talk more specifics in what we're going to do with our simple API example. 

We're going to deploy a `Deployment` kubernetes resource with a `Service` resource exposing the Deployment's pod's port 8080 through an Elastic Load Balancer, or, ELB which we can then CNAME in route53 and expose to the rest of the world; yes that was a mouth full so let's digest what I just said. 

I didn't mention this before, but Kubernetes is made of a bunch of "resources" that specify the state of the world. The smallest `unit` in Kubernetes would be a Pod which is essentially your docker container but with a few bells and whistles. From there you have a `Deployment | StatefulSet | DaemonSet | CronJob | Job` which all wrap the Pod resource. There's tons of information on the internet explaining the use-case for each (Google it on your own) but for this example we need a Deployment. Why? Because our service can be easily replicated horizontally across multiple nodes and AZs and we don't care about any state since its just an API (any real state would exist in a DB not on disk in a volume). 

Even that above was a lot to take in, so let's take a brief tour of the cluster you should have made in [Step 5](p5-kops.md)

One last thing, as I mentioned above namespaces are how we group anything and everything in Kubernetes. You could group it by domains in a SOA, or just have one main namespace for everything (**not recommended**).

```
// list namespaces in the cluster
$ kubectl get namespaces
bflobox           Active   30h
default           Active   30h
kube-node-lease   Active   30h
kube-public       Active   30h
kube-system       Active   30h

// get "everything in the namespace"
$ kubectl get all --namespace bflobox
NAME                               READY   STATUS    RESTARTS   AGE
pod/bflobox-api-8458b56d8f-skrzz   1/1     Running   0          29m

NAME                  TYPE           CLUSTER-IP       EXTERNAL-IP                                                               PORT(S)                      AGE
service/bflobox-api   LoadBalancer   100.71.182.220   fdsafd-fdsafds.us-east-1.elb.amazonaws.com   443:30294/TCP,80:31135/TCP   3h27m

NAME                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/bflobox-api   1/1     1            1           3h27m

NAME                                     DESIRED   CURRENT   READY   AGE
replicaset.apps/bflobox-api-8458b56d8f   1         1         1       120m

// get only pods
$ kubectl get pods --namespace bflobox
NAME                           READY   STATUS    RESTARTS   AGE
bflobox-api-8458b56d8f-skrzz   1/1     Running   0          30m
```

Right there we looked at all namespaces in the cluster, then all resources in our bflobox namespace, then did some filtering to just see the pods in our namespace. Remember, a pod is a runner container of the image we pushed to ECR.

In the Kubernetes world all of these resources are made up of `yaml` files that specify anything and everything about itself. From the ports it exposes, to the labels and annotations set, to the name and namespace its deployed to, to container level information like environment variables, secrets, the image URL, and much more. 

## Helm

Without going too deep in the kubernetes resources induvidually I figured it'd be easier to just go straight to Helm. Think of this as the package manager for Kubernetes, where you can easily install a whole bundle of resources (which are just plain yaml files) in one command. The structure is really simple `helm install {CHART_PATH} --name {NAME} --namespace {NAMESPACE}` and boom you can view, create and delete hundreds of yaml configurations with no sweat. 

It's a life saver whether you believe it or not. 

If you head over to [example/helm](../example/helm) you'll see the helm chart I've created for our bflobox API. It has a `Deployment` and `Service` which are our kubernetes resources we'll use to deploy and expose the API running on `:8080`.

You can easily install the chart by running the following command which will simply install everything for you.

```
$ helm upgrade --install bflobox-api example/helm --namespace bflobox --set containers.bfloBoxApi.image.tag={ECR_IMAGE_TAG_HERE} --wait
```

## [[Next Section] (Part 8: CI/CD)](p8-ci-cd.md)