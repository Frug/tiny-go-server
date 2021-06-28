# Tiny go server
Build this image (or skip this step since it's already on docker hub):
```
make build
```

Run this image (on port 8080):

```
docker pull dasfrugen/tiny-go-server:1.0
docker run -p 8080:8080 --rm dasfrugen/tiny-go-server:1.0
```

Check its output

```
curl localhost:8080
> Pong
```

You can change the listening port of the container by setting a `PORT` environment variable.

This image is built with ideas from [Nick Gauthier](https://www.cloudbees.com/blog/building-minimal-docker-containers-for-go-applications) and [Adriaan de Jonge](https://xebia.com/blog/create-the-smallest-possible-docker-container/).

## But Why
I needed a tiny image that would run and pass an http health check, and couldn't find one prebuilt on docker hub.

This image is under 5mb and fits that requirement.

There's an even smaller possibility using Rust: https://blog.semicolonsoftware.de/building-minimal-docker-containers-for-rust-applications/
