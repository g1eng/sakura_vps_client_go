#!/bin/sh -ex

API_VERSION=`jq .info.version < spec/spec.json`
cp README.md README.md.swp
awk -v version=": ${API_VERSION:-不明}"  '{ gsub(": 不明",version); print }' < README.tpl > README.md
{
  echo
  echo
  cat README.md.swp
  echo 'Nomura Suzume <ＳＵＺＵＭe[att]ｅa.g1e.org>'
  echo
  echo "## 本リポジトリのコード生成の再現方法について"
  echo 
  echo 'Makefileに記載のopenapi-generatorのバージョンがご自身のマシンに導入されていることをご確認の上、`make renew`を実行してください。'
  echo 
}  >> README.md
rm README.md.swp 
