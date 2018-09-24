package util

import (
	"path"
	"runtime"
)

func GetCurrentFile() string {
	_, filename, _, _ := runtime.Caller(1)
	return filename
}

func GetCurrentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func GetRootDir() string {
	currentDir := GetCurrentDir()
	return path.Dir(currentDir)
}
