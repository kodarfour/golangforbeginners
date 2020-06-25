package main

import "fmt"

func main() {
	x := 1
	y := 5

	if x > 0 {
		fmt.Println(x)
	} else {
		fmt.Println(y)
	}

	///for init; condition; post { } - a loop using the syntax borrowed from C;
	///for condition { } - a while loop, and;
	///for { } - an endless loop.

	upToten := 0
	for i := 0; i <= 10; i++ {
		count := upToten + i
		fmt.Printf("%v ", count)
	}

	//if using a special character they can hold a value
	//over 1 (ex. ő is 2)
	for num, char := range "Hello There sir" {
		fmt.Printf("\n %c is the character, %d is the position \n", char, num)
	}

	fmt.Println()

	lmao := 0

	switch lmao {
	case 0:
		fmt.Println("test")
		fallthrough
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("five")
	case 4:
		break
	default:
		fmt.Println("three")
	}

	//declaring arrays
	var arr [10]int
	arr[0] = 42

	//to have numbers already inside
	arr2 := [4]int{1, 2, 3, 4}

	fmt.Println("\nArray Dropdown")

	for i := 0; i < 4; i++ {
		fmt.Printf("%d \n", arr2[i])
	}

	//slices are a more lightweight and customizable array

	arr2slice := arr2[1:3] //includes the low (1) and excludes the high (4)

	fmt.Printf("This is the slice %v\n", arr2slice)

	// u can make slices larges by utilizing append and copy

	arr2slice2 := append(arr2slice, 5, 7)

	fmt.Printf("This is the slice now %v\n", arr2slice2)

	arr2slice3 := append(arr2slice, arr2slice2...) //the dots specify your appending another slice not one value
	fmt.Printf("This is the slice now %v\n", arr2slice3)

	testmake := make([]int, 5) //creates empty slice of that size

	fmt.Printf("This is the slice now %v\n", testmake)

	n1 := copy(testmake, arr2slice2[0:]) // copies arr2slices data into testmake and assigns n1 to
	// the number of values copied into the slice "arr2slice2"

	fmt.Printf("This is the slice now %v and n1 is %v\n", testmake, n1)

	//maps
	//declaring maps works map[firstype]secondtype
	//u must use make to create an empty map

	dcspeedsters := map[string]int{
		"Wally West": 1, "Barry Allen": 2,
		"Reverse-Flash": 3, "Superman": 4,
		"The Black Racer": 5, "Cheetah": 6,
		"Wonder Woman": 7, "God Speed": 8,
		"Shazam": 9, "Kid Flash": 10,
	}

	fmt.Println("\nUnedited Map")

	for name, num := range dcspeedsters {
		//in range loops, u can use "_" if you are trying to focus on one  variable
		fmt.Printf("%s is the number %d speedster in the DC Universe.\n", name, num)
	}

	//adding new hero
	dcspeedsters["Quicksilver"] = 11

	//deleting hero
	delete(dcspeedsters, "Cheetah")
	//to test for existence you would use the following: value, present := monthdays["Jan"]
	//It’s more Go like to name present “ok”, and use: v, ok := monthdays["Jan"].
	//In Go we call this the “comma ok” form
	//in general the syntax delete(m, x) will delete the map entry retrieved by the expression m[x].

	fmt.Println("\nEdited Map")

	for name, num := range dcspeedsters {
		//in range loops, u can use "_" if you are trying to focus on one  variable
		fmt.Printf("%s is the number %d speedster in the DC Universe.\n", name, num)
	}

}
