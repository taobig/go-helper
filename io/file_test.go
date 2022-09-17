package io

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestPathExist(t *testing.T) {
	t.Parallel()

	path, err := os.UserHomeDir()
	if err != nil {
		t.Errorf("UserHomeDir() error:%v", err.Error())
	}

	{
		expected := true
		actual1, err := PathExist(path)
		if err != nil {
			t.Errorf("TestPathExist(%v) error:%v", path, err.Error())
		} else {
			if actual1 != expected {
				t.Errorf("TestPathExist(%v):%v; expected %v", path, actual1, expected)
			}
		}
	}

	{
		expected := false
		path += fmt.Sprintf("%d", rand.Int63())
		actual2, err := PathExist(path)
		if err != nil {
			t.Errorf("TestPathExist(%v) error:%v", path, err.Error())
		} else {
			if actual2 != expected {
				t.Errorf("TestPathExist(%v):%v; expected %v", path, actual2, expected)
			}
		}
	}

}

func TestWriteFileAndReadFile(t *testing.T) {
	t.Parallel()

	path := "test_file.txt"
	expected := []byte("hello")
	err := WriteFile(path, expected, 0644)
	if err != nil {
		t.Errorf("WriteFile to %v error:%v", path, err.Error())
	} else {
		content, err := ReadFile(path)
		if err != nil {
			t.Errorf("ReadFile(%v) error:%v", path, err.Error())
		} else {
			if !bytes.Equal(content, expected) {
				t.Fatal("WriteFile != ReadFile")
			}
			//t.Logf("file content:%v", string(content))
			err = os.Remove(path)
			if err != nil {
				t.Errorf("remove(%v) error:%v", path, err.Error())
			}
		}
	}

}
