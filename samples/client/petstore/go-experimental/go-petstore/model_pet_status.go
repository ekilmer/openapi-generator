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
	"encoding/json"
	"fmt"
)

// PetStatus pet status in the store
type PetStatus string

// List of PetStatus
const (
	PET_STATUS_AVAILABLE PetStatus = "available"
	PET_STATUS_PENDING PetStatus = "pending"
	PET_STATUS_SOLD PetStatus = "sold"
)

func (v *PetStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PetStatus(value)
	for _, existing := range []PetStatus{ "available", "pending", "sold",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PetStatus", value)
}

// Ptr returns reference to PetStatus value
func (v PetStatus) Ptr() *PetStatus {
	return &v
}

type NullablePetStatus struct {
	value *PetStatus
	isSet bool
}

func (v NullablePetStatus) Get() *PetStatus {
	return v.value
}

func (v *NullablePetStatus) Set(val *PetStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePetStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePetStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePetStatus(val *PetStatus) *NullablePetStatus {
	return &NullablePetStatus{value: val, isSet: true}
}

func (v NullablePetStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePetStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
