# Chapter-1

## Introduction to containers

[This Blog post](https://www.freecodecamp.org/news/a-beginner-friendly-introduction-to-containers-vms-and-docker-79a9e3e119b/) from Freecodecamp is little overwhelming but covers pretty much all the details one needs to know before starting with containers.

## Using docker engine
You can install docker on you machine by following the instruction from the [official page](https://docs.docker.com/desktop/install/windows-install/). Lsat time i used docker on windows, I didn't have a good experience. Docker on windows depends on [WSL](https://docs.microsoft.com/en-us/shows/). If you don't want to bother your self with docker installation, you can [docker playground](https://labs.play-with-docker.com/) as your online lab environment. You can use this lab env to follow along with this workshop.

## Containerising your first app
Move to the root of `Chapter-1` before running the commands.

- building container image
```bash
docker build -t workshop:v1 -f dockerfile.v1 .
```

- listing all the images on your machine
``` bash
docker images
```

- running images 
```bash
docker run -p 3000:3000 workshop:v1 .
```

- listing all the containers running on your machine 
```bash
docker ps
docker ps -a
```

- stopping a running container 
```bash 
docker stop <container_ir or container_name>
```

- digging through the image
```bash
docker 