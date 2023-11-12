package filex

import (
	"bufio"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFile(t *testing.T) {
	data := []byte("Hello World!")
	{
		filename := "test_file.txt"
		err := WriteFile(filename, data, os.ModePerm)
		require.NoError(t, err)
		err = os.Remove(filename)
		require.NoError(t, err)
	}
	{
		filename := "dir/test_file.txt"
		err := WriteFile(filename, data, os.ModePerm)
		require.ErrorContains(t, err, "no such file or directory")
		//err = os.Remove(filename)
		//require.NoError(t, err)
	}
	{
		filename := "dir/test_file.txt"
		option := &Option{
			AutoCreateParentDir: true,
			DirPerm:             os.ModePerm,
		}
		err := WriteFile(filename, data, os.ModePerm, option)
		require.NoError(t, err)
		{
			err = os.Remove(filename)
			require.NoError(t, err)

			err = os.RemoveAll(filepath.Dir(filename))
			require.NoError(t, err)
		}
	}
}

func TestWriteJsonFile(t *testing.T) {
	data := map[string]interface{}{
		"hello": "world",
		"foo":   "bar",
		"abc":   "xyz",
		"obj":   map[string]string{"a": "b"},
	}
	{
		filename := "test_file.json"
		err := WriteJsonFile(filename, data, os.ModePerm)
		require.NoError(t, err)
		err = os.Remove(filename)
		require.NoError(t, err)
	}

	{
		filename := "dir/test_file.json"
		err := WriteJsonFile(filename, data, os.ModePerm)
		require.ErrorContains(t, err, "no such file or directory")
		//err = os.Remove(filename)
		//require.NoError(t, err)
	}
	{
		filename := "dir/test_file.json"
		jsonOption := &JsonOption{}
		err := WriteJsonFile(filename, data, os.ModePerm, jsonOption)
		require.ErrorContains(t, err, "no such file or directory")
		//err = os.Remove(filename)
		//require.NoError(t, err)
	}

	func() {
		filename := "dir/test_file.json"
		jsonOption := &JsonOption{Option: Option{
			AutoCreateParentDir: true,
			DirPerm:             os.ModePerm,
		}}
		err := WriteJsonFile(filename, data, os.ModePerm, jsonOption)
		require.NoError(t, err)
		defer func() {
			err = os.Remove(filename)
			require.NoError(t, err)

			err = os.RemoveAll(filepath.Dir(filename))
			require.NoError(t, err)
		}()
		f, err := os.Open(filename)
		require.NoError(t, err)
		defer func() {
			err = f.Close()
			require.NoError(t, err)
		}()
		var lines []string
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		require.NoError(t, scanner.Err())
		require.Equal(t, 1, len(lines))
	}()
	func() {
		filename := "dir/test_file.json"
		jsonOption := &JsonOption{
			Option:              Option{AutoCreateParentDir: true, DirPerm: os.ModePerm},
			UseMarshalIndent:    true,
			MarshalIndentPrefix: "",
			MarshalIndentIndent: "  ",
		}
		err := WriteJsonFile(filename, data, os.ModePerm, jsonOption)
		require.NoError(t, err)
		defer func() {
			err = os.Remove(filename)
			require.NoError(t, err)

			err = os.RemoveAll(filepath.Dir(filename))
			require.NoError(t, err)
		}()
		f, err := os.Open(filename)
		require.NoError(t, err)
		defer func() {
			err = f.Close()
			require.NoError(t, err)
		}()
		var lines []string
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		require.NoError(t, scanner.Err())
		require.Equal(t, 8, len(lines))
	}()

}
