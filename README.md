# Hello World

[![Build Status](https://travis-ci.org/rohangupta5328/hello-world.svg?branch=master)](https://travis-ci.org/rohangupta5328/hello-world)

This project is just to practice GoLang and stay up to date with the latest and greatest stuff. Keeping my golang skills fresh in my free time.


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