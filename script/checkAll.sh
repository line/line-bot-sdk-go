#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

echo "Run tests"
bash script/test.sh

echo "Check gofmt"
bash script/fmt.sh
git diff --exit-code

echo "Run go vet"
go vet $(go list ./... | grep -v /examples/)

echo "Compile examples"
while IFS= read -r -d '' file; do
  dir=$(dirname "$file")
  (
    cd "$dir"
    go build -o /dev/null
  )
done < <(find ./examples/ -name '*.go' -print0)
