# これは何

さくらのVPS APIをGoで利用するための非公式クライアントライブラリです。

このライブラリは[上流の仕様書](https://manual.sakura.ad.jp/vps/api/api-doc/)が [CC-BY 4.0国際ライセンス](https://creativecommons.org/licenses/by/4.0/deed.ja)で許諾する部分を切り出して作成した[仕様書](/spec/spec.json)を元に、OpenAPI Generator の Go generator を用いてコード生成を行ったものです。

このパッケージは[さくらインターネット](https://www.sakura.ad.jp)が著作権を有する仕様書(バージョン: 不明)を元として作成されています。


# 利用方法

本リポジトリでは、クライアントライブラリの包括的な利用方法のサポートを提供しません。
その代わり、仕様書に含まれるAPIの利用方法をご自身のエディタで参照できるよう、/docs配下にAPIの利用方法のドキュメンテーションが格納されています。

goplsのようなLSPサーバーを用いれば、それぞれの関数の利用方法について有益なサジェストを得ることができるでしょう。


# ライセンス

Apache 2.0

ただし、[仕様書](/spec/spec.json)は[CC-BY 4.0国際ライセンス](https://creativecommons.org/licenses/by/4.0/deed.ja)で許諾されます。


