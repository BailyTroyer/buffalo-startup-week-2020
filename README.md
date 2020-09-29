# Buffalo Startup Week 2020

## From localhost to Prod

### 0.5 Preface

The following mini-series is the text version of a local tech talk held for Buffalo's Startup week. This repo will contain all examples used in the talk, thus, feel free to use these for your own service too!

### 1. Overview

You just made your billion dollar startup idea but you only have it running on localhost. Where do you go from there? Do you pay monthly for Heroku or Netlify? What if you could manage it all on your own while saving monies? 

This repo will go through deploying an API and React App running on localhost to a live running .com domain. We’ll use tools including Docker, Make, Github Actions, ECR, Route53, Cloudfront and more!

### 2. Buffalo Box

Before we get our hands dirty let's go through what we're building exactly...

Alrighty - let's say you're a new Buffalo based startup called "BuffBox." You're a next generation subscription service that ships out Buffalo themed swag, food and more! Sweet :chicken-wing: :buffalo:

You and your dev team just created a sweet website in React, and a backend API in Golang. It all works on `localhost:3000` and `localhost:8080` however your customers can't view anything ... because bflobox.com doesn't exist!

So we need to somehow ship our code to a **live** environment that _real_ users can interact and use 24-7. That's what we're going to do in this repo. 

### 3. The Process

**1. (API)** First let's make sure we're using version control: I personally love Github. After that we need to easily run our codebase from anywhere; and I mean anywhere. On Linux, Ubuntu, Centos, MacOS, or your grandmas Windows XP laptop XD. To do so we'll use a fancy containerization tool called Docker. You've probably heard of it, and if you haven't that's ok (I guess patrick wasn't the only one living under a rock).

**2. (API)**  Once we have our services containerized, we're going to want to store our image remotely so we can **run** it in a production environment later!. For that we'll use ECR to privately store our bundled application.

**3. (API)** The next series of steps sort of fall in place in parallel. We're going to dive into AWS and see how we can run our docker image in a Kubernetes **cluster**. Similarlly we're going to want to easily access this service using our domain name we bought: bflobox.com. To do so we're going to interact with some AWS services: Route53 & ELB. More on that later.

**4. (API)** At this point, **a.** we've built our image **b.** stored it remotely **c.** ran it remotely in a cluster exposed by an elastic load balancer which sits behind our public facing DNS entry for `bflobox.com`. Woah!

**5. (API)** What happens next? What if we push a new update to the site? Do we have to manually go through that process again?? What if we could easily "do all the things" but only when we push code to our trunk in Github (i.e. the `main/master` branch)? We'll use Github Actions, Github's version of CI/CD.

**6. (Web)** Great, our API is running at api.bflobox.com. What about our website? Since we want to build assets and serve them like any other website, we're going to use S3 as our storage with CloudFront as our delivery method. 

### Off Topic

Shout out to postman, tableplus, figma, Spotify, and Zsh for being awesome. 
