package main
import "fmt"
func main() {
	//fmt.Println("hello world\n")

	var x int;
	x = 3;
	fmt.Println(x)
	x=19
	fmt.Println(x)
	//x="助教很帥" > 型態錯了
	
	var x1 float64 = 3.14159
	fmt.Println(x1)

	var x2 rune = 'a'
	fmt.Println(x2)



}
// ----------------------------------------------------------------

// 撰寫程式 > 建置 (build) > run code
// compile: go build <code file>