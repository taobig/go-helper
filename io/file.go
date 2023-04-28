package io

import (
	"io/ioutil"
	"os"
)

// Deprecated 使用原生的ioutil.ReadFile(<Go1.16) or os.ReadFile(>=Go1.16)
func ReadFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	return data, err
}

// Deprecated 使用原生的ioutil.WriteFile(<Go1.16) or os.WriteFile(>=Go1.16)
func WriteFile(filePath string, content []byte, perm os.FileMode) error {
	err := ioutil.WriteFile(filePath, content, perm)
	return err
}

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
