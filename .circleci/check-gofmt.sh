#!/bin/bash
# Copyright 2019 Tricot Inc.
# Use of this source code is governed by the license in the LICENSE file.

# Checks that files are gofmt-ed (with -s).

set -e

cd "$(dirname "$BASH_SOURCE")/.."

ERRORS=$(find . -name '*.go' | \
    xargs -r gawk \
        'FNR > 3 || /Code generated .* DO NOT EDIT\./ {nextfile}; {print FILENAME; nextfile}' | \
    xargs -r gofmt -s -l)
if [[ -z "$ERRORS" ]]; then
    exit 0
fi

echo "gofmt errors found in:"
echo "$ERRORS"
exit 1
