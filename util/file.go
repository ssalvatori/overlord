package util

import "os"

func FileExists(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	f.Close()

	return nil
}
