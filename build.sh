#!/bin/bash

set -eu

build_server() {
  echo "build_server start..."
  export GO111MODULE=on
  # go mod tidy
  # go mod vendor

  go build -o bin/user-rpc services/user/rpc/user.go
  go build -o bin/user-api services/user/api/user.go
  go build -o bin/wallpaper-rpc services/wallpaper/rpc/wallpaper.go
  go build -o bin/wallpaper-api services/wallpaper/api/wallpaper.go

  chmod +x bin/*

  echo "build_server successful"
}

run_server() {
  echo "run_server start..."
  nohup ./bin/user-rpc -f services/user/rpc/etc/user.yaml &
  nohup ./bin/user-api -f services/user/api/etc/user-api.yaml &
  nohup ./bin/wallpaper-rpc -f services/wallpaper/rpc/etc/wallpaper.yaml &
  nohup ./bin/wallpaper-api -f services/wallpaper/api/etc/wallpaper-api.yaml &
  echo "run_server successful"
}

build_server
#run_server