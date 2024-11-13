#!/bin/sh

# 检查环境变量是否存在
if [ -z "$DEFAULT_API_URL" ]; then
  echo "DEFAULT_API_URL is not set"
  exit 0
fi

# 更新 config.js 文件
sed -i "s|DEFAULT_API_URL|DEFAULT_API_URL: '$DEFAULT_API_URL'//|g" ./public/config.js

echo "config.js updated"
