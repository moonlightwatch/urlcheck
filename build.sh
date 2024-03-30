#!/bin/bash

go build -o ./bin/urlcheck ./
CGO_ENABLED=0 GOOS=windows  GOARCH=amd64 go build -o ./bin/urlcheck.exe ./

cp /usr/local/go/misc/wasm/wasm_exec.js bin/
GOOS=js GOARCH=wasm go build -ldflags '-s -w' -o bin/main.wasm ./