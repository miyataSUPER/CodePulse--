package main

import (
	"fmt"
	"strings"
)

// 概要:
//
//	指定されたファイルパスがアーカイブファイルかどうかを判定する関数。
//
// 主な仕様:
//   - .tar, .tar.gz, .tar.bz2, .zip, .war, .ear のいずれかの拡張子であればtrueを返す。

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

// 概要:
//
//	指定されたファイルパスがテキストファイルかどうかを判定する関数。
//
// 主な仕様:
//   - .txt, .md, .csv, .log のいずれかの拡張子であればtrueを返す。

func IsTextFile(path string) bool {
	targets := []string{
		".txt", ".md", ".csv", ".log",
	}
	for _, suffix := range targets {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}

// 概要:
//
//	指定されたファイルパスが画像ファイルかどうかを判定する関数。
//
// 主な仕様:
//   - .jpg, .jpeg, .png, .gif, .bmp, .svg, .webp のいずれかの拡張子であればtrueを返す。

func IsImageFile(path string) bool {
	targets := []string{
		".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg", ".webp",
	}
	for _, suffix := range targets {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}

// 概要:
//
//	指定されたファイルパスが音声ファイルかどうかを判定する関数。
//
// 主な仕様:
//   - .mp3, .wav, .ogg, .flac, .aac のいずれかの拡張子であればtrueを返す。

func IsAudioFile(path string) bool {
	targets := []string{
		".mp3", ".wav", ".ogg", ".flac", ".aac",
	}
	for _, suffix := range targets {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}

// 概要:
//
//	指定されたファイルパスが動画ファイルかどうかを判定する関数。
//
// 主な仕様:
//   - .mp4, .avi, .mov, .wmv, .flv, .mkv のいずれかの拡張子であればtrueを返す。

func IsVideoFile(path string) bool {
	targets := []string{
		".mp4", ".avi", ".mov", ".wmv", ".flv", ".mkv",
	}
	for _, suffix := range targets {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}

// 概要:
//
//	指定されたファイルパスがドキュメントファイルかどうかを判定する関数。
//
// 主な仕様:
//   - .pdf, .doc, .docx, .xls, .xlsx, .ppt, .pptx のいずれかの拡張子であればtrueを返す。

func IsDocumentFile(path string) bool {
	targets := []string{
		".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx",
	}
	for _, suffix := range targets {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}

// 概要:
//
//	指定されたファイルパスがスクリプトファイルかどうかを判定する関数。
//
// 主な仕様:
//   - .sh, .bat, .ps1, .py, .rb, .pl, .js, .ts のいずれかの拡張子であればtrueを返す。

func IsScriptFile(path string) bool {
	targets := []string{
		".sh", ".bat", ".ps1", ".py", ".rb", ".pl", ".js", ".ts",
	}
	for _, suffix := range targets {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}

// 概要:
//
//	指定されたファイルパスが実行ファイルかどうかを判定する関数。
//
// 主な仕様:
//   - .exe, .dll, .so, .bin, .app のいずれかの拡張子であればtrueを返す。

func IsExecutableFile(path string) bool {
	targets := []string{
		".exe", ".dll", ".so", ".bin", ".app",
	}
	for _, suffix := range targets {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}

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


