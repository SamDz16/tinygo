package main

import (
	"strings"
)

// Declare a main function, this is the entrypoint into our go module
// That will be run. In our example, we won't need this
func main() {}

// This exports an add function.
// It takes in two 32-bit integer values
// And returns a 32-bit integer value.
// To make this function callable from JavaScript,
// we need to add the: "export add" comment above the function
//export add
func add(x int, y int) int {
	return x + y
}

var alphabet = "abcdefghijklmnopqrstuvwxyz,! "

func IntToArrayString(tab []int) string {
	var str string

	for i := 0; i < len(tab); i++ {
		str += alphabet[tab[i] : tab[i]+1]
	}
	return str
}

func StringToIntArray(str string) []int {

	var tab [100]int

	for i := 0; i < len(str); i++ {
		tab[i] = strings.Index(alphabet, str[i:i+1])
	}

	return tab[:]
}

// //export hello
// func hello(tab []int) []int {

// 	return tab[:]
// }

//export hello
func hello(char int) int {

	return char
}
