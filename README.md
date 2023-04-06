![PLATEAU VIEW 2.0](docs/logo.png)

# PLATEAU VIEW 2.0

PLATEAU VIEW 2.0 は以下のシステムにより構成されます。

- **PLATEAU CMS**: ビューワーに掲載する各種データの管理を行う。
- **PLATEAU Editor**: ビューワーの作成・公開をノーコードで行う。
- **PLATEAU VIEW**: 様々なPLATEAU関連データセットの可視化が可能なWebアプリケーション。

詳細は[「実証環境構築マニュアル 第3.0版」](https://www.mlit.go.jp/plateau/file/libraries/doc/plateau_doc_0009_ver03.pdf)を参照してください。  
開発環境の構築手法は[terraform](terraform)を参照してください。

## フォルダ構成

- [cms](cms): PLATEAU CMS
- [editor](editor): PLATEAU Editorとして採用されているOSS「[Re:Earth](https://github.com/reearth/reearth)」
- [plugin](plugin): PLATEAU VIEWで使用するRe:Earthのプラグイン
- [server](server): PLATEAU CMS のサイドカーサーバー（CMSと共に補助的に動作するサーバーアプリケーション）
- [terraform](terraform): PLATEAU VIEW 2.0をクラウド上に構築するためのTerraform

## ライセンス

[Apache License Version 2.0](LICENSE)
