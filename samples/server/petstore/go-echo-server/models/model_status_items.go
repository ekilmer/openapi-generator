package models

type StatusItems string

// List of StatusItems
const (
	AVAILABLE StatusItems = "available"
	PENDING StatusItems = "pending"
	SOLD StatusItems = "sold"
)
