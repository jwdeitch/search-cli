#!/usr/bin/env bash

env GOOS=windows GOARCH=amd64 go build -o dist/win/s .
env GOOS=linux GOARCH=arm go build -o dist/linux/s .
go build -o dist/osx/s .