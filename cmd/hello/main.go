package main

import "fmt"

func main() {
	for range 2 {
		fmt.Println(greeting())
	}
}

func greeting() string {
	return "hello from qegg pr mode"
}
