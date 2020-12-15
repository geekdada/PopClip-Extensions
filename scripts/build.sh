#!/bin/bash

go build -o ./flomo.popclipext ./src/flomo

for name in Douban flomo
do
  zip -r "${name}.zip" "./${name}.popclipext"
  mv -f "${name}.zip" "./Downloads/${name}.popclipextz"
done
