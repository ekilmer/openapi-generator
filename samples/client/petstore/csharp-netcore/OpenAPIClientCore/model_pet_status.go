/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore
// PetStatus pet status in the store
type PetStatus string

// List of PetStatus
const (
	PET_STATUS_AVAILABLE PetStatus = "available"
	PET_STATUS_PENDING PetStatus = "pending"
	PET_STATUS_SOLD PetStatus = "sold"
)