package api

import (
	"log"
	"os"
	"path"
)

// Exists .
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir .
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile .
func IsFile(path string) bool {
	return !IsDir(path)
}

// Basedir root Dir
func Basedir(staticDir, tmpDir string) (baseDir string, staticPath string, tempPath string) {
	baseDir, err := os.Getwd()
	if err != nil {
		log.Println(err.Error())
		return "", "", ""
	}
	staticPath = path.Join(baseDir, staticDir)
	// temp files directory
	tempPath = path.Join(baseDir, tmpDir)
	return
}
