package main

import (
	"fmt"
	"strconv"
)

//asd
type Contact struct {
	num          int
	firstName   string
	lastName    string
}

func (c Contact) getFirstName() string {
	return c.firstName
}

func (c Contact) getLastName() string {
	return c.lastName
}

func (c Contact) getNum() int {
	return c.num
}

func (c *Contact) setFirstName(firstname string) {
	c.firstName = firstname
}

func (c *Contact) setLastName(lastname string) {
	c.lastName = lastname
}

func (c *Contact) toString() string {
	return c.getFirstName() + " " + c.getLastName() + " " + strconv.Itoa(c.getNum())
}

func main() {

var newContact Contact

newContact.setLastName("Rakhmanov")
newContact.setFirstName("Azizbek")
newContact.num = +998997779977
//fmt.Println(newContact)

fmt.Println(newContact.toString())
}
