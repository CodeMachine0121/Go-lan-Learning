package main
import(
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request){
	r.ParseForm() // 解析參數
	
	fmt.Println(r.Form)
	fmt.Println("Path: ", r.URL.Path)
	fmt.Println("Scheme: ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"]) // 上傳的變數名稱
		
	for k, v := range r.Form{
		fmt.Println("################################")
		fmt.Println("Key: ",k)
		fmt.Println("Value: ", strings.Join(v,""))
	}
	fmt.Println("################################")

	fmt.Fprintln(w, "Hello astaxie!") // 寫入到 w 傳給客戶端
}

func main(){
	http.HandleFunc("/", sayhelloName) // 設定存取路由
	err:=http.ListenAndServe(":9090", nil) // 設定監聽的 port
	
	if err != nil{
		log.Fatal("ListenAndServe: ",err)
	}
}