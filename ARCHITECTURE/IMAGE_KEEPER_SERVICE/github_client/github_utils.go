package github_client

import (
	"io"
	"os"
)

func createFile(filename string, r io.Reader) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, r)
	if err != nil {
		return err
	}

	return nil
}
