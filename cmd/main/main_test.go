// 概要:
//
//	IsArchiveFile関数のテーブル駆動テスト。
//
// 主な仕様:
//   - さまざまな拡張子のファイル名に対してアーカイブ判定の真偽を検証する。
//
// 制限事項:
//   - テストケースは必要に応じて追加可能。
//
// TODO:
//   - 増加する拡張子や異常系のテスト追加。
//
// FIXME:
//   - テスト失敗時の出力改善。
package main

import (
	"fmt"
	"testing"
)

func TestIsArchiveFile(t *testing.T) {
	testCases := []struct {
		given string // 入力ファイル名
		want  bool   // 期待値
	}{
		// テストケースはここに追加
	}

	for _, tc := range testCases {
		got := IsArchiveFile(tc.given)
		if got != tc.want {
			t.Errorf("IsArchiveFile(%q) = %v, want %v", tc.given, got, tc.want)
		}
	}
}

func ExampleIsArchiveFile() {
	fmt.Println(IsArchiveFile("hoge.tar"))
	fmt.Println(IsArchiveFile("hoge.tar.gz"))
	fmt.Println(IsArchiveFile("hoge.zip"))
	fmt.Println(IsArchiveFile("hoge.gz"))
	// Output:
	// true
	// true
	// true
	// false
}
