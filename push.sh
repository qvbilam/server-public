#!/bin/bash
# shellcheck disable=SC2086

echo -e "\033[0;32mPlease input server version\033[0m"
read tag
imageName=qvbilam/api-server-public-alpine
originImageName=registry.cn-hangzhou.aliyuncs.com/qvbilam/api-server-public

# build image
docker build -t ${imageName} .
# login hub
docker login --username=13501294164 registry.cn-hangzhou.aliyuncs.com
# tag image
docker tag ${imageName} ${originImageName}:${tag}
# push image
docker push ${originImageName}:${tag}