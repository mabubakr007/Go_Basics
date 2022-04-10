package main

import "fmt"

func main() {

	grades := [3]int{80, 90, 100}
	fmt.Printf("Grades: %v\n", grades)

	grades2 := [...]int{82, 92, 102}
	fmt.Printf("Grades: %v\n", grades2)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Hassaan"
	fmt.Printf("Students: %v\n", students)

	fmt.Printf("Student#2: %v\n", students[1])

	fmt.Printf("Number of Students: %v\n", len(students))

	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}

	fmt.Println(identityMatrix)

	var idm_mat [2][2]int
	idm_mat[0] = [2]int{1, 0}
	idm_mat[1] = [2]int{0, 1}

	fmt.Println(idm_mat)

	// arrays are considered values in Go
	// in C/C++ a copy of an array actually point to the same data
	// but in Go it actually makes a new copy of the old data
	// in passing arrays via functions Go will make copy of the entire array
	// it is better to use pointers instead
	orig_a := [...]int{1, 2, 3}
	copy_a := orig_a
	copy_a[1] = 5
	fmt.Println(copy_a)

	// array are not reference but value types so
	c := &orig_a // not pointing to the same data
	c[1] = 6
	fmt.Println(c)
	// slices -> because array length has to be known at compile time
	// slices are always backed by arrays that are handling the data

	slice := []int{1, 2, 3}
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// slice are reference type so no need to pass the address '&' operator
	d := slice
	d[1] = 7
	fmt.Println(d)

	a_s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b_s := a_s[:]   // slice of all elements
	c_s := a_s[3:]  // from 4th onwards
	d_s := a_s[6:]  // from 6th onwards
	e_s := a_s[3:6] // 4,5,6th elements
	// slicing operators can work on both arrays or slices
	a_s[5] = 42

	// all slices point to the same underlying data hence all will reflect this change
	fmt.Println(a_s)
	fmt.Println(b_s)
	fmt.Println(c_s)
	fmt.Println(d_s)
	fmt.Println(e_s)

	s_a := []string{"This", "is", "an", "array", "of", "strings"}
	fmt.Println(s_a)

	// slicing can also be done using a make function

	// every value in go is initialized to 0 and so are arrays and slices
	make_slice := make([]int, 3)
	fmt.Println(make_slice)
	fmt.Println(len(make_slice))

	// make a slice with an array of underlying 100 elements
	// this is because unlike array slices do not have to be of a fixed size
	// we can add new elements to a slice during its lifetime
	make_slice2 := make([]int, 3, 100)
	fmt.Println(cap(make_slice2))

	make_slice = append(make_slice, 1)
	fmt.Println(make_slice)
	fmt.Println(len(make_slice)) // length and capacity is increased by appending
	// go copies all of the data to another array with more space for the data
	// can go bad if we use for very large array -> use the three arguement make array functions instead
	make_slice2 = append(make_slice2, 1, 2, 3, 4, 5, 6)
	fmt.Println(make_slice2)

	// can not append a integer array to the slice so we can use the concept
	// of spreading as it is available in JS
	make_slice = append(make_slice, []int{10, 11, 12}...)
	fmt.Println(make_slice)

	// slice allows us to push elements using append function
	// pop can be done using the shift function

	b_m := make_slice[:len(make_slice)-1] // remove from the end
	fmt.Println(b_m)
	c_m := append(make_slice2[:2], make_slice2[3:]...)
	fmt.Println(c_m)
}
