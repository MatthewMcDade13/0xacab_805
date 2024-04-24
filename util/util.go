package util

import "os"

const DEFAULT_FILE_PERM = 0666

func IsValidFsPath(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}
