#!/bin/bash
cd crawl
GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap *.go
zip -r bootstrap.zip bootstrap
cd ../serve
GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap *.go
zip -r bootstrap.zip bootstrap
