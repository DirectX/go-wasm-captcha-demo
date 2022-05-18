#!/bin/sh

mkdir -p wasm
GOOS=js GOARCH=wasm go build -o wasm/captcha.wasm
