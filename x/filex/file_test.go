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
		filename := "test_file.txt"
		err := WriteFile(filename, data, 0666)
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
		err := WriteFile(filename, data, os.ModePerm, AutoCreateParentDirOption)
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
		err := WriteJsonFile(filename, data, os.ModePerm, &JsonOption{})
		require.ErrorContains(t, err, "no such file or directory")
		//err = os.Remove(filename)
		//require.NoError(t, err)
	}

	func() {
		filename := "dir/test_file.json"
		err := WriteJsonFile(filename, data, os.ModePerm, AutoCreateParentDirJsonOption)
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
		err := WriteJsonFile(filename, data, os.ModePerm, MarshalIndentJsonOption)
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

func TestWriteCsvFile(t *testing.T) {
	var err error
	{
		filename := "test.csv"
		// this defines the header value and data values for the new csv file
		headers := []string{"name", "age", "sex"}
		data := [][]string{
			{"Alice", "25", "Female"},
			{"Bob", "30", "Male"},
			{"Charlie", "35", "Male"},
		}
		err = WriteCsvFile(filename, headers, data)
		require.NoError(t, err)
		err = os.Remove(filename)
		require.NoError(t, err)
	}

	{
		filename := "test.csv"
		// this defines the header value and data values for the new csv file
		headers := []string{"name", "desc"}
		data := [][]string{
			{"Alice", "hello, ' world"},   // data contains comma and single quote
			{"Bob", "hello, \" world"},    // data contains comma and double quote
			{"Charlie", "hello, ` world"}, // data contains comma and back quote
		}
		err = WriteCsvFile(filename, headers, data)
		require.NoError(t, err)
		err = os.Remove(filename)
		require.NoError(t, err)
	}

	{
		filename := "dir/test.csv"
		// this defines the header value and data values for the new csv file
		headers := []string{"name", "desc"}
		data := [][]string{
			{"Alice", "hello, ' world"},   // data contains comma and single quote
			{"Bob", "hello, \" world"},    // data contains comma and double quote
			{"Charlie", "hello, ` world"}, // data contains comma and back quote
		}
		err = WriteCsvFile(filename, headers, data)
		require.ErrorContains(t, err, "no such file or directory")
	}

	{
		filename := "dir/test.csv"
		// this defines the header value and data values for the new csv file
		headers := []string{"name", "desc"}
		data := [][]string{
			{"Alice", "hello, ' world"},   // data contains comma and single quote
			{"Bob", "hello, \" world"},    // data contains comma and double quote
			{"Charlie", "hello, ` world"}, // data contains comma and back quote
		}
		err = WriteCsvFile(filename, headers, data, AutoCreateParentDirOption)
		require.NoError(t, err)
		defer func() {
			err = os.Remove(filename)
			require.NoError(t, err)

			err = os.RemoveAll(filepath.Dir(filename))
			require.NoError(t, err)
		}()
	}
}

func TestReadCsvFile(t *testing.T) {
	var err error

	filename := "test.csv"
	// this defines the header value and data values for the new csv file
	headers := []string{"name", "desc"}
	data := [][]string{
		{"Alice", "hello, ' world"},   // data contains comma and single quote
		{"Bob", "hello, \" world"},    // data contains comma and double quote
		{"Charlie", "hello, ` world"}, // data contains comma and back quote
	}
	err = WriteCsvFile(filename, headers, data, AutoCreateParentDirOption)
	require.NoError(t, err)
	defer func() {
		err = os.Remove(filename)
		require.NoError(t, err)
	}()

	rows, err := ReadCsvFile(filename)
	require.NoError(t, err)
	require.Equal(t, 4, len(rows))
	require.Equal(t, 2, len(rows[0]))
	require.Equal(t, 2, len(rows[1]))
	require.Equal(t, 2, len(rows[2]))
	require.Equal(t, 2, len(rows[3]))
	require.Equal(t, "name", rows[0][0])
	require.Equal(t, "Alice", rows[1][0])
	require.Equal(t, "Bob", rows[2][0])
	require.Equal(t, "Charlie", rows[3][0])
	require.Equal(t, "desc", rows[0][1])
	require.Equal(t, "hello, ' world", rows[1][1])
	require.Equal(t, "hello, \" world", rows[2][1])
	require.Equal(t, "hello, ` world", rows[3][1])
}
