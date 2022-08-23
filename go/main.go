package main

import (
	"fmt"
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
	fmt.Println("Web server start (функция server_start) server url: http://localhost:" + port)
	http.HandleFunc(home_page_url, WEBgraph)
	http.ListenAndServe(":"+port, nil)
}
func write() {
	fmt.Println("Функция write")
}
func WEBgraph(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HTML PAGE home")
}
