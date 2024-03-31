#!/usr/bin/sh

go build -o ./tmp/cli ./cmd/cli
./tmp/cli --register
