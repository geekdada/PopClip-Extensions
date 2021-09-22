#!/bin/sh

export BIN="./flomo-macos"
export ARCH=`uname -m`

if [ "$ARCH" = "x86_64" ]; then
  BIN="$BIN-amd64"
else
  BIN="$BIN-arm64"
fi

$BIN "$@"
