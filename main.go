package main

import (
	"fmt"
	"os"
)


func main() {

	arguments := os.Args[1:]

	if IsAndroid() {
		fmt.Println(StatusOK, "Running on Android.")
	} else {
		fmt.Println(StatusError, "Not an Android device.")
		os.Exit(1)
	}
	if IsRoot() {
		fmt.Println(StatusOK, "Running as root.")
	} else {
		fmt.Println(StatusError, "Not running as root.")
		os.Exit(1)

	}
	if ModulesPathExists() {
		fmt.Println(StatusOK, "Magisk modules path exists")
	} else {
		fmt.Println(StatusWarn, "Magisk modules path does not exist. \nmodule-* packages won't be available.")
		os.Exit(1)

	}
	if UserDirectoryExists() {
		fmt.Println(StatusOK, "Pie directory exists.")
	} else {
		fmt.Println(StatusError, "Pie directory does not exist.")
		fmt.Println(StatusWarn, "Creating Pie directory.")
		BuildUserDirectory()
	}

	if len(arguments) > 0 {
		if arguments[0] == InstallCharacter {
			//todo
		}
	}
}
