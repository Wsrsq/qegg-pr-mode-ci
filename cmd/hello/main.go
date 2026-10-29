package main

import (
	"fmt"
	"os"
)

// version is overridden at build time so a binary can identify its release.
var version = "dev"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Println(version)
		return
	}
	name := "friend"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Println(greeting(name))
	fmt.Println(farewell(name))
}

func greeting(name string) string {
	return "hello from qegg pr mode, " + name + "!"
}

func farewell(name string) string {
	return "goodbye, " + name + "!"
}
