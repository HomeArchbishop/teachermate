#!/bin/sh

while [ $# -gt 0 ]; do
  case "$1" in
    --server-port=*)
      SERVER_PORT="${1#*=}"
      ;;
    --client-api-host=*)
      CLIENT_API_HOST="${1#*=}"
      ;;
  esac
  shift
done

if [ ! -d "./build" ]; then
  mkdir build
fi

cp config.yaml.example ./build/config.yaml

sed -i "s/port:[[:space:]][0-9.]*/port: ${SERVER_PORT}/" ./build/config.yaml
sed -i "s/server:[[:space:]].*/server: ${CLIENT_API_HOST}/" ./build/config.yaml
