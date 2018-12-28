package syracuse

import "time"

// Citizen ...
type Citizen struct {
	ID       string `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	FullName string `json:"full_name" db:"full_name"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"-" db:"deleted_at"`
}

// Citizens ...
type Citizens interface {
	Get(*CitizensQuery) (*Citizen, error)
	Select() ([]*Citizen, error)
	Create(*Citizen) error
	Update(*Citizen) error
	Delete(*Citizen) error
}

// CitizensStore ...
type CitizensStore interface {
	Get(*CitizensQuery) (*Citizen, error)
	Select() ([]*Citizen, error)
	Create(*Citizen) error
	Update(*Citizen) error
	Delete(*Citizen) error
}

// CitizensQuery ...
type CitizensQuery struct {
	ID       string
	Email    string
	FullName string
}
