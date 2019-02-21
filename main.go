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

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	muhammad := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil.",
	}

	sages := []sage{buddha, gandhi, mlk, jesus, muhammad}

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
