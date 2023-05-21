package main

import "fmt"

type contact_info struct {
	email string
	zip_code int
}

type person struct {
	first_name string
	last_name string
	contact contact_info
	// contact_info
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// Go is a pass by value language: Copy of the value is passed to the function
// To pass by reference, use pointers
func (p *person) update_name(new_first_name string) {
	// This means we want to manipulate the value the pointer is referencing 
	(*p).first_name = new_first_name
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{first_name: "Alex", last_name: "Anderson"}
	var alex person
	alex.first_name = "Alex"
	alex.last_name = "Anderson"
	// alex.contact = contact_info{email: "alex.anderson@gmail", zip_code: 12345}
	alex.contact.email = "alex.anderson@gmail.com"
	alex.contact.zip_code = 12345
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
	// alex_pointer := &alex
	alex.update_name("Alexis")
	alex.print()
}

// Value types: int, float, string, bool, structs
// Reference types: slices, maps, channels, pointers, functions