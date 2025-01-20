package fileutils

import "os"

func WriteTextFile(filename string, content string) error {
	err := os.WriteFile(filename, []byte(content), 0666)
	return err
}
