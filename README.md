# ブログAPI プロジェクト

このプロジェクトは、「[APIを作りながら進むGo中級者への道](https://techbookfest.org/product/jXDAEU1dR53kbZkgtDm9zx?productVariantID=dvjtgpjw8VDTXNqKaanTVi)」という書籍を参考に実装したブログAPIです。

## 学習内容

本プロジェクトを通じて、以下の内容を学習しました：

1. **ユニットテスト**
   - `testing`パッケージの活用
   - `httptest`パッケージを使用したHTTPハンドラのテスト
   - 前処理と後処理の共通化

2. **3層アーキテクチャ**
   - サービス層、コントローラ層、ルータ層の分離
   - クリーンアーキテクチャの基本概念の適用

3. **ミドルウェア**
   - ロギング機能の実装

4. **並行処理とベンチマーク**
   - Goルーチンとチャネルの基本
   - `testing.B`を使用したベンチマークテスト

5. **認証**
   - GoogleのSingle Sign-On (SSO) の基本概念

6. **コンテキスト管理**
   - `context`パッケージを使用したトレースIDの管理

## 参考文献

- [APIを作りながら進むGo中級者への道](https://techbookfest.org/product/jXDAEU1dR53kbZkgtDm9zx?productVariantID=dvjtgpjw8VDTXNqKaanTVi)
