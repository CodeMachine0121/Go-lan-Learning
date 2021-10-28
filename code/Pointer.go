package main

import "fmt"
func main(){
	var x int =3
	var xPtr *int=&x

	fmt.Println("x addr: ",xPtr)
	fmt.Println("x value: ",*xPtr)
	
}