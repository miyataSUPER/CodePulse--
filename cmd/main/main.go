// 概要: ファイル種別判定のデモンストレーションを行うメインプログラム
// 主要仕様: 複数のファイル名に対して種別判定関数を呼び出し、結果を出力する
// 制限事項: ファイル拡張子のみで判定。実際のファイル内容は考慮しない
// TODO: 拡張子以外の判定方法も検討する
// FIXME: 拡張子が大文字の場合の判定追加検討
package main

import (
	"fmt"
)

func main() {
	// ファイル種別判定のデモンストレーション
	testFiles := []string{
		"document.txt",
		"photo.jpg",
		"music.mp3",
		"movie.mp4",
		"report.pdf",
		"script.py",
		"program.exe",
		"archive.zip",
		"data.csv",
		"unknown.xyz",
	}

	fmt.Println("=== ファイル種別判定デモ ===")
	for _, filename := range testFiles {
		fmt.Printf("ファイル名: %-15s => ", filename)

		var fileType string
		switch {
		case IsTextFile(filename):
			fileType = "テキスト"
		case IsImageFile(filename):
			fileType = "画像"
		case IsAudioFile(filename):
			fileType = "音声"
		case IsVideoFile(filename):
			fileType = "動画"
		case IsDocumentFile(filename):
			fileType = "ドキュメント"
		case IsScriptFile(filename):
			fileType = "スクリプト"
		case IsExecutableFile(filename):
			fileType = "実行ファイル"
		case IsArchiveFile(filename):
			fileType = "アーカイブ"
		default:
			fileType = "不明"
		}

		fmt.Println(fileType)
	}
}
