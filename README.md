# go_patterns
go patterns designs

## Description
This repo contains simple functional builder patterns design implemented by using go language

main structures

-- Person
-- PersonBuilder: has a reference pointer to Person
-- NewPersonBuilder: return a pointer to PersonBuilder
-- Build: a PersonBuilder function that return the final person instance.
-- Parameters: Setting via PersonBuilder functions that returns a pointer to PersonBuilder  


## Final Example

```go
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
```