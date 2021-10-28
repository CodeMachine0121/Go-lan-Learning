package main
import "fmt"

func main() {
	var x int = 10
	var y int = 20
	var z int = func2(x,y)
	fmt.Println(z)
}


func func2(x int ,y int) int {
	var z int = x+y+3;
	return z;
}


