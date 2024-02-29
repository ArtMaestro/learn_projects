package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name      string
	Age       uint16
	Money     int16
	Service   string
	Time      uint16
	Master    string
	Happiness uint16
	Hobbies   []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money equal: %d", u.Name, u.Age, u.Money)
}
func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {

	male := User{
		Name:      "Gena",
		Age:       27,
		Money:     1400,
		Service:   "fade",
		Time:      60,
		Master:    "Victor",
		Happiness: 5,
		Hobbies:   []string{"footbal", "dkdsdd", "sfsdfqas"}}

	// male.setNewName("Alex")
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, male)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello bro")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	// male := User{name:"Gena", age: 27 , money: 1400, service: "fade", time: 60, master: "Victor", happiness: 5 }
	handleRequest()

}
