#!/bin/sh

mkdir -p cache
docker run -ti -p 8080:8080 -v`pwd`/cache:/go -v`pwd`:/data --workdir=/data golang:1.20 go $@
