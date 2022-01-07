#!/bin/bash

mkdir releases
GOOS=darwin GOARCH=amd64 go build -o releases/mok-$1-darwin-amd64/mok
GOOS=darwin GOARCH=arm64 go build -o releases/mok-$1-darwin-arm64/mok
GOOS=linux GOARCH=amd64 go build -o releases/mok-$1-linux-amd64/mok
GOOS=linux GOARCH=arm64 go build -o releases/mok-$1-linux-arm64/mok
GOOS=windows GOARCH=amd64 go build -o releases/mok-$1-windows-amd64/mok.exe

cd releases
zip -r mok-$1-darwin-amd64.zip mok-$1-darwin-amd64
zip -r mok-$1-darwin-arm64.zip mok-$1-darwin-arm64
zip -r mok-$1-linux-amd64.zip mok-$1-linux-amd64
zip -r mok-$1-linux-arm64.zip mok-$1-linux-arm64
zip -r mok-$1-windows-amd64.zip mok-$1-windows-amd64

rm -r mok-$1-darwin-amd64
rm -r mok-$1-darwin-arm64
rm -r mok-$1-linux-amd64
rm -r mok-$1-linux-arm64
rm -r mok-$1-windows-amd64
