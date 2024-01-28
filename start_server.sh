#!/bin/sh
# Start redis server

set -e
tmpFile=$(mktemp)
go build -o "$tmpFile" src/*.go
exec "$tmpFile" "$@"
