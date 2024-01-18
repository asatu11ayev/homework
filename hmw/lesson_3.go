package main

import "fmt"

func main() {




	var name string
	fmt.Scan(&name)

	var age int
	_,err :=fmt.Scan(&age)
	fmt.Println(err)

	var work string
	fmt.Scan(&work)
	fmt.Printf("My name:%s\n My age:%d\n My work:%s", name, age, work)

	formatedStr :=fmt.Sprintf("My name:%s,My age:%d, My work:%s", name, age, work)
	fmt.Print("Creat:",formatedStr)
}


