package file

import (
	"io"
	"os"
)

func ReadFileContent(filePath string) ([]byte, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(pwd + filePath)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = file.Close()
	}()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}
