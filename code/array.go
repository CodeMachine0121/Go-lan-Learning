package main
import "fmt"

func main(){
	var ary1 [10]int // 長度可有可無
	var i int
	fmt.Print("Array: [")
	for i = 0; i <10; i++{
		ary1[i]=i*10
		fmt.Print(ary1[i]," ")
	}

	fmt.Println("]\nlength of array: ",len(ary1))
	
	
}