package main

import "fmt"

func main() {

	// all the keys will be of string type and all value will be integer type -> has to be consistent for all elements in the map
	// all keys should be of data types that can be tested for equality -> string, bool, numeric types, arrays
	// can not use slices, maps and other functions as they can not be tested for equality
	statePops := map[string]int{
		"Sindh":       20,
		"Punjab":      45,
		"Balochistan": 15,
		"Khyber":      10,
	}
	fmt.Println(statePops)

	// m := map[[]int]string{} -> this will give errors

	m := map[[3]int]string{}
	fmt.Println(m) // this will work because it is treating the key as an array

	// can also use make function to make map

	statePop := make(map[string]int) // use this syntax to declare a map that you will fill later on using a loop or something
	fmt.Println(statePop)

	fmt.Println(statePops["Sindh"]) // access a particular value within the map
	statePops["Kashmir"] = 7

	fmt.Println(statePops["Kashmir"])

	fmt.Println(statePops) // will return calues in a different order to what we enter values as maps can not guarantee similar order as arrays and slices

	delete(statePops, "Kashmir")
	fmt.Println(statePops)

	fmt.Println(statePops["Kashmir"]) // return 0 even when the key value pair has been deleted

	pop, ok := statePops["Sndh"]
	fmt.Println(pop, ok) // use this ok check to see that does the value even exist in the map because even if a key doesnot exist the map will return a 0 value for it

	pakpops := statePops
	pakpops["Sindh"] = 21

	fmt.Println(statePops) // since assigning another variable equal to statepops; it is passed as a reference and so changes in one effect the other

}
