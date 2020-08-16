/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore
// PetStatus pet status in the store
type PetStatus string

// List of PetStatus
const (
	PETSTATUS_PET_STATUS_AVAILABLE PetStatus = "available"
	PETSTATUS_PET_STATUS_PENDING PetStatus = "pending"
	PETSTATUS_PET_STATUS_SOLD PetStatus = "sold"
)
