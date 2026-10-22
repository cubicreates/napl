package main

import (
	"fmt"
	"strconv"
)




func main(){
	v, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println("not an integer")
		return
	}
	fmt.Println(v)
}