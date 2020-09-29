# Part 6 (Web App)

## Overview

I won't lie, our happy little webapp hasn't had as much love as you probably imagined. After all its the one customer facing piece of software your users care about. However, quite frankly its easy as sh** to deploy a webapp in 2020. Like, **really really** easy. And I'm not talking about using Heroku or Netlify, we're going to use AWS and its going to be great. Let's go!

## Intro To Codebase

Let's first head over to the webapp in [example/bflobox](../example/bflobox). This is a simple React app that runs on `localhost:3000` when running the dev server.

Assuming you have yarn installed locally (if not you can [here](https://classic.yarnpkg.com/en/docs/install/#mac-stable)) go ahead and start the app like below. Note once you run `yarn start` your browser should pop up with localhost automatically. 

```
$ yarn install
$ yarn start
```

:tada:

Welcome to the frontend of our new startup ;) It's not much but once we figure out how to deploy and setup CI/CD you can just run with it!

## AWS CloudFront

So before we were running a dev server that automatically bundles and runs your application with hot-reload enabled. This is **not** how we're going to deploy to production. Instead we are going to bundle (more information on that [here](https://reactjs.org/docs/code-splitting.html)) which essentially turns all of your react-specific javascript code into viewable HTML that users can view without having React installed. It's really nice.

Thankfully `create-react-app` (the tool used to bootstrap this project) comes with a fancy yarn command to bundle the static assets that we need to "put somewhere." Just run `yarn build` and you should see a new directory `build` with an `index.html` inside. This is what we're going to deploy to production.

To do so, we're going to push the following build directory to S3, and then host it behind a CDN (Content Delivery Network). More information on CDNs can be found [here](https://www.cloudflare.com/learning/cdn/what-is-a-cdn/).

First we need an S3 bucket, which can be done really simply by following this doc. After that all we need to do is copy our built assets and push them up to our bucket like so. 

**Note:** In this case our s3 bucket is called www.bflobox.com.

```
$ aws s3 sync ./build s3://www.bflobox.com --cache-control max-age=30
```

You're going to want to create a bucket policy with the following information to allow anonymous read acess since this is a public website:

```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "PublicReadGetObject",
            "Effect": "Allow",
            "Principal": "*",
            "Action": "s3:GetObject",
            "Resource": "arn:aws:s3:::{S3_BUCKET_HERE}/*"
        }
    ]
}
```

From there all we need to do is create a CDN in CloudFront, which can be easily done in the console under **CloudFront > Create Distribution > Web (Get Started)**

From there you will need to configure the following parts:

* Origin Domain Name - 
* Origin Path - 
* Viewer Protocol Policy > Redirect HTTP to HTTPS -
* Distribution Settings > Price Class - I only used US, Canada & Europe since its cheaper
* Alternate Domain Names (CNAMEs) - for this I used www.bflobox.com & bflobox.com since not everyone uses www anymore :(
* SSL Certificate > Custom SSL Certificate > Select your cert we made in p4-dns

Then select **Create Distribution** ... 

From there your distribution will have a domain name that goes directly to our webapp example that we pushed to s3. The last final touch would be creating a `CNAME` Route53 entry for www.bflobox.com and bflobox.com that maps to our distributon: `fdsaffdafds.cloudfront.net`



```
$ aws cloudfront create-distribution \
    --origin-domain-name awsexamplebucket.s3.amazonaws.com \
    --default-root-object index.html
```
