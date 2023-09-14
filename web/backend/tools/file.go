package tools

import "os"

func Write(context []byte, path string) error {
	err := os.WriteFile(path, context, 0644)
	if err != nil {
		return err
	}
	return nil
}

func CreateDir(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	return nil
}
