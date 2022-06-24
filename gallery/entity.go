package gallery

import "time"

type Gallery struct {
	ID        int
	ProductID int
	Image     string
	IsDefault bool
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
