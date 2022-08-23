package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("")
	write()
	server_start("/")
}
func server_start(home_page_url string) {
	var port string
	fmt.Print("Введите порт сервера:")
	fmt.Scan(&port)
	fmt.Println("Web server start, server url: http://localhost:" + port)
	http.HandleFunc(home_page_url, WEBgraph)
	http.HandleFunc("/err", err_page)
	http.ListenAndServe(":"+port, nil)
}
func write() {
	fContent, err := ioutil.ReadFile("txt_info/dino.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fContent))
}
func WEBgraph(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HTML PAGE home")
}
func err_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "it'page err 404")
}
