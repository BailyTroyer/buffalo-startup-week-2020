# Part 2 (Docker)

## Intro 

Now that we're on the same page about git and some best practices let's dive into getting our stuff hosted somewhere!

If you haven't already noticed the `example/` directory at the root of this repo take a chance to get slightly familiar with what we're deploying. You don't have to understand what language/framework they're using, but be aware that we have **both** an API and a webapp. One has a dev server that runs on `localhost:3000` and the other runs regularly on `localhost:8080`. Note the webapp will __not__ run on a port in production, rather we'll build static assets and ship them to a remote bucket that we'll host on a CDN, or, content delivery network. More on that later...

## Building With Docker

Let's first start with the API located in [example/api](../example/api). This is a simple Go API that runs on localhost:8080. If you're not familiar with Go, you probably don't have it installed and that's more of a reason why we're using Docker! Not only does it make deploying easier, but it also makes developer onboarding much simpler since you don't have to worry about dev dependencies, you just need Docker installed (docs on installing docker can be found [here](https://docs.docker.com/get-docker/)).

1st we're going to need a `Dockerfile` to containerize and "bundle" our service and all of its dependencies. Think of Docker as a Virtual Machine that runs its own operating system separate from your host machine (technically containerization uses the same kernel as the host, but we won't talk about that here).

The basic "commands" or keywords you'll want to know and understand are:

1. `FROM` - this is the base image you're Dockerfile is based off-of. There's an "app store" full of nifty base images you can use on Dockerhub that have most of the heavy lifting already done for you. Common examples are golang, python, alpine, node and much much more!

2. `WORKDIR` - you're stuff is going in a directory somewhere right? this just sets that :)

3. `COPY | ADD` - the whole point in creating a Dockerfile is to run some code from your machine in a "container". This simply copies your local code and places it where you specify in the image.

4. `RUN` - depending on your usecase you might want to install some dependencies in your image, from `go get ...` to `npm install` to base packages in the images OS using `apk, apt-get, yum`

5. `ENTRYPOINT | CMD` - this is the money maker here. This simply runs your app so when we eventually run your container, this command is what's being run. Think of this is the command used to start your application.

```
FROM base-image:tag-name

WORKDIR some/image/directory

COPY local/path image/path

RUN some command maybe apt-get, yum install, apk add?

ENTRYPOINT ["path/to/binary"]
```

From there once our Dockerfile is built we can build the image by running the following:

```
$ docker build -t api:latest -f Dockerfile . 
```

In english we are docker `build`ing, `-t` tagging our image as api:latest and `-f Dockerfile` building the image based on the file we called Dockerfile.

Once built you'll see the following successful message:

```
...
Removing intermediate container f1009d955d22
 ---> 7d34317b688c
Successfully built 7d34317b688c
Successfully tagged api:latest
```

After that we can simply run our container like so:

```
$ docker run api:latest
```

From there we should be able to go to `localhost:8080` and see our API running....

...

...

OK, I wasn't so honest :( we forgot one thing. We want to port forward the running app in our container to an open port on our machine. Simply add the following argument to our run statement above `-p 8080:8080` which would make for `docker run -p 8080:8080 api:latest`

Sweet :tada:

Remember when I said there was a fancy "app store" for docker images that houses tons of pre-built nifty images you can use? We're going to host and store this image we just made in AWS using a tool called ECR, or, Elastic Container Registry. Its essentially a private version of docker hub, since I doubt you want random people pulling your api image and stealing your precious startup code ;)

We'll dive into creating the registry next since we need an AWS account however for the sake of doing all things docker right here and now you would essentially run the following to push your image to your registry.

```
$ aws ecr create-repository --region us-east-1 --repository-name bflobox-api
$ docker tag api:latest ECR_REGISTRY_URL/api:latest
$ docker push ECR_REGISTRY_URL/api:latest
```

:tada: :allthethings: