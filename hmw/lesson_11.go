package main

import "fmt"

type Client struct {
	Id int
	Name string 
	Basket []BasketProduts
  }

  type BasketProduts struct {
	ProductId int
	Quantity int
  }
  
  type Product struct {
	Id int
	Name string
	Price int
  }


 func main() {
	Client (
		Id: 123456,
	    Name: "Umar",
        Basket: []BasketProduts{"banan","olma"},
 )
 }