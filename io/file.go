package io

import (
	"io/ioutil"
	"os"
)

func ReadFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	return data, err
}

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
