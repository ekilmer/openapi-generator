/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
	"fmt"
)

// EnumArraysArrayEnumItems the model 'EnumArraysArrayEnumItems'
type EnumArraysArrayEnumItems string

// List of EnumArraysArrayEnumItems
const (
	ENUM_ARRAYS_ARRAY_ENUM_ITEMS_FISH EnumArraysArrayEnumItems = "fish"
	ENUM_ARRAYS_ARRAY_ENUM_ITEMS_CRAB EnumArraysArrayEnumItems = "crab"
)

func (v *EnumArraysArrayEnumItems) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumArraysArrayEnumItems(value)
	for _, existing := range []EnumArraysArrayEnumItems{ "fish", "crab",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumArraysArrayEnumItems", value)
}

// Ptr returns reference to EnumArraysArrayEnumItems value
func (v EnumArraysArrayEnumItems) Ptr() *EnumArraysArrayEnumItems {
	return &v
}

type NullableEnumArraysArrayEnumItems struct {
	value *EnumArraysArrayEnumItems
	isSet bool
}

func (v NullableEnumArraysArrayEnumItems) Get() *EnumArraysArrayEnumItems {
	return v.value
}

func (v *NullableEnumArraysArrayEnumItems) Set(val *EnumArraysArrayEnumItems) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumArraysArrayEnumItems) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumArraysArrayEnumItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumArraysArrayEnumItems(val *EnumArraysArrayEnumItems) *NullableEnumArraysArrayEnumItems {
	return &NullableEnumArraysArrayEnumItems{value: val, isSet: true}
}

func (v NullableEnumArraysArrayEnumItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumArraysArrayEnumItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

