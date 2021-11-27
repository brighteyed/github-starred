package repository

import "fmt"

// Starred represents starred repository
type Starred struct {
	Url         string
	Description string
}

// NewStarred returns new starred repository
func NewStarred(url, description string) *Starred {
	return &Starred{
		Url:         url,
		Description: description,
	}
}

// String returns string representation of the object so it
// implements fmt.Stringer interface
func (s *Starred) String() string {
	return fmt.Sprintf("%s – %s", s.Url, s.Description)
}
