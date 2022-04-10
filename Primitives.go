package main

import (
	"fmt"
)

func main() {
	// boolean variables can be used as flag
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	// boolean variables can also be used for logical tests
	k := 1 == 1
	m := 1 == 2

	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", m, m)

	// unlike in other languages like C or C++ where initialized variables
	// can have garbage data
	// variables in Go are by default initialized to 0 value

	var l bool
	fmt.Printf("%v, %T\n", l, l)

	// for all numeric values; 0 value is going to be 0 for int it is 0
	// for float it will be 0.0

	// int will be atleast 32bit
	// but we have int8 -128 to 128
	// int16 -32768 to 32767
	// int32 -2147483648 to 2147483647
	// int64 is quintillions large

	var un_n uint16 = 42
	fmt.Printf("%v, %T\n", un_n, un_n)

	// unsigned int will be atleast 32bit
	// but we have uint8 0 to 255
	// int16 0 to 65535
	// int32 0 to 4294967295
	// we do not have any unsigned 64 bit integer in Go

	// Basic Arithmetics Operators

	a := 10 // in binary it is 1010
	b := 3  // in binary it is 0011

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) // arithmetic does not lead directly to type conversion; we are in essence doing integer division
	fmt.Println(a % b)

	var c int = 100
	var d uint8 = 30
	// (c + d)  can not do this because of mismatched types

	// instead we must do explicit type conversion
	fmt.Println(c + int(d))

	// Bit Operators

	fmt.Println(a & b)  // and -> output 0010
	fmt.Println(a | b)  // or -> output 1011
	fmt.Println(a ^ b)  // exclusive or (1 not both) -> output 1001
	fmt.Println(a &^ b) // and not (opposite of or) -> output 1000

	// bit shifting

	e := 8              // 2^3
	fmt.Println(e << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(e >> 3) // 2^3 / 2^3 = 2^0 = 1

	// float32 +/- 1.18e^-38 to +/- 3.4e^38
	// float32 +/- 2.23e^-308 to +/- 1.8e^308

	// Declaring floating point literals

	f := 3.14
	f = 13.7e72 // can use exponential notation in go
	f = 2.1e14
	fmt.Printf("%v, %T\n", f, f)

	f = 10.2
	g := 3.7

	fmt.Println(f + g)
	fmt.Println(f - g)
	fmt.Println(f * g)
	fmt.Println(f / g)
	// % and bitwise operator are not available with floating point

	// we can work with complex numbers in go as well
	// complex 64 or complex 128 -> made up of two float32s or two float64s
	var comp complex64 = 1 + 2i
	fmt.Printf("%v %T\n", comp, comp)

	// operator available to us are the simple operator
	comp1 := 1 + 2i
	comp2 := 2 + 4i

	fmt.Println(comp1 + comp2)
	fmt.Println(comp1 - comp2)
	fmt.Println(comp1 * comp2)
	fmt.Println(comp2 / comp1)

	// two functions called real and imag to access the parts of a complex number
	fmt.Printf("%v %T\n", real(comp), real(comp))
	fmt.Printf("%v %T\n", imag(comp), imag(comp))

	var n_comp = complex(5, 12)
	fmt.Printf("%v %T\n", n_comp, n_comp)

	//
	s := "This is a string"
	fmt.Printf("%v, %T\n", s[2], s[2]) // this output 105, unint8
	// this is because strings are an alias for bytes and

	fmt.Printf("%v, %T\n", string(s[2]), string(s[2]))
	// this will correct it

	// s[2] = "u"  -> this will cause an error as strings are immutable

	s1 := "This is a s1 string"
	s2 := "This is a s2 string"

	fmt.Printf("%v,%T\n", s1+s2, s1+s2)

	// convert string to a collection of bytes
	bites := []byte(s)                  // a collection of unsigned 8 bit integers
	fmt.Printf("%v,%T\n", bites, bites) // this will return the ascii value of all the letters in the strings

	// if we send files/ strings around we can use the collection of bytes to do that

	// Rune datatype -> utf32 -> can be 32 bit long but does not have to be 32 bits
	// any utf8 is also a utf32 value
	// single quotes for runes
	// rune is the same as an int32
	r := 'a'
	fmt.Printf("%v, %T\n", r, r)

	var r1 rune = 'b'
	fmt.Printf("%v, %T\n", r1, r1)

	// for working with runes we use the strings package

}
