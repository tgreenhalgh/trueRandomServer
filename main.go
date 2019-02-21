package main

import (
	"log"
	"os"
	"text/template"
)

// pacage level scope
var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	// Must does error checking, ParseGlob gets all the matching files
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating the file", err)
	}
	defer nf.Close()

	err = tpl.ExecuteTemplate(nf, "tpl.gohtml", buddha)
	if err != nil {
		log.Fatalln(err)
	}
}
