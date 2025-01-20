package data

import "time"

type Workshop struct {
	Course // Embeddding - Adding types without an identifier
	Date   time.Time
	Name   string
	Instructor
}

// Factory pattern
func NewWorkshop(name string) Workshop {
	w := Workshop{}
	w.Name = name
	w.Date = time.Now()
	w.Course.Name = name
	return w
}
