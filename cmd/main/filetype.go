package main

import "strings"

func hasExtension(path string, exts []string) bool {
	for _, ext := range exts {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}
	return false
}

func IsTextFile(path string) bool {
	exts := []string{".txt", ".md", ".csv", ".log"}
	return hasExtension(path, exts)
}

func IsImageFile(path string) bool {
	exts := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg", ".webp"}
	return hasExtension(path, exts)
}

func IsAudioFile(path string) bool {
	exts := []string{".mp3", ".wav", ".ogg", ".flac", ".aac"}
	return hasExtension(path, exts)
}

func IsVideoFile(path string) bool {
	exts := []string{".mp4", ".avi", ".mov", ".wmv", ".flv", ".mkv"}
	return hasExtension(path, exts)
}

func IsDocumentFile(path string) bool {
	exts := []string{".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx"}
	return hasExtension(path, exts)
}

func IsScriptFile(path string) bool {
	exts := []string{".sh", ".bat", ".ps1", ".py", ".rb", ".pl", ".js", ".ts"}
	return hasExtension(path, exts)
}

func IsExecutableFile(path string) bool {
	exts := []string{".exe", ".dll", ".so", ".bin", ".app"}
	return hasExtension(path, exts)
}

func IsArchiveFile(path string) bool {
	exts := []string{".zip", ".tar", ".gz", ".bz2", ".rar", ".7z", ".xz", ".tar.gz", ".tar.bz2", ".tar.xz", ".tgz", ".tbz2", ".lz", ".lzma", ".z", ".cab", ".arj", ".war", ".ear"}
	return hasExtension(path, exts)
}
