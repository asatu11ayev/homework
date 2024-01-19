package main

import "fmt"

func main (){
     var a int
	 var b int
	 var g string
	 fmt.Println("Brinchi sonni kiriting")
	 fmt.Scan(&a)
	 fmt.Println("Ikkinchi sinni kiriting")
	 fmt.Scan(&b)
	 fmt.Println("Aperator kiriting")
	 fmt.Scan(&g)

	 if g == "+"{
       fmt.Println(a+b)
	 }else if g == "-"{
		fmt.Println(a-b)
	 }else if g == "*"{
		fmt.Println(a*b)
	 }else if g == "/"{
		fmt.Println(a/b)
	 }

    //  b := (111%11 == 1 && 1*0 == 0) || !(0/1 == 1) || (3695/2 == 1847)
	//  fmt.Print(b)

    // a := (!(true && false) || !(true || false))&&((true && false) && !(true || false))
	// fmt.Print(a)

    // a := 5
	// b := 5
	// d := 5
	// c := 5
	// p := a+b+c+d
	// r := 2*(p+p)
	// fmt.Print(r)

    //  a := 6
	//  b := 7
	//  c := 8
	//  d := 9
	//  p := a+b+c+d
	//  r := 2*(p+p)
	//  fmt.Print(r)

	// a := 12
	// d := a*a
	// fmt.Print(d)

	// var d float64  = 3.14
	// var a float64  = 3.14
	// c := d*a
	// fmt.Print(c)

    // a := 6
	// s := a*a
	// fmt.Print(s)

    // a := 6
	// b := 9
    // c := 3
	// s := a+b>c && b+c>a
    // fmt.Print(s)

	// a := 5
	// b := 7
	// a,b=b,a
	// fmt.Print(a,b)
}