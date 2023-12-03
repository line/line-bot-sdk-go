#!/bin/bash
set -ex
goimports -w "$*"
gofmt -l -s -w "$*"
