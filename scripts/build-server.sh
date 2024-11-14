#!/bin/sh

if [ ! -d "./build" ]; then
  mkdir build
fi

go build -buildvcs=false -o ./build/ ./cmd/server/...
