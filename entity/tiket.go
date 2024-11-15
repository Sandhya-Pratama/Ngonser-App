package entity

import (
	"time"
)

// Ticket structure
type Ticket struct {
	ID          int64     `json:"id"`
	Image       string    `json:"image"`
	Location    string    `json:"location"`
	Date        string    // Format: YYYY-MM-DD
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	Status      string    `json:"Status"` // e.g., 'available', 'sold'
	Quota       int64     `json:"Quota"`
	Category    string    `json:"category"`
	Terjual     int64     `json:"Terjual"` // e.g., 1000, 5000, 10000
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	DeletedAt   time.Time `json:"-"`
}

// NewTicket function
func NewTicket(image, location, date, title, description, status, category string, price, quota, terjual int64) *Ticket {
	return &Ticket{
		Image:       image,
		Location:    location,
		Date:        date,
		Title:       title,
		Description: description,
		Price:       price,
		Status:      status,
		Quota:       quota,
		Category:    category,
		Terjual:     terjual,
	}
}

func UpdateTicket(id int64, image, location, date, title, description, status, category string, price, quota, terjual int64) *Ticket {
	return &Ticket{
		ID:          id,
		Image:       image,
		Location:    location,
		Date:        date,
		Title:       title,
		Description: description,
		Price:       price,
		Status:      status,
		Quota:       quota,
		Category:    category,
		Terjual:     terjual,
	}
}
