package main

import (
	"fmt"
	"strings"
)

func thuchanh() {
	const programme = "Gioi thieu ban than"

	var fullname string = "Tran Quang Huy"
	age := 19
	email := "quanghuyy0411@gmail.com"
	hobbies := []string{"code", "nghe nhac", "cau long"}

	fmt.Println(strings.Repeat("=", 40))
	fmt.Println(programme)
	fmt.Println(strings.Repeat("=", 40))

	fmt.Println("Ten:", fullname)
	fmt.Println("Tuoi:", age)
	fmt.Println("Email:", email)
	fmt.Println("So thich:", hobbies)

	//sothich
	fmt.Println("So thich:")
	for i, hobby := range hobbies {
		fmt.Println(i+1, "-", hobby)
	}

	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("Rất vui được làm quen với mọi người!")
	fmt.Println(strings.Repeat("=", 40))
}
