package main

import "fmt"
func main() {
	var x int

	for x=0; x<10; x++ { 

		if x == 4 {
			fmt.Println("continue")
			continue
		}else if  x == 8{
			fmt.Println("break")
			break
		}else {
			fmt.Println("x: ", x)
		}
	}
}