package main

import (
	"html/template"
	"log"
	"net/http"
)

type IndexData struct {
	Title   string
	Content string
}

func test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	tmpl := template.Must(template.ParseFiles("./index.html"))
	data := new(IndexData)
	data.Title = "HomePage"
	data.Content = "First Page"
	tmpl.Execute(w, data) // 把網頁內容送回去

	//w.Write([]byte()) // 寫入要回傳的東西
}
func main() {
	http.HandleFunc("/", test)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
