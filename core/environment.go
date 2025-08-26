// basic environment checks
package core

import (
	"os"
)

// checks if the current system is an Android system.
func IsAndroid() bool {
	_, err := os.Stat("/system/build.prop")
	return err == nil
}

// checks if the euid of current user is root (0)
func IsRoot() bool {
	euid := os.Geteuid()
	if euid == 0 {
		return true
	} else {
		return false
	}
}

// checks if there is a magisk modules path
func ModulesPathExists() bool {
	_, err := os.Stat("/data/adb/modules")
	return err == nil
}

func UserDirectoryExists() bool {
	_, err := os.Stat("/data/" + ProgramName)
	return err == nil
}

func BuildUserDirectory() {
	mkdirIfNotExists(UserDirectory)
	mkdirIfNotExists(CacheDirectory)
}
