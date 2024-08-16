/***
The aim of this exercise is to build a structured web page, use a template, and fill it with parameters from the
querystring.
***/

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var tmplStr = `
<html>
	<h1>Customer {{ .ID }}
	{{ if .ID }}
		<p>Details</p>
		<ul>
		{{ if .Name }}<li>Name: {{ .Name }}</li>{{ end }}
		{{ if .Surname }}<li>Surname: {{ .Surname }}</li>{{ end }}
		{{ if .Age }}<li>Age: {{ .Age }}</li>{{ end }}
		</ul>
	{{ else }}
		<p>Data Not Available</p>
	{{ end }}
</html>
`

type Customer struct {
	ID      int
	Name    string
	Surname string
	Age     int
}

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query()

	c := Customer{}

	id, ok := val["id"]
	if ok {
		c.ID, _ = strconv.Atoi(strings.Join(id, ","))
	}

	name, ok := val["name"]
	if ok {
		c.Name = strings.Join(name, ",")
	}

	surname, ok := val["surname"]
	if ok {
		c.Surname = strings.Join(surname, ",")
	}

	age, ok := val["age"]
	if ok {
		c.Age, _ = strconv.Atoi(strings.Join(age, ""))
	}

	tmpl, _ := template.New("test").Parse(tmplStr)

	tmpl.Execute(w, c)
}

func main() {
	http.Handle("/", Hello{})
	fmt.Println("Starting server on port :8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
