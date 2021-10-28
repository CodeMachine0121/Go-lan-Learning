package main
import "fmt"

func main() {

	var x int
	fmt.Scanln(&x) // 輸入
	fmt.Println("x: ",x)
	fmt.Println("x(mem addr): ", &x) //輸出記憶體位址


	if x > 100{
		x+=100
	}else if x<50{
		x-=50
	}else{
		x+=x
	}

	fmt.Println("x: ", x)
	
}