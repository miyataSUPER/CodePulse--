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

func TestIsTextFile(t *testing.T) {
	testCases := []struct {
		given string // 入力ファイル名
		want  bool   // 期待値
	}{
		// テストケースはここに追加
	}

	for _, tc := range testCases {
		got := IsTextFile(tc.given)
		if got != tc.want {
			t.Errorf("IsTextFile(%q) = %v, want %v", tc.given, got, tc.want)
		}
	}
}

func TestIsImageFile(t *testing.T) {
	testCases := []struct {
		given string // 入力ファイル名
		want  bool   // 期待値
	}{
		// テストケースはここに追加
	}

	for _, tc := range testCases {
		got := IsImageFile(tc.given)
		if got != tc.want {
			t.Errorf("IsImageFile(%q) = %v, want %v", tc.given, got, tc.want)
		}
	}
}

func TestIsAudioFile(t *testing.T) {
	testCases := []struct {
		given string // 入力ファイル名
		want  bool   // 期待値
	}{
		// テストケースはここに追加
	}

	for _, tc := range testCases {
		got := IsAudioFile(tc.given)
		if got != tc.want {
			t.Errorf("IsAudioFile(%q) = %v, want %v", tc.given, got, tc.want)
		}
	}
}

func TestIsVideoFile(t *testing.T) {
	testCases := []struct {
		given string // 入力ファイル名
		want  bool   // 期待値
	}{
		// テストケースはここに追加
	}

	for _, tc := range testCases {
		got := IsVideoFile(tc.given)
		if got != tc.want {
			t.Errorf("IsVideoFile(%q) = %v, want %v", tc.given, got, tc.want)
		}
	}
}

func TestIsDocumentFile(t *testing.T) {
	testCases := []struct {
		given string // 入力ファイル名
		want  bool   // 期待値
	}{
		// テストケースはここに追加
	}

	for _, tc := range testCases {
		got := IsDocumentFile(tc.given)
		if got != tc.want {
			t.Errorf("IsDocumentFile(%q) = %v, want %v", tc.given, got, tc.want)
		}
	}
}

func TestIsScriptFile(t *testing.T) {
	testCases := []struct {
		given string // 入力ファイル名
		want  bool   // 期待値
	}{
		// テストケースはここに追加
	}

	for _, tc := range testCases {
		got := IsScriptFile(tc.given)
		if got != tc.want {
			t.Errorf("IsScriptFile(%q) = %v, want %v", tc.given, got, tc.want)
		}
	}
}

func TestIsExecutableFile(t *testing.T) {
	testCases := []struct {
		given string // 入力ファイル名
		want  bool   // 期待値
	}{
		// テストケースはここに追加
	}

	for _, tc := range testCases {
		got := IsExecutableFile(tc.given)
		if got != tc.want {
			t.Errorf("IsExecutableFile(%q) = %v, want %v", tc.given, got, tc.want)
		}
	}
}

func ExampleIsTextFile() {
	fmt.Println(IsTextFile("readme.txt"))
	fmt.Println(IsTextFile("document.md"))
	fmt.Println(IsTextFile("data.csv"))
	fmt.Println(IsTextFile("image.jpg"))
	// Output:
	// true
	// true
	// true
	// false
}

func ExampleIsImageFile() {
	fmt.Println(IsImageFile("photo.jpg"))
	fmt.Println(IsImageFile("image.png"))
	fmt.Println(IsImageFile("logo.svg"))
	fmt.Println(IsImageFile("document.txt"))
	// Output:
	// true
	// true
	// true
	// false
}

func ExampleIsAudioFile() {
	fmt.Println(IsAudioFile("song.mp3"))
	fmt.Println(IsAudioFile("music.wav"))
	fmt.Println(IsAudioFile("sound.ogg"))
	fmt.Println(IsAudioFile("video.mp4"))
	// Output:
	// true
	// true
	// true
	// false
}

func ExampleIsVideoFile() {
	fmt.Println(IsVideoFile("movie.mp4"))
	fmt.Println(IsVideoFile("video.avi"))
	fmt.Println(IsVideoFile("clip.mov"))
	fmt.Println(IsVideoFile("audio.mp3"))
	// Output:
	// true
	// true
	// true
	// false
}

func ExampleIsDocumentFile() {
	fmt.Println(IsDocumentFile("report.pdf"))
	fmt.Println(IsDocumentFile("document.docx"))
	fmt.Println(IsDocumentFile("spreadsheet.xlsx"))
	fmt.Println(IsDocumentFile("script.py"))
	// Output:
	// true
	// true
	// true
	// false
}

func ExampleIsScriptFile() {
	fmt.Println(IsScriptFile("script.py"))
	fmt.Println(IsScriptFile("build.sh"))
	fmt.Println(IsScriptFile("app.js"))
	fmt.Println(IsScriptFile("document.pdf"))
	// Output:
	// true
	// true
	// true
	// false
}

func ExampleIsExecutableFile() {
	fmt.Println(IsExecutableFile("program.exe"))
	fmt.Println(IsExecutableFile("library.dll"))
	fmt.Println(IsExecutableFile("binary.bin"))
	fmt.Println(IsExecutableFile("script.py"))
	// Output:
	// true
	// true
	// true
	// false
}
