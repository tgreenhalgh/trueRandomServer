package main

import "fmt"

type person struct {
	fname string
	lname string
}

// (p person) attaches this func to the struct `person`
func (p person) speak() {
	fmt.Println(p.fname, `says, "Good morning!"`)
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)

}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	xi := []int{2, 4, 5, 4, 3}
	fmt.Println(xi)

	m := map[string]string{
		"Name": "Thomas",
		"Age":  "50",
	}
	fmt.Println(m)

	p1 := person{
		fname: "Thomas",
		lname: "Greenhalgh",
	}
	fmt.Println(p1)
	// p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}

	// sa1.speak()
	// sa1.person.speak()
	saySomething(p1)
	saySomething(sa1)
}
