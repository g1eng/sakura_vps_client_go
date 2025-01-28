#!/bin/sh -ex

API_VERSION=`jq .info.version < spec/spec.json`
awk "{ gsub(\": 不明\",\": ${API_VERSION:-不明}\"); print }" < README.tpl > README.md
