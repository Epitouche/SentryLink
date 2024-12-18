#!/usr/bin/env bash

set -e
set -a
source ../.env
go test -cover -coverprofile=coverage.out -race ./...
set +a
