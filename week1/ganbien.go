package main

import "fmt"

func gan() {
	var ho string = "Tran Quang" // khai bao bien voi kieu string
	name:= "Huy"				// khai bao ngan gon
	age := 19
	fmt.Println("Xin chao, toi ten la", ho, name, "va toi", age, "tuoi");
	const sothich = "code" + ", " + "nghe nhac"		//khai bao voi hang so, khong the thay doi gia tri sau khi da gan
	const sothich2 = "cau long"
	fmt.Println("So thich cua toi la:",sothich + " va " + sothich2);
	
}