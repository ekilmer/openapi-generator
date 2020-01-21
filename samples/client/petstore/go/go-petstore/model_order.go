/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore
import (
	"time"
)
// Order struct for Order
type Order struct {
	Id int64 `json:"id,omitempty"`
	PetId int64 `json:"petId,omitempty"`
	Quantity int32 `json:"quantity,omitempty"`
	ShipDate time.Time `json:"shipDate,omitempty"`
	Status OrderStatus `json:"status,omitempty"`
	Complete bool `json:"complete,omitempty"`
}
