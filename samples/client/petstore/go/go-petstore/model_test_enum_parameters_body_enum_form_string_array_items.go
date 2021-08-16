/*
OpenAPI Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
	"fmt"
)

// TestEnumParametersBodyEnumFormStringArrayItems the model 'TestEnumParametersBodyEnumFormStringArrayItems'
type TestEnumParametersBodyEnumFormStringArrayItems string

// List of testEnumParametersBodyEnumFormStringArrayItems
const (
	GREATER_THAN TestEnumParametersBodyEnumFormStringArrayItems = ">"
	DOLLAR TestEnumParametersBodyEnumFormStringArrayItems = "$"
)

var allowedTestEnumParametersBodyEnumFormStringArrayItemsEnumValues = []TestEnumParametersBodyEnumFormStringArrayItems{
	">",
	"$",
}

func (v *TestEnumParametersBodyEnumFormStringArrayItems) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TestEnumParametersBodyEnumFormStringArrayItems(value)
	for _, existing := range allowedTestEnumParametersBodyEnumFormStringArrayItemsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TestEnumParametersBodyEnumFormStringArrayItems", value)
}

// NewTestEnumParametersBodyEnumFormStringArrayItemsFromValue returns a pointer to a valid TestEnumParametersBodyEnumFormStringArrayItems
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTestEnumParametersBodyEnumFormStringArrayItemsFromValue(v string) (*TestEnumParametersBodyEnumFormStringArrayItems, error) {
	ev := TestEnumParametersBodyEnumFormStringArrayItems(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TestEnumParametersBodyEnumFormStringArrayItems: valid values are %v", v, allowedTestEnumParametersBodyEnumFormStringArrayItemsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TestEnumParametersBodyEnumFormStringArrayItems) IsValid() bool {
	for _, existing := range allowedTestEnumParametersBodyEnumFormStringArrayItemsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to testEnumParametersBodyEnumFormStringArrayItems value
func (v TestEnumParametersBodyEnumFormStringArrayItems) Ptr() *TestEnumParametersBodyEnumFormStringArrayItems {
	return &v
}

type NullableTestEnumParametersBodyEnumFormStringArrayItems struct {
	value *TestEnumParametersBodyEnumFormStringArrayItems
	isSet bool
}

func (v NullableTestEnumParametersBodyEnumFormStringArrayItems) Get() *TestEnumParametersBodyEnumFormStringArrayItems {
	return v.value
}

func (v *NullableTestEnumParametersBodyEnumFormStringArrayItems) Set(val *TestEnumParametersBodyEnumFormStringArrayItems) {
	v.value = val
	v.isSet = true
}

func (v NullableTestEnumParametersBodyEnumFormStringArrayItems) IsSet() bool {
	return v.isSet
}

func (v *NullableTestEnumParametersBodyEnumFormStringArrayItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestEnumParametersBodyEnumFormStringArrayItems(val *TestEnumParametersBodyEnumFormStringArrayItems) *NullableTestEnumParametersBodyEnumFormStringArrayItems {
	return &NullableTestEnumParametersBodyEnumFormStringArrayItems{value: val, isSet: true}
}

func (v NullableTestEnumParametersBodyEnumFormStringArrayItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestEnumParametersBodyEnumFormStringArrayItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

