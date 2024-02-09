package main

import (
	"fmt"
)

var pl = fmt.Println

type customer struct {
	name    string
	address string
	balance float64
}

func getCustomerInfo(c customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.balance)
}

func newCustomerAdd(c *customer, address string) {
	c.address = address
}

type rectangle struct {
	length, height float64
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}

type contact struct {
	firstName string
	lastName  string
	phone     string
}

type business struct {
	name    string
	address string
	contact
}

func (b business) info() {
	fmt.Printf("Contact at %s is %s %s \n", b.name, b.contact.firstName, b.contact.phone)
}

func main() {
	var newCustomer customer
	newCustomer.name = "Tom Smith"
	newCustomer.address = "5 Main st"
	newCustomer.balance = 234.56

	getCustomerInfo(newCustomer)

	newCustomer2 := customer{"Sally Smith", "123 main", 0}
	pl(newCustomer2)

	rectangle1 := rectangle{10, 15}
	pl(rectangle1.Area())

	contact1 := contact{
		"James", "Wang", "0612345678",
	}
	business1 := business{
		"ABC plumbing", "234 north street", contact1,
	}
	business1.info()

}
