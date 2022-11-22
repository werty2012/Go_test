package person

import "fmt"

type Person struct {
	FirstName string
	Age       int
}

func (p *Person) SetFirstName(name string) {
	p.FirstName = name

}
func (p Person) GetFirstName() string {
	return p.FirstName
}

func (p Person) GetAge() int {
	return p.Age
}

func (p Person) SetAge(i int) {
	p.Age = i
	fmt.Print(p)
}
