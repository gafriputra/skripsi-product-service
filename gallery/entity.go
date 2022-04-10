package gallery

import "time"

type image struct {
	ID        int
	ProductID int
	image     string
	IsDefault bool
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
