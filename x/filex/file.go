package filex

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"path/filepath"
)

type Option struct {
	AutoCreateParentDir bool
	DirPerm             os.FileMode
}

func WriteFile(filename string, data []byte, perm os.FileMode, options ...*Option) error {
	var err error
	var option *Option
	if len(options) > 0 {
		option = options[0]
	} else {
		option = &Option{
			AutoCreateParentDir: false,
		}
	}

	if option.AutoCreateParentDir {
		err = os.MkdirAll(filepath.Dir(filename), option.DirPerm)
		if err != nil {
			return err
		}
	}

	err = os.WriteFile(filename, data, perm)
	if err != nil {
		return err
	}
	return nil
}

type JsonOption struct {
	Option

	UseMarshalIndent    bool
	MarshalIndentPrefix string
	MarshalIndentIndent string
}

func WriteJsonFile(filename string, data interface{}, perm os.FileMode, options ...*JsonOption) error {
	var err error
	var option *JsonOption
	if len(options) > 0 {
		option = options[0]
	} else {
		option = &JsonOption{
			Option: Option{
				AutoCreateParentDir: false,
			},
			UseMarshalIndent: false,
		}
	}
	if option.AutoCreateParentDir {
		err = os.MkdirAll(filepath.Dir(filename), option.DirPerm)
		if err != nil {
			return err
		}
	}

	var jsonEncodedContent []byte
	if option.UseMarshalIndent {
		jsonEncodedContent, err = json.MarshalIndent(data, option.MarshalIndentPrefix, option.MarshalIndentIndent)
	} else {
		jsonEncodedContent, err = json.Marshal(data)
	}
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, jsonEncodedContent, perm)
	if err != nil {
		return err
	}
	return nil
}

func WriteCsvFile(filename string, headers []string, data [][]string, options ...*Option) error {
	var err error
	var option *Option
	if len(options) > 0 {
		option = options[0]
	} else {
		option = &Option{
			AutoCreateParentDir: false,
		}
	}
	if option.AutoCreateParentDir {
		err = os.MkdirAll(filepath.Dir(filename), option.DirPerm)
		if err != nil {
			return err
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	//defer writer.Flush()
	err = writer.Write(headers)
	if err != nil {
		return err
	}
	//for _, row := range data {
	//	err = writer.Write(row)
	//	if err != nil {
	//		return err
	//	}
	//}
	err = writer.WriteAll(data)
	if err != nil {
		return err
	}

	return nil
}
