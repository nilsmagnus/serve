#!/bin/sh

cd ..
env GOOS=linux GOARCH=386 go build && mv serve dist/linux
env GOOS=windows GOARCH=386 go build && mv serve.exe dist/win
env GOOS=darwin GOARCH=386 go build && mv serve dist/darwin