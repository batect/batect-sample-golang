#!/usr/bin/env bash

set -euo pipefail

unformatted_files=$(gofmt -l $(find . -name '*.go' -not -path './.batect/caches/*'))

if [ "$unformatted_files" == "" ]; then
    echo "All files are formatted correctly."
else
    echo "The following files are not formatted correctly:"
    echo "$unformatted_files"
    echo
    echo "Run 'go fmt' to reformat them."
    exit 1
fi
