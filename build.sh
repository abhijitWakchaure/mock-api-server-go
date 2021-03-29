#!/usr/bin/env bash
packr2
appName="mock-api-server-go"
appVersion="v1.0.0"
export CGO_ENABLED=0
echo "### Building for platform: linux/amd64"
env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/${appName}-${appVersion}-linux-amd64
echo "### Building for platform: windows/amd64"
env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o dist/${appName}-${appVersion}-win-amd64.exe
echo "### Building for platform: darwin/amd64"
env GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o dist/${appName}-${appVersion}-darwin-amd64
packr2 clean
echo "### Building docker image"
docker build -t abhijitwakchaure/mock-api-server-go .
