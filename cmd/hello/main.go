package main

import "fmt"

func main() {
	for range 2 {
		fmt.Println(greeting("friend"))
	}
}

func greeting(name string) string {
	return "hello from qegg pr mode, " + name
}
