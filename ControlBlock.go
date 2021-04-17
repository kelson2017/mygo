package main

import "fmt"

func main() {

	var a int
	a = 3
	switch a {
	case 0:
		fmt.Println("this is 0!")
	case 1:
		fmt.Println("this is 1.")
	case 2, 3:
		fmt.Println("this is 2, or 3")
	default:
		fmt.Println("default")
	}

}
