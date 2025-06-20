package main

import "strings"

// IsArchiveFile checks if path is an archive file based on extension.
func IsArchiveFile(path string) bool {
	targets := []string{".tar", ".tar.gz", ".tar.bz2", ".zip", ".war", ".ear", ".rar"}
	for _, suffix := range targets {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}

// IsTextFile checks if path is a text file based on extension.
func IsTextFile(path string) bool {
	targets := []string{".txt", ".md", ".csv", ".log"}
	for _, suffix := range targets {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}

// IsImageFile checks if path is an image file based on extension.
func IsImageFile(path string) bool {
	targets := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg", ".webp"}
	for _, suffix := range targets {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}

// IsAudioFile checks if path is an audio file based on extension.
func IsAudioFile(path string) bool {
	targets := []string{".mp3", ".wav", ".ogg", ".flac", ".aac"}
	for _, suffix := range targets {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}

// IsVideoFile checks if path is a video file based on extension.
func IsVideoFile(path string) bool {
	targets := []string{".mp4", ".avi", ".mov", ".wmv", ".flv", ".mkv"}
	for _, suffix := range targets {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}

// IsDocumentFile checks if path is a document file based on extension.
func IsDocumentFile(path string) bool {
	targets := []string{".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx"}
	for _, suffix := range targets {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}

// IsScriptFile checks if path is a script file based on extension.
func IsScriptFile(path string) bool {
	targets := []string{".sh", ".bat", ".ps1", ".py", ".rb", ".pl", ".js", ".ts"}
	for _, suffix := range targets {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}

// IsExecutableFile checks if path is an executable file based on extension.
func IsExecutableFile(path string) bool {
	targets := []string{".exe", ".dll", ".so", ".bin", ".app"}
	for _, suffix := range targets {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}
