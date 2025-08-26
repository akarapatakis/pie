package core

import "os"

func mkdirIfNotExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModeDir|0755)
		return err == nil
	}
	return false
}

// touch = unix command to create a file
func touchIfNotExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_, err = os.Create(path)
		return err == nil
	}
	return false
}
