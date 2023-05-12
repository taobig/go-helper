package io

import (
	"os"
	"path/filepath"
)

func PathExist(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return true, nil
	} else if os.IsExist(err) {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

func CreateParentDir(filePath string) error {
	dir := filepath.Dir(filePath)
	// If path is already a directory, MkdirAll does nothing and returns nil.
	err := os.MkdirAll(dir, os.ModePerm)
	return err
}
