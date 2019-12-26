#!/bin/sh
docker build -t welcome:build . -f Dockerfile.build
docker create --name extract welcome:build
docker cp extract:/welcome ./welcome
docker rm -f extract

docker build --no-cache -t welcome:small . -f Dockerfile.run
#rm -rf ./welcome
