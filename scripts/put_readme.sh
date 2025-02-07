#!/bin/sh -ex

API_VERSION=`jq .info.version < spec/spec.json`
cp README.md README.md.swp
awk -v version=": ${API_VERSION:-不明}"  '{ gsub(": 不明",version); print }' < README.tpl > README.md
{
  echo
  echo
  cat README.md.swp
}  >> README.md
rm README.md.swp 
