# アプリケーション名とバージョン情報

App := 'CodePulse++'
Version := if path_exists("VERSION") == "true" { `cat VERSION 2>/dev/null || echo "0.1.0"` } else { "0.1.0" }

# ヘルプメッセージの表示
@help:
    echo "Build tool for {{ App }} v{{ Version }} with Just"
    echo "Usage: just <recipe>"
    echo ""
    just --list

# 依存関係の整理
@tidy:
    go mod tidy

# ビルドとテストの実行
build: test
    cd cmd/main && go build -o ../../{{ App }} ./...

# テストの実行
@test:
    cd cmd/main && go test -v -covermode=count -coverprofile=../../coverage.out ./...

# 実行ファイルの起動
@run: build
    ./{{ App }}

# カバレッジレポートの表示
@coverage: test
    go tool cover -html=coverage.out -o coverage.html
    echo "カバレッジレポートが coverage.html に生成されました"

# クリーンアップ
@clean:
    rm -f {{ App }} coverage.out coverage.html coverage.lcov
