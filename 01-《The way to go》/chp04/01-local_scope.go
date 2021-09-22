package main

var a = "G"

// This is main function in 01-Locak
func main() {
	n() // G
	m() // O
	n() // G
}

// this is function n
func n() { print(a) }

// this is func m
func m() {
	a := "O"
	print(a)
}
