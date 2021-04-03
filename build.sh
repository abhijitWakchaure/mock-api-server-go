#!/usr/bin/env bash
echo "### Deleting old dist files..."
rm -rf dist/*
packr2
APP_NAME="mock-api-server-go"
APP_VERSION="v1.0.4"
export CGO_ENABLED=0
echo "### Building for platform: linux/amd64"
env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/${APP_NAME}-${APP_VERSION}-linux-amd64
echo "### Building for platform: windows/amd64"
env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o dist/${APP_NAME}-${APP_VERSION}-win-amd64.exe
echo "### Building for platform: darwin/amd64"
env GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o dist/${APP_NAME}-${APP_VERSION}-darwin-amd64
packr2 clean
echo "### Building docker image"
docker build -t abhijitwakchaure/${APP_NAME}:${APP_VERSION} .
echo "### Pushing docker image to docker hub"
docker push abhijitwakchaure/${APP_NAME}:${APP_VERSION}