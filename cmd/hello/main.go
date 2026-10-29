package main

import (
	"fmt"
	"os"
)

func main() {
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
