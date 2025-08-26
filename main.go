package main

import (
	core "akarapatakis/pie/core"
	"fmt"
	"os"
)

func main() {

	arguments := os.Args[1:]

	if core.IsAndroid() {
		fmt.Println(core.StatusOK, "Running on Android.")
	} else {
		fmt.Println(core.StatusError, "Not an Android device.")
		os.Exit(1)
	}
	if core.IsRoot() {
		fmt.Println(core.StatusOK, "Running as root.")
	} else {
		fmt.Println(core.StatusError, "Not running as root.")
		os.Exit(1)

	}
	if core.ModulesPathExists() {
		fmt.Println(core.StatusOK, "Magisk modules path exists")
	} else {
		fmt.Println(core.StatusWarn, "Magisk modules path does not exist. \nmodule-* packages won't be available.")
		os.Exit(1)

	}
	if core.UserDirectoryExists() {
		fmt.Println(core.StatusOK, "Pie directory exists.")
	} else {
		fmt.Println(core.StatusError, "Pie directory does not exist.")
		fmt.Println(core.StatusWarn, "Creating Pie directory.")
		core.BuildUserDirectory()
	}

	if len(arguments) < 2 {
		fmt.Println(core.StatusError, "No arguments provided.")
		os.Exit(1)
	} else {
		switch args[1] {
		case "-I":
			core.Install(args[2:])
		case "-R":
			core.Uninstall(args[2:])
		}
	}
}
