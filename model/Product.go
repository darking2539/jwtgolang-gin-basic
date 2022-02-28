package model

import "time"

// Product 
type Product struct {
	ID        uint      `gorm:"primary_key"`
	Name      string    `json:"name"`
	Stock     int64     `json:"stock"`
	Price     float64   `json:"price"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"time"`
}