#!/bin/sh -ex

API_VERSION=`jq .info.version < spec/spec.json`
cp README.md README.md.swp
awk "{ gsub(\": 不明\",\": ${API_VERSION:-不明}\"); print }" < README.tpl > README.md
{
  echo
  echo
  cat README.md.swp
}  >> README.md
rm README.md.swp 
