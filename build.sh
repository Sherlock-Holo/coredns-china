#!/usr/bin/env bash

git clone --depth 1 https://github.com/coredns/coredns.git

cd coredns

git checkout -b build

go get -v github.com/Sherlock-Holo/coredns-china/pkg/plugin

echo china-list:github.com/Sherlock-Holo/coredns-china/pkg/plugin >> plugin.cfg

go generate coredns.go

go mod tidy

go mod download

go mod tidy

make
