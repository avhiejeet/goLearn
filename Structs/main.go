package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	abhijit := person{
		firstName: "Abhijeet",
		lastName:  "Rajbanshi",
		contactInfo: contactInfo{
			email:   "avhiejeet@gmail.com",
			zipCode: 44600,
		},
	}

	sushant := person{
		firstName: "Sushant",
		lastName:  "Adhikari",
		contactInfo: contactInfo{
			email:   "devsushant46@gmail.com",
			zipCode: 44600,
		},
	}

	abhijit.updateFirstName("Abhijit")
	sushant.updateLastName("Shrestha")
	abhijit.print()
	sushant.print()
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (pointerToPerson *person) updateLastName(newLastName string) {
	(*pointerToPerson).lastName = newLastName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
