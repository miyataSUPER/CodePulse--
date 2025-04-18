# プロジェクト名: CodePulse--
![Status](https://img.shields.io/badge/status-開発中-blue)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
[![Go Report Card](https://goreportcard.com/badge/github.com/miyataSUPER/CodePulse--)](https://goreportcard.com/report/github.com/miyataSUPER/CodePulse--)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## 概要

指定されたディレクトリ内のソースコードに関する統計情報（行数、ファイル数、言語ごとの内訳など）を表示するコマンドラインインターフェース（CLI）アプリケーションです。`tokei`リスペクト。

## 主な仕様

1.  **統計情報の収集:**
    *   指定されたディレクトリ（デフォルトはカレントディレクトリ）を再帰的に探索します。
    *   認識されたプログラミング言語のソースファイルを識別します。
    *   各ファイルについて、コード行、コメント行、空白行の数をカウントします。
    *   言語ごとに統計情報を集計します。
    *   プロジェクト全体の統計情報（合計ファイル数、合計行数など）を集計します。
2.  **結果の表示:**
    *   言語ごとのファイル数、コード行、コメント行、空白行、合計行数を表形式でコンソールに表示します。
    *   プロジェクト全体の合計ファイル数、コード行、コメント行、空白行、合計行数を表示します。
3.  **対応言語:**
    *   初期バージョンでは、いくつかの主要な言語に対応します（具体的な言語は要検討）。
    *   将来的には対応言語を容易に拡張できる設計を目指します。



## インストール

```bash
# プロジェクトの初期化
go mod init github.com/miyataSUPER/CodePulse--

# 依存関係のインストール
go mod tidy

# ビルド
go build
```

## 使用方法

### 基本的な使い方

```bash
# カレントディレクトリの統計情報を表示（デフォルト設定）
./codepulse

# 特定のディレクトリの統計情報を表示
./codepulse ./path/to/directory
```

### オプション

```bash
# ヘルプを表示
./codepulse --help

# 真偽値オプション
./codepulse --enable        # 特定の機能を有効化
./codepulse --verbose       # 詳細な出力を表示

# 出力形式の指定
./codepulse --format json   # JSON形式で出力（デフォルト）
./codepulse --format csv    # CSV形式で出力
./codepulse --format xml    # XML形式で出力

# 数値オプション
./codepulse --number 10     # 表示する結果の数を指定

# 文字列オプション
./codepulse --title "統計情報"  # 出力のタイトルを指定
```

### 出力例（デフォルト形式）

```
統計情報
========

Language  Files  Lines  Code  Comments  Blanks
Go        5      1000   800   100       100
Python    3      500    400   50        50
Total     8      1500   1200  150       150
```

### 出力例（JSON形式）

```json
{
  "title": "統計情報",
  "languages": [
    {
      "name": "Go",
      "files": 5,
      "lines": 1000,
      "code": 800,
      "comments": 100,
      "blanks": 100
    },
    {
      "name": "Python",
      "files": 3,
      "lines": 500,
      "code": 400,
      "comments": 50,
      "blanks": 50
    }
  ],
  "total": {
    "files": 8,
    "lines": 1500,
    "code": 1200,
    "comments": 150,
    "blanks": 150
  }
}
```

## テスト

```bash
# すべてのテストを実行
go test -v

# 特定のテストを実行
go test -v -run Test_hello
```

## プロジェクト構造

```
.
├── main.go          # メインアプリケーションコード
├── main_test.go     # テストコード
├── go.mod           # Goモジュール定義
├── LICENSE          # ライセンスファイル
└── README.md        # プロジェクトドキュメント
```


