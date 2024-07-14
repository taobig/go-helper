package io

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
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
		require.NoError(t, err)

		if !bytes.Equal(content, expected) {
			t.Fatal("WriteFile != ReadFile")
		}

		err = AppendFile(path, " world\n", 0644)
		require.NoError(t, err)
		content, err = ReadFile(path)
		require.NoError(t, err)
		require.Equal(t, "hello world\n", string(content))

		//t.Logf("file content:%v", string(content))
		err = os.Remove(path)
		require.NoErrorf(t, err, "remove(%v) error failed", path)
	}
}

func TestCreateFileParentDir(t *testing.T) {
	filepath := os.TempDir() + "/" + strconv.FormatInt(time.Now().Unix(), 10) + "/a/b/c/d"
	filepath2 := filepath + "/e"
	{
		err := CreateParentDir(filepath)
		if err != nil {
			t.Fatalf("create dir a/b/c error: %v", err)
		}

		err = WriteFile(filepath, []byte("content"), 0644)
		if err != nil {
			t.Fatalf("create file %s error: %v", filepath, err)
		}
		defer func() {
			os.Remove(filepath)
		}()

		err = CreateParentDir(filepath2)
		if err != nil {
			// expected
		} else {
			t.Fatal("CreateParentDir without error")
		}
	}
}
