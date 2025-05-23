package main

import "strings"

// 概要:
//
//	指定されたファイルパスがアーカイブファイルかどうかを判定する関数。
//
// 主な仕様:
//   - .tar, .tar.gz, .tar.bz2, .zip, .war, .ear のいずれかの拡張子であればtrueを返す。
//
// 制限事項:
//   - 拡張子の追加・削除はtargetsスライスを編集することで対応可能。
//
// TODO:
//   - 拡張子の大文字小文字対応。
//   - より多くのアーカイブ形式への対応。
//
// FIXME:
//   - パスが空文字や不正な場合のエラーハンドリング。
func IsArchiveFile(path string) bool {
	targets := []string{
		".tar", ".tar.gz", ".tar.bz2",
		".zip", ".war", ".ear",
	}
	for _, suffix := range targets {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}
