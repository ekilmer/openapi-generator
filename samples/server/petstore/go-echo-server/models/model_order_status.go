package models
// OrderStatus : Order Status
type OrderStatus string

// List of OrderStatus
const (
	PLACED OrderStatus = "placed"
	APPROVED OrderStatus = "approved"
	DELIVERED OrderStatus = "delivered"
)
