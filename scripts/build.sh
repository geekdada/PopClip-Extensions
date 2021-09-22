#!/bin/bash

for arch in arm64 x86_64
do
  GOOS=darwin GOARCH="$arch" go build -o "./flomo.popclipext/flomo-macos-$arch" ./src/flomo
done

for name in flomo
do
  zip -r "${name}.zip" "./${name}.popclipext"
  mv -f "${name}.zip" "./Downloads/${name}.popclipextz"
done
