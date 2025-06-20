// 概要: ファイル種別判定のデモンストレーションを行うメインプログラム
// 主要仕様: 複数のファイル名に対して種別判定関数を呼び出し、結果を出力する
// 制限事項: ファイル拡張子のみで判定。実際のファイル内容は考慮しない
// TODO: 拡張子以外の判定方法も検討する
// FIXME: 拡張子が大文字の場合の判定追加検討
package main

import (
	"fmt"
	"os"
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
		info, err := GetFileInfo(filename)
		if err != nil {
			// ファイルが存在しない場合は種別のみ表示
			if os.IsNotExist(err) {
				fmt.Printf("ファイル名: %-15s => %s (ファイルなし)\n", filename, GetFileType(filename))
				continue
			}
			fmt.Printf("%s の読み込みに失敗しました: %v\n", filename, err)
			continue
		}
		fmt.Printf("ファイル名: %-15s => %s, サイズ: %d bytes, トークン数: %d\n",
			filename, info.Type, info.Size, info.Tokens)
	}
}
