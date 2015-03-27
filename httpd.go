package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
)

var (
	templates *template.Template
)

type Person struct {
	Id   int
	Name string
	Age  int
	Desc string
}

func findPerson(userId int) *Person {
	return &Person{Id: 1, Name: "Unknown", Age: 20, Desc: "an unknown person"}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(r.URL.Path[len("/view/"):])
	person := findPerson(userId)
	templates.ExecuteTemplate(w, "view.html", person)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "edit.html", nil)
}

func init() {
	templates, _ = template.ParseFiles("view.html", "edit.html")
}

func main() {
	/*
		http.HandleFunc("/view/", viewHandler)
		http.HandleFunc("/edit/", editHandler)
		http.ListenAndServe(":8080", nil)
	*/

	r := gin.Default()
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		message := "Hello " + name
		c.String(http.StatusOK, message)
	})

	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Params.ByName("name")
		action := c.Params.ByName("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	r.Run(":8081")
}
