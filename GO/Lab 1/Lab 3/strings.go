package main

import "fmt"

func main() {
	// b:=12
	// c:=12.4
	// d:=true
	// a:="new day"

	// var e int = 34
	// var f float32 = 34.0
	// var g float64 = 34.44
	// var h string = "34"
	// var i bool = false

	fmt.Print("Enter the marks: ")
	var str int
	fmt.Scan(&str)

	if str>=33{
		fmt.Print("You Pass!!")
	}else{
		fmt.Print("You Failed!!")
	}

	fmt.Print("\n")

	var age int

	fmt.Printf("%d", age)


}
