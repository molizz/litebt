#!/usr/bin/env bash

export GOOS=linux
export GOARCH=amd64

cd spider && go build -o ../cmd/spider -ldflags "-s -w"
cd ..
