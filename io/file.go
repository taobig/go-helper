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

func AppendFile(filename string, content string, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perm)
	if err != nil {
		return err
	}
	_, err = f.WriteString(content)
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}
