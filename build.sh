#!/usr/bin/env sh

set -x

gitVersion="$(git rev-parse --abbrev-ref HEAD)-$(git rev-parse --short HEAD)"
buildTime="$(date -u '+%Y-%m-%d %H:%M:%S') UTC"

sed -i "s/__GIT_VERSION/$gitVersion/g" version/version.go
sed -i "s/__BUILD_TIME/$buildTime/g" version/version.go

go build -o bin/HomeAI