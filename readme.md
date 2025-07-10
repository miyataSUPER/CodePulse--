# プロジェクト名: CodePulse++
![Status](https://img.shields.io/badge/status-開発中-blue)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
[![Go Report Card](https://goreportcard.com/badge/github.com/miyataSUPER/CodePulse--)](https://goreportcard.com/report/github.com/miyataSUPER/CodePulse--)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## 概要

ファイル名の拡張子に基づいてファイル種別を判定するGo言語ライブラリです。テキスト、画像、音声、動画、ドキュメント、スクリプト、実行ファイル、アーカイブなど、様々なファイル種別を自動判定できます。

- **ファイル種別判定のロジックは `filetype.go` に実装されています。**
- **`main.go` はデモ用のエントリポイントです。**

## 主な仕様

1.  **ファイル種別判定機能:**
    *   ファイル名（拡張子）に基づいて、8つのカテゴリのファイル種別を判定します。
    *   テキストファイル（.txt, .md, .csv, .log）
    *   画像ファイル（.jpg, .jpeg, .png, .gif, .bmp, .svg, .webp）
    *   音声ファイル（.mp3, .wav, .ogg, .flac, .aac）
    *   動画ファイル（.mp4, .avi, .mov, .wmv, .flv, .mkv）
    *   ドキュメントファイル（.pdf, .doc, .docx, .xls, .xlsx, .ppt, .pptx）
    *   スクリプトファイル（.sh, .bat, .ps1, .py, .rb, .pl, .js, .ts）
    *   実行ファイル（.exe, .dll, .so, .bin, .app）
    *   アーカイブファイル（.zip, .tar, .gz, .bz2, .rar, .7z, .xz, .tar.gz, .tar.bz2, .tar.xz, .tgz, .tbz2, .lz, .lzma, .z, .cab, .arj, .war, .ear）

2.  **ファイル情報取得機能:**
    *   `GetFileInfo(path string) (*FileInfo, error)` でファイルの詳細情報を取得
    *   ファイル種別、サイズ（バイト）、トークン数（単語数）を一括取得
    *   トークン数は全ファイルタイプで計算（バイナリファイルでも実行される）
    *   トークン数は空白文字（スペース、タブ、改行）で区切られた単語数をカウント
    *   空ファイルの場合は0、バイナリファイルの場合は意味のない値が返される

3.  **API設計:**
    *   各ファイル種別ごとに個別の判定関数を提供します。
    *   すべての関数は `func IsXxxFile(path string) bool` の形式です。
    *   拡張子の追加・削除は各関数内の配列を編集することで容易に対応可能です。

4.  **デモプログラム:**
    *   `main.go` でファイル種別判定のデモンストレーションを実行できます。
    *   日本語でファイル種別を表示します。

## 制限事項

*   現在は拡張子のみに基づく判定で、ファイルの内容（MIMEタイプなど）は考慮しません。
*   拡張子の大文字小文字の違いには対応していません（将来対応予定）。
*   複数の拡張子を持つファイル（例: .tar.gz）は、より一般的なカテゴリ（アーカイブ）として判定されます。
*   トークン数計算は全ファイルタイプで実行されるため、バイナリファイル（画像、動画、実行ファイルなど）でも不必要にファイル全体を読み取ります。
*   大きなバイナリファイルの場合、メモリ使用量が増加する可能性があります。

## インストール

```bash
# プロジェクトのクローン
git clone https://github.com/miyataSUPER/CodePulse--

# ディレクトリに移動
cd CodePulse--/cmd/main

# 依存関係のインストール
go mod tidy

# ビルド
go build .
```

### Dockerイメージの利用

Dockerfile を使ってコンテナイメージをビルド・実行できます。イメージは GHCR (GitHub Container Registry) にも公開されます。

```bash
# イメージをビルド
docker build -t codepulse .

# 実行
docker run --rm codepulse
```

## 使用方法

### デモプログラムの実行

```bash
# デモプログラムを実行
./main
```

### API使用例

```go
package main

import (
    "fmt"
)

func main() {
    // テキストファイルの判定
    fmt.Println(IsTextFile("document.txt"))  // true
    fmt.Println(IsTextFile("readme.md"))     // true
    
    // 画像ファイルの判定
    fmt.Println(IsImageFile("photo.jpg"))    // true
    fmt.Println(IsImageFile("logo.png"))     // true
    
    // 音声ファイルの判定
    fmt.Println(IsAudioFile("music.mp3"))    // true
    fmt.Println(IsAudioFile("sound.wav"))    // true
    
    // 動画ファイルの判定
    fmt.Println(IsVideoFile("movie.mp4"))    // true
    fmt.Println(IsVideoFile("clip.avi"))     // true
    
    // ドキュメントファイルの判定
    fmt.Println(IsDocumentFile("report.pdf")) // true
    fmt.Println(IsDocumentFile("data.xlsx"))  // true
    
    // スクリプトファイルの判定
    fmt.Println(IsScriptFile("script.py"))   // true
    fmt.Println(IsScriptFile("build.sh"))    // true
    
    // 実行ファイルの判定
    fmt.Println(IsExecutableFile("app.exe")) // true
    fmt.Println(IsExecutableFile("lib.dll")) // true
    
    // アーカイブファイルの判定
    fmt.Println(IsArchiveFile("data.zip"))   // true
    fmt.Println(IsArchiveFile("backup.tar.gz")) // true

    // ファイル情報の取得（テキストファイル）
    info, err := GetFileInfo("sample.txt")
    if err == nil {
        fmt.Printf("Type: %s, Size: %d bytes, Tokens: %d\n",
            info.Type, info.Size, info.Tokens)
    }
    
    // バイナリファイルでもトークン数が計算される（注意）
    imageInfo, err := GetFileInfo("photo.jpg")
    if err == nil {
        fmt.Printf("Image: Type=%s, Size=%d bytes, Tokens=%d\n",
            imageInfo.Type, imageInfo.Size, imageInfo.Tokens)
    }
    
    // 複数行のテキストファイルでのトークン数例
    scriptInfo, err := GetFileInfo("script.py")
    if err == nil {
        fmt.Printf("Script: Type=%s, Size=%d bytes, Tokens: %d\n",
            scriptInfo.Type, scriptInfo.Size, scriptInfo.Tokens)
    }
}
```

### デモ出力例

```
=== ファイル種別判定デモ ===
ファイル名: document.txt    => テキスト (ファイルなし)
ファイル名: photo.jpg       => 画像 (ファイルなし)
ファイル名: music.mp3       => 音声 (ファイルなし)
ファイル名: movie.mp4       => 動画 (ファイルなし)
ファイル名: report.pdf      => ドキュメント (ファイルなし)
ファイル名: script.py       => スクリプト (ファイルなし)
ファイル名: program.exe     => 実行ファイル (ファイルなし)
ファイル名: archive.zip     => アーカイブ (ファイルなし)
ファイル名: data.csv        => テキスト (ファイルなし)
ファイル名: unknown.xyz     => 不明 (ファイルなし)
```

### トークン数計算の詳細例

```go
// 例1: シンプルなテキストファイル
// 内容: "Hello world example"
// 結果: Tokens = 3

// 例2: 複数行のテキストファイル
// 内容: 
// "package main
// import "fmt"
// func main() {
//     fmt.Println("Hello")
// }"
// 結果: Tokens = 12

// 例3: 空のテキストファイル
// 内容: ""
// 結果: Tokens = 0

// 例4: バイナリファイル（画像）
// 内容: バイナリデータ
// 結果: Tokens = 意味のない値（例: 12345）

// 例5: CSVファイル
// 内容: "name,age,city\nJohn,25,Tokyo\nJane,30,Osaka"
// 結果: Tokens = 8
```

## API リファレンス

### ファイル種別判定関数

| 関数名 | 対応拡張子 | 戻り値 |
|--------|------------|--------|
| `IsTextFile(path string) bool` | .txt, .md, .csv, .log | ファイルがテキストファイルかどうか |
| `IsImageFile(path string) bool` | .jpg, .jpeg, .png, .gif, .bmp, .svg, .webp | ファイルが画像ファイルかどうか |
| `IsAudioFile(path string) bool` | .mp3, .wav, .ogg, .flac, .aac | ファイルが音声ファイルかどうか |
| `IsVideoFile(path string) bool` | .mp4, .avi, .mov, .wmv, .flv, .mkv | ファイルが動画ファイルかどうか |
| `IsDocumentFile(path string) bool` | .pdf, .doc, .docx, .xls, .xlsx, .ppt, .pptx | ファイルがドキュメントファイルかどうか |
| `IsScriptFile(path string) bool` | .sh, .bat, .ps1, .py, .rb, .pl, .js, .ts | ファイルがスクリプトファイルかどうか |
| `IsExecutableFile(path string) bool` | .exe, .dll, .so, .bin, .app | ファイルが実行ファイルかどうか |
| `IsArchiveFile(path string) bool` | .zip, .tar, .gz, .bz2, .rar, .7z, .xz, .tar.gz, .tar.bz2, .tar.xz, .tgz, .tbz2, .lz, .lzma, .z, .cab, .arj, .war, .ear | ファイルがアーカイブファイルかどうか |
| `GetFileInfo(path string) (*FileInfo, error)` | ー | ファイル種別・サイズ・トークン数をまとめて取得（全ファイルタイプでトークン数計算） |
| `FileInfo` 構造体 | ー | Type: ファイル種別, Size: サイズ（バイト）, Tokens: トークン数（単語数） |

### トークン数計算の仕様

**計算方法:**
- `strings.Fields()` を使用して空白文字（スペース、タブ、改行）で区切られた単語数をカウント
- 空ファイルの場合は0を返す
- バイナリファイルでも計算されるが、意味のない値が返される

**注意事項:**
- 全ファイルタイプで計算されるため、大きなバイナリファイルでは処理時間が長くなる
- 日本語テキストの場合、文字単位ではなく単語単位でカウントされる
- 特殊文字や記号も単語としてカウントされる

## テスト

```bash
# すべてのテストを実行
go test -v

# Example関数のテストを実行
go test -v -run Example

# 特定のテストを実行
go test -v -run TestIsArchiveFile
```

## プロジェクト構造

```
.
├── cmd/
│   └── main/
│       ├── main.go          # デモ用エントリポイント
│       ├── filetype.go      # ファイル種別判定ロジック
│       └── main_test.go     # テストコード・Example関数
├── .github/
│   └── workflows/
│       ├── build.yaml       # CI/CDワークフロー
│       └── update_version.yaml # バージョン自動更新ワークフロー
├── go.mod                   # Goモジュール定義
├── LICENSE                  # ライセンスファイル
└── README.md               # プロジェクトドキュメント
```

## TODO

1. **拡張性の向上**
   - 拡張子の大文字小文字対応
   - より多くのファイル形式への対応
   - 設定ファイルによる拡張子カスタマイズ

2. **機能追加**
   - MIMEタイプによる判定
   - ファイル内容による判定
   - 複数ファイルの一括判定
   - ファイル種別に応じたトークン数計算の最適化（テキストファイルのみ計算）

3. **テスト充実**
   - エラーケースのテスト追加
   - パフォーマンステスト
   - 異常系のテスト


