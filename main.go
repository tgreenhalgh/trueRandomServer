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

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

func init() {
	// Must does error checking, ParseGlob gets all the matching files
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	// anonymous struct (i.e. not declared type blah struct{})
	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating the file", err)
	}
	defer nf.Close()

	err = tpl.ExecuteTemplate(nf, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
