# アプリケーション名とバージョン情報

App := 'CodePulse--'
Version := `grep '^const VERSION = ' main.go | sed "s/^VERSION = \"\(.*\)\"/\1/g"`

# ヘルプメッセージの表示
@help:
    echo "Build tool for {{ App }} {{ Version }} with Just"
    echo "Usage: just <recipe>"
    echo ""
    just --list

# ビルドとテストの実行
build: test
    go build -o codepulse main.go

# テストの実行
test:
    go test -covermode=count -coverprofile=coverage.out ./... 
