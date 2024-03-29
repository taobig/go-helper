// +build go1.16

package io

import (
	"os"
)

// ReadFile 使用原生的ioutil.ReadFile(<Go1.16) or os.ReadFile(>=Go1.16)
func ReadFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	return data, err
}

// WriteFile 使用原生的ioutil.WriteFile(<Go1.16) or os.WriteFile(>=Go1.16)
func WriteFile(filePath string, content []byte, perm os.FileMode) error {
	err := os.WriteFile(filePath, content, perm)
	return err
}
