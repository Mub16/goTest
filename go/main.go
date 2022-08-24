package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

//////////FUNCTION//////////////

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
	men := User{Name: "Steave", Age: 20, Money: -228, Happines: 0.6, Avg_grades: 0, Hobbies: []string{"Video Games", "Mining", "Card colection"}}
	//men.setNewName("Bob")
	//fmt.Fprintf(w, men.getAllInfo())
	tmpl, _ := template.ParseFiles("HTML/main.html", "HTML/contact.html")
	//tmpl0, _ := template.ParseFiles("HTML/contact.html")
	tmpl.Execute(w, men)
	//tmpl0.Execute(w, men)
}
func (user0 User) getAllInfo() string {
	return fmt.Sprintf("Name: %s. \n Age: %d \n Money:%v", user0.Name, user0.Age, user0.Money)
}
func (user0 *User) setNewName(NewName string) {
	user0.Name = NewName
}
func err_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "it'page err 404")
}

// ////////TYPE//////////////
type User struct {
	Name                        string
	Age                         uint32
	Avg_grades, Happines, Money float64
	Hobbies                     []string
}
