#!/bin/sh

# update config.js
if [ ! -z "$DEFAULT_API_URL" ]; then
  sed -i "s|DEFAULT_API_URL|DEFAULT_API_URL: '$DEFAULT_API_URL',//|g" ./public/config.js
fi

if [ ! -z "$DEFAULT_API_PROTOCOL" ]; then
  sed -i "s|DEFAULT_API_PROTOCOL|DEFAULT_API_PROTOCOL: '$DEFAULT_API_PROTOCOL',//|g" ./public/config.js
fi

echo "config.js updated"
