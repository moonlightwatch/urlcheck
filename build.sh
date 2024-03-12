#!/bin/bash

go build -o ./bin/urlcheck ./
CGO_ENABLED=0 GOOS=windows  GOARCH=amd64 go build -o ./bin/urlcheck.exe ./