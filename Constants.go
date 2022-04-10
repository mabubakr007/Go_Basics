package main

import (
	"fmt"
)

const a int = 14

// untyped constant
const iot = iota

// we can have a const block instead

const (
	ai = iota
	bi = iota
	ci = iota
)

// enumerated constants
const (
	// the iota is tied to the scope of a block
	// so using it in another block resets it
	a2 = iota
	b2
	c2
)

// use constant to check if a value has been assigned to a variable
// by declaring an error iota

const (
	error_spec = iota
	car_spec
	bike_spec
	bus_spec
)

const (
	_ = iota //it specifies that we donot care what value is given here by iota

)

const (
	_ = iota + 5 //it specifies that we donot care what value is given here by iota
	// can be used to have a fixed offset
	catspec
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	// bit shifting 1 places and onwards
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	// contant aren't named all caps but like variables

	// const must be declared at compile time not runtime
	// const check float64 = math.Sin(1.57) -> will throw an error
	// const can not be equal to flags in code
	const myConst int = 42
	fmt.Printf("%v, %T \n", myConst, myConst)

	// const can be made of any primitive types

	// consts can be shadowed
	const a int = 4
	fmt.Printf("%v, %T \n", a, a)

	// const can always be added to a variables
	var b int = 43
	fmt.Printf("%v, %T \n", b+a, b+a)

	// we can make use of the compiler guessing the type of variable
	// by doing the following thing

	const a1 = 43
	var b1 uint16 = 33
	fmt.Printf("%v, %T \n", a1+b1, a1+b1) // the compiler is infact replacing every instance of a with the value stored within it
	fmt.Printf("%v, %T \n", 4+b1, 4+b1)   // this is same as the lines above for the compiler

	fmt.Printf("%v \n", ai)
	fmt.Printf("%v \n", bi)
	fmt.Printf("%v \n", ci)
	// iota is changing value as constants are being assigned

	fmt.Printf("%v \n", a2)
	fmt.Printf("%v \n", b2)
	fmt.Printf("%v \n", c2)
	// even when we dont write iota when declaring b2, c2 the compiler infers the order of assignment and assign values to the variables accordingly

	var spec int
	fmt.Printf("%v \n", spec == car_spec)

	fmt.Printf("%v \n", catspec) // will evaluate to 6 as iota starts at 5

	filesize := 400000000.0
	fmt.Printf("%.2f GB\n", filesize/GB)

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b \n", roles)
	fmt.Printf("Is Admin? %v \n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v \n", isHeadquarters&roles == isHeadquarters)
}
