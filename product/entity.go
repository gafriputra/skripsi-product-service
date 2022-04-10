package product

import "time"

type Product struct {
	ID          int
	CaregoryID  int
	Name        string
	Slug        string
	Description string
	Price       int
	Status      bool
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
