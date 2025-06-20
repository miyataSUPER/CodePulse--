package main

import (
	"io/fs"
	"os"
	"strings"
)

// FileInfo holds file type, size in bytes, and token count.
type FileInfo struct {
	Type   string
	Size   int64
	Tokens int
}

// GetFileType returns the human-readable category for the file based on its extension.
func GetFileType(path string) string {
	switch {
	case IsTextFile(path):
		return "テキスト"
	case IsImageFile(path):
		return "画像"
	case IsAudioFile(path):
		return "音声"
	case IsVideoFile(path):
		return "動画"
	case IsDocumentFile(path):
		return "ドキュメント"
	case IsScriptFile(path):
		return "スクリプト"
	case IsExecutableFile(path):
		return "実行ファイル"
	case IsArchiveFile(path):
		return "アーカイブ"
	default:
		return "不明"
	}
}

// GetFileInfo reads the file at the given path and returns its type, size, and token count.
func GetFileInfo(path string) (*FileInfo, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	tokens := 0
	if len(data) > 0 {
		tokens = len(strings.Fields(string(data)))
	}

	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	return &FileInfo{
		Type:   GetFileType(path),
		Size:   info.Size(),
		Tokens: tokens,
	}, nil
}

// helper for tests to create a file with content.
func createTempFile(dir, pattern, content string) (string, fs.FileInfo, error) {
	file, err := os.CreateTemp(dir, pattern)
	if err != nil {
		return "", nil, err
	}
	if _, err := file.WriteString(content); err != nil {
		file.Close()
		os.Remove(file.Name())
		return "", nil, err
	}
	if err := file.Close(); err != nil {
		return "", nil, err
	}
	info, err := os.Stat(file.Name())
	return file.Name(), info, err
}
