package main

import (
	"log"
	"os"
	"text/template"
)

// pacage level scope
var tpl *template.Template

func init() {
	// Must does error checking, ParseGlob gets all the matching files
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	sages := []string{"Ghandi", "MLK", "Buddha", "Jesus", "Muhammad"}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating the file", err)
	}
	defer nf.Close()

	err = tpl.ExecuteTemplate(nf, "tpl.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
