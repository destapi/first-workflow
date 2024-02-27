package main

import "fmt"

func main(){
	fmt.Println("Hello World!")

	var x int = 1
	var y int
	var ip *int

	fmt.Println(y)

	ip = &x
	y = *ip

	fmt.Println(y)

	var num int

	fmt.Printf("Number?")

	num, err = fmt.Scan(&num)

	fmt.Printf("Entered num = %d", num)
}