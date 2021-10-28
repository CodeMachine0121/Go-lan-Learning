package main
import "fmt"

func main(){
	var ary1 [10]int // 長度可有可無
	var i int
	for i = 0; i <10; i++{
		ary1[i]=i*10
	}

	fmt.Println(len(ary1))
	
	
}