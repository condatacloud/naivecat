#! /bin/bash

BuildTime=`date '+%Y-%m-%d %H:%M:%S'`
GoVersion=`go version`
GitHash=`git rev-parse --short HEAD`
GitBranch=`git name-rev --name-only HEAD`

go build -ldflags "-X 'naivecat/model.GoVersion=$GoVersion'
                   -X 'naivecat/model.GitHash=$GitHash'
                   -X 'naivecat/model.BuildTime=$BuildTime'
                   -X 'naivecat/model.GitBranch=$GitBranch'" -o bin/naivecat naivecat.go

# windows build
export GOOS=windows
export GOARCH=amd64
export CGO_ENABLED=1
export CC=x86_64-w64-mingw32-gcc
export CXX=x86_64-w64-mingw32-g++
export HOST=x86_64-w64-mingw32
go build -ldflags "-X 'naivecat/model.GoVersion=$GoVersion'
                   -X 'naivecat/model.GitHash=$GitHash'
                   -X 'naivecat/model.BuildTime=$BuildTime'
                   -X 'naivecat/model.GitBranch=$GitBranch' -s -w -H=windowsgui -extldflags=-static" -p 4 -v -o bin/naivecat.exe 
