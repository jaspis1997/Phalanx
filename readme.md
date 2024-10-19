# Phalanx

レガシー版から移行中のため、動作しない。

## Project Overview / プロジェクト概要

Phalanxは、自分専用の動画サイトを作成できるプロジェクトです。数多くの動画を自由に楽しむことができ、簡単に動画を検索する機能も備えています。また、ユーザーが慣れ親しんだUIを採用しており、使いやすさが特徴です。
友達や家族との思い出やSNSでバズった動画など、たくさんの動画を手軽に、自分だけで楽しめるようなソフトウェアを目指しています。

現在、各ロジックを検証し、設計を行っている段階です。要求の変更に伴い、現在、レガシー版から以降の途中であり、各機能は動作しません。

## Technology Stack / 採用技術

Phalanxでは、各種機能を最適に実現するために以下の技術を採用しています。

### Frontend program language / TypeScript

型安全性を高め、開発効率を向上させるため、フロントエンドではTypeScriptを使用しています。

- [TypeScript: JavaScript With Syntax For Types.](https://www.typescriptlang.org/)

#### UI Framework / SolidJS

高速なレンダリングとリアクティブなUIを提供するため、UIフレームワークとしてSolidJSを採用しました。
コンポーネントとして、Material UIのSolidJS版である`SUID`を使用します。

- [リアクティブJavascriptライブラリ](https://www.solidjs.com/)
- [SUID](https://suid.io/)

#### Module Bundler / Vite

高速なビルドと開発体験を実現するため、フロントエンドのバンドラーにはViteを選択しました。

- [Vite | 次世代フロントエンドツール](https://ja.vite.dev/)

#### TypeScript Linter & Formatter / Biome.js

コード品質の向上と自動フォーマットを行うため、LintとFormatterとしてBiome.jsを使用しています。
PrettierやESLintよりも簡単に利用できます。

- [Biome、Webのためのツールチェーン](https://biomejs.dev/ja/)

#### Node.js Package Manager / pnpm

高速かつ効率的な依存関係管理を実現するため、pnpmを採用しました。
Node.jsのバージョンも管理でき、より厳格なPackage管理が実現できます。厳格なPackage管理やNode.jsのバージョン管理により、環境固有のエラーの削減が見込めます。

- [Fast, disk space efficient package manager | pnpm](https://pnpm.io/ja/)

### Backend program language / Goland

シングルバイナリ形式でのデプロイが求められるため、Goを採用しました。また、GoはAPIサーバーの構築に強く、高速なバックエンド処理が可能です。

- [The Go Programming Language](https://go.dev/)

#### CLI Libraly / cobra

CLIツールの開発を簡単に行うため、GoのコマンドライブラリであるCobraを使用しています。

- [spf13/cobra: A Commander for modern Go CLI interactions](https://github.com/spf13/cobra)

#### Web framework　/ Gin

高速かつ軽量なAPI開発を行うため、GoのWebフレームワークであるGinを選びました。

- [Gin Web Framework](https://gin-gonic.com/ja/)

#### Database Object Relational  Mapping / Bun

SQLデータベースとの効率的なやり取りを実現するため、GoのORMライブラリであるBunを採用しています。
SQLファーストな使い方ができ、SQLを知っているのであれば比較的簡単に理解できます。

- [uptrace/bun: SQL-first Golang ORM](https://github.com/uptrace/bun)

### Container / Docker

コンテナベースの開発環境を提供し、一貫した動作を保証するため、Dockerを採用しています。
Minioやソフトウェアの動作検証にDockerを使用しています。最終的なリリースではDockerで利用することは想定していません。これは、だれでも手軽に利用するうえで、Dockerを前提にするのはハードルが高いためです。

また、Docker Desktopの代わりに、`Podman`を使用しています。

- [Docker: コンテナー アプリケーション開発の加速](https://www.docker.com/ja-jp/)
- [Podman](https://podman.io/)

#### Object Storage / Minio

サムネイルの保存には、オブジェクトストレージのMinioを使用し、大容量データの管理を効率化しています。
当初は、Webサービスとして提供することも視野に入れていたため、採用していました。現在、ソフトウェアとして利用できる要求のほうが強いため、将来的には変更する予定です。

- [MinIO | S3 Compatible Storage for AI](https://min.io/)

## Features / 機能

- 動画再生
- タイトルを使った動画検索
- ランダムに選ばれた動画の一覧表示
- お気に入り機能
- サムネイルの自動生成および変更

## Installation / インストール方法

TBD

- 実行ファイルとして提供する
- 必要なリソースはサーバーから自動でDLすることを目指す
- アップデートも自動で行うことができるようにする

## Usage / 使い方

TBD

## Configuration / 設定

TBD
