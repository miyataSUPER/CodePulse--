package main

import "strings"

func ext(path string) string {
	if i := strings.LastIndex(path, "."); i >= 0 {
		return strings.ToLower(path[i+1:])
	}
	return ""
}

func IsTextFile(path string) bool {
	switch ext(path) {
	case "txt", "md", "csv", "log":
		return true
	}
	return false
}

func IsImageFile(path string) bool {
	switch ext(path) {
	case "jpg", "jpeg", "png", "gif", "bmp", "svg", "webp":
		return true
	}
	return false
}

func IsAudioFile(path string) bool {
	switch ext(path) {
	case "mp3", "wav", "ogg", "flac", "aac":
		return true
	}
	return false
}

func IsVideoFile(path string) bool {
	switch ext(path) {
	case "mp4", "avi", "mov", "wmv", "flv", "mkv":
		return true
	}
	return false
}

func IsDocumentFile(path string) bool {
	switch ext(path) {
	case "pdf", "doc", "docx", "xls", "xlsx", "ppt", "pptx":
		return true
	}
	return false
}

func IsScriptFile(path string) bool {
	switch ext(path) {
	case "sh", "bat", "ps1", "py", "rb", "pl", "js", "ts":
		return true
	}
	return false
}

func IsExecutableFile(path string) bool {
	switch ext(path) {
	case "exe", "dll", "so", "bin", "app":
		return true
	}
	return false
}

func IsArchiveFile(path string) bool {
	switch ext(path) {
	case "zip", "tar", "gz", "bz2", "rar", "7z", "xz", "tar.gz", "tar.bz2", "tar.xz", "tgz", "tbz2", "lz", "lzma", "z", "cab", "arj", "war", "ear":
		return true
	}
	return false
}
