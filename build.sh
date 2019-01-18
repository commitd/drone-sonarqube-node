#!/bin/bash

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o drone-sonarqube-node

docker build -t committed/drone-sonarqube-node .