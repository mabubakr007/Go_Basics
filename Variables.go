package main

import (
	"fmt"
	"strconv"
)

// declaring variables at the package level.
// you have to use the complete syntax for declaration

var check int = 42

// if we want to declare multiple variables we can do so at a package level
// and inside a single var code block

// variable naming convention for visibility scope

// lower case first letter means it is scoped to the package level
// anything using package can not see or use it
var ilp int = 10

// upper case first letter will be exported at package level to outside world
// so anything using this package can make use of this variable
var ILP int = 10

var (
	actorName     string = "Adam Scott"
	companionName string = "Britt Lower"
	season        int    = 1
)

var (
	counter int = 0
)

func main() {

	// the length of the variable name should be proportional to its usage/ lifespan
	// all acronyms should be upper case -> HTTP, URL

	// declare and give the value to a variable
	var i int
	i = 42
	fmt.Println(i)
	i = 27
	fmt.Println(i)
	// assign value at time of declaration

	// gives a bit of control because if you declare 99
	// then compiler assumes value is of integer type but
	// 99.0 it will assume float64
	// if you want to have float32 you can use this declaration type
	var j int = 45
	fmt.Println(j)

	// assign a value and compiler will determine what value it needs to have
	z := 42
	fmt.Println(z)

	// allows us to define the a formatted print statement
	// %v for value and %T for type
	fmt.Printf("%v, %T", j, j)

	fmt.Printf("%v \n", check)

	// redeclaration of variables (check declared twice)
	var check int = 32
	// error no new variables
	// can not redefine the variable but assign it a value

	// check := 22

	// this however will work
	check = 22

	// the inner most variable takes precedence
	// function level check takes precedence over package level check
	// this is known as shadowing variable
	fmt.Println(check)

	// local variables in go always have to be used if they are declared
	// var no_use = 10 -> get error var declared but not used

	// this variable is scoped sto the block and never visible outside
	var ilp int = 11
	fmt.Println(ilp)

	// variable type conversion
	var diff float64
	// float64 is a conversion function
	diff = float64(i)

	// conver float32 to int -> lose info

	// var diff2 float64 = 64.5
	// the statement below  will result in error and go will not allow it
	// diff2 = int32(diff2)

	fmt.Printf("%T \n", diff)

	var k1 string
	k1 = string(j)

	// returns value as * meaning compiler sees this as a string of bytes
	// can use the strconv library instead
	fmt.Printf("%v, %T\n", k1, k1)

	k1 = strconv.Itoa(j)
	// converts to string properly
	fmt.Printf("%v, %T\n", k1, k1)
}
