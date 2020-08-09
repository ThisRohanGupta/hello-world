# Hello World
[![Build Status](https://cloud.drone.io/api/badges/ThisRohanGupta/hello-world/status.svg)](https://cloud.drone.io/ThisRohanGupta/hello-world)


## Dockerfile 

``` 
FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
```

Repo: https://hub.docker.com/repository/docker/srgupta5328/go-k8s 

## Sample GoLang Project
This project is a sample project to demo CI/CD. 

## Deployed by Harness

url: https://app.harness.io 
