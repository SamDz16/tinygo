package main

import (
	"fmt"
	"strings"
)

// Create a byte (uint8, not Go byte) buffer, which will be available in Wasm Memory.
// We can then share this buffer with JS and Wasm.
const BUFFER_SIZE int = 2

var buffer [BUFFER_SIZE]uint8

// Declare a main function, this is the entrypoint into our go module
// That will be run. In our example, we won't need this
func main() {
	fmt.Println(hello())
}

// Function to return a pointer (Index) to our buffer in wasm memory
//export getWasmMemoryBufferPointer
func getWasmMemoryBufferPointer() *[BUFFER_SIZE]uint8 {
	return &buffer
}

// Function to store the passed value at index 0,
// in our buffer
//go:export storeValueInWasmMemoryBufferIndexZero
func storeValueInWasmMemoryBufferIndexZero(value uint8) {
	buffer[0] = value
}

// Function to read from index 1 of our buffer
// And return the value at the index
//go:export readWasmMemoryBufferAndReturnIndexOne
func readWasmMemoryBufferAndReturnIndexOne() uint8 {
	return buffer[1]
}

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

func IntArrayToString(tab []int) string {
	var str string

	for i := 0; i < len(tab); i++ {
		str += alphabet[tab[i] : tab[i]+1]
	}
	return str
}

func StringToIntArray(str string) []int {

	tab := []int{}

	for i := 0; i < len(str); i++ {
		tab = append(tab, strings.Index(alphabet, str[i:i+1]))
	}

	return tab[:]
}

// //export hello
// func hello(tab []int) []int {

// 	return tab[:]
// }

//export hello
func hello() *string {
	var str *string

	tmp := "Hello, World!"
	str = &tmp

	return str
}
