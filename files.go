package frutils

import (
	"os"
	"strings"
)

func CreateFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = file.WriteString(data)
	return err
}

func GetFilenameExtension(filename string) string {
	splitted := strings.Split(filename, ".")
	return splitted[len(splitted)-1]
}
