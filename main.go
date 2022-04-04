package main

import "fmt"

const PatternName = "Functional builder Pattern design"

//Person struct
type Person struct {
	// Personal details
	name    string
	address string
	pin     string
	// Job details
	workAddress string
	company     string
	position    string
	salary      int
}

type PersonBuilder struct {
	person *Person
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{
		person: &Person{},
	}
}

func (pb *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*pb}
}

func (pb *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*pb}
}

// Address builder

func (pab *PersonAddressBuilder) At(address string) *PersonAddressBuilder {
	pab.person.address = address
	return pab
}

func (pab *PersonAddressBuilder) WithPostalCode(pin string) *PersonAddressBuilder {
	pab.person.pin = pin
	return pab
}

// Job Builder

func (pjb *PersonJobBuilder) As(position string) *PersonJobBuilder {
	pjb.person.position = position
	return pjb
}

func (pjb *PersonJobBuilder) For(company string) *PersonJobBuilder {
	pjb.person.company = company
	return pjb
}

func (pjb *PersonJobBuilder) In(companyAddress string) *PersonJobBuilder {
	pjb.person.workAddress = companyAddress
	return pjb
}

func (pjb *PersonJobBuilder) WithSalary(salary int) *PersonJobBuilder {
	pjb.person.salary = salary
	return pjb
}

func (pb *PersonBuilder) Build() *Person {
	return pb.person
}

func main() {
	p := fmt.Println
	p(PatternName)

	pb := NewPersonBuilder()
	pb.Lives().
		At("Bangalore").
		WithPostalCode("560102").
		Works().
		As("Software Engineer").
		For("IBM").
		In("Bangalore").
		WithSalary(150000)
	person := pb.Build()

	p(person)

}
