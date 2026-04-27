package main

import "fmt"

func sum(  x , y int) int {
	return x + y
}

//go tu hieu x va y deu la int

func ham() {
	a := 5
	b := 10
	res := sum(a,b)
	fmt.Println("Tong cua", a, "va ", b, "la:", res)
}