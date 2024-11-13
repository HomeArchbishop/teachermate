#!/bin/sh

if [ ! -d "./build" ]; then
  mkdir build
fi

go build -o ./build/ ./cmd/server/...
