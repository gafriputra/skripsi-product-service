package document

import "time"

type Document struct {
	ID           int
	ProductID    int
	Name         string
	DocumentLink string
	Type         string
	Status       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
