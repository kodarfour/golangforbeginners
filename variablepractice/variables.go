package main

import "fmt"

func main() {

	const lol = iota
	//will print zero since its outside the grouping

	const ( //can only be numbers, strings, or booleans
		lol1 = iota
		lol2 = iota
		//iotas increase by one every time its used consecutively
	)

	fmt.Println(lol, lol1, lol2)

	var (
		g int
		m float32
	) // u can group variables when declaring type

	//////////////
	var a int
	var b bool
	a = 15
	b = false
	//you can declare them like this but its not as effective

	//this is a deduced declaration meaning the compiler assumes what type your declaring
	x := false
	l := 94.3

	fmt.Println(a, b, x, l, g, m)

	var s string = "hello world"
	test := "hello"
	d := 'c' //prints ascii value
	//u cannot change strings in go

	fmt.Println(s, test, d)
}
