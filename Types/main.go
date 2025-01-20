package main

import (
	"fmt"

	"manantaneja.com/go/types/data"
)

func main() {
	inst := data.Instructor{Id: 1, FirstName: "Some"}
	// inst2 := data.NewInstructor("second", "last")

	course1 := data.Course{Id: 123, Name: "Basics of Go", Author: inst}
	fmt.Printf("%s\n", inst.GetData())
	fmt.Printf("%s\n", course1.String())

	workshop1 := data.NewWorkshop("Workshop")

}
