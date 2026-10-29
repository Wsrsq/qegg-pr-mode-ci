package main

import "fmt"

func main() {
	for range 2 {
		fmt.Println(greeting("friend"))
	}
	fmt.Println(farewell("friend"))
}

func greeting(name string) string {
	return "hello from qegg pr mode, " + name + "!"
}

func farewell(name string) string {
	return "goodbye, " + name + "!"
}
