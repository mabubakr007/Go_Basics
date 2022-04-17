package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int // field name and type of the field
	actorname  string
	companions []string
}

// uppercase exported and if lower case so only remains in the current package
type Doctor2 struct {
	number     int
	actorname  string
	companions []string
	episodes   []string
}

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal // embedding an animal inside a Bird
	// animal Animal -> this would technically be correct but now animal would be a field of the Bird struct
	SpeedKPH float32
	CanFly   bool
}

// tags in structs

type Tank struct {
	Name string `required max: "100"`
}

func main() {
	aDoctor := Doctor{
		number:    1,
		actorname: "Hugh Laurie",
		companions: []string{ // initializing a slice so have to specify the type
			"Omar Epps",
			"Olivia Wilde",
			"Kal Penn",
			"Robert Sean Leonard"},
	}
	fmt.Println(aDoctor)

	// dot syntax to access a particular value of the struct
	fmt.Println(aDoctor.number)

	fmt.Println(aDoctor.companions[1])

	//this style can lead to errors if we redefine our struct and add or reposition the field in the struct
	aDoctor2 := Doctor{
		1,
		"Robert Sean Leonard",
		[]string{
			"Omar Epps",
			"Hugh Laurie",
		},
	}
	fmt.Println(aDoctor2)

	// with this style of declaration can even skips some fields and not initialize values for it
	anewDoctor := Doctor2{
		number:    1,
		actorname: "Omar Epps",
		companions: []string{
			"Olivia Wilde",
			"Hugh Laurie",
		},
	}
	fmt.Println(anewDoctor)

	// anonymous struct
	anonDoctor := struct{ name string }{name: "John Doe"}
	fmt.Println(anonDoctor)

	anotherDoctor := anonDoctor
	anotherDoctor.name = "Tim Apple"

	//unlike slices and maps; structs are value tyoes rather than reference types so changing one does not effect the other
	fmt.Println(anonDoctor)

	miscDoctor := &anonDoctor
	miscDoctor.name = "Jane Doe"

	// passing the reference allows us to change the underlying value/ data
	fmt.Println(anonDoctor)

	// Go doesnot have/ support Inheritance
	// Instead it has the concept of Composition through embedding -> can use it when Polymorphism is not important

	// Bird is an animal x
	// Bird has animal (like characteristics)

	b := Bird{}
	b.Name = "Ostrich"
	b.Origin = "Africa"
	b.SpeedKPH = 50.32
	b.CanFly = false
	fmt.Println(b)
	fmt.Println(b.Name)

	b2 := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}

	fmt.Println(b2.Name)

	// access the tag
	t := reflect.TypeOf(Tank{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
