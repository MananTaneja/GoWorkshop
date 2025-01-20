package data

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

// Factory - Constructor alternative
func NewInstructor(name string, lastName string) Instructor {
	// generate a random id here
	// add custom default score
	return Instructor{FirstName: name, LastName: lastName}
}

func (i Instructor) GetData() string {
	return fmt.Sprintf("First name: %v", i.FirstName)
}
