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
