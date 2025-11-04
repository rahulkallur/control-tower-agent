#!/bin/sh

# build docker image
docker buildx build --platform linux/amd64 --tag 794213689372.dkr.ecr.ap-south-1.amazonaws.com/control-tower-agent:1.0.3 --push .