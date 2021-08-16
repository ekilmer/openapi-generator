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

// OuterEnumIntegerDefaultValue the model 'OuterEnumIntegerDefaultValue'
type OuterEnumIntegerDefaultValue int32

// List of OuterEnumIntegerDefaultValue
const (
	OUTER_ENUM_INTEGER_DEFAULT_VALUE__0 OuterEnumIntegerDefaultValue = 0
	OUTER_ENUM_INTEGER_DEFAULT_VALUE__1 OuterEnumIntegerDefaultValue = 1
	OUTER_ENUM_INTEGER_DEFAULT_VALUE__2 OuterEnumIntegerDefaultValue = 2
)

var allowedOuterEnumIntegerDefaultValueEnumValues = []OuterEnumIntegerDefaultValue{
	0,
	1,
	2,
}

func (v *OuterEnumIntegerDefaultValue) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OuterEnumIntegerDefaultValue(value)
	for _, existing := range allowedOuterEnumIntegerDefaultValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OuterEnumIntegerDefaultValue", value)
}

// NewOuterEnumIntegerDefaultValueFromValue returns a pointer to a valid OuterEnumIntegerDefaultValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOuterEnumIntegerDefaultValueFromValue(v int32) (*OuterEnumIntegerDefaultValue, error) {
	ev := OuterEnumIntegerDefaultValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OuterEnumIntegerDefaultValue: valid values are %v", v, allowedOuterEnumIntegerDefaultValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OuterEnumIntegerDefaultValue) IsValid() bool {
	for _, existing := range allowedOuterEnumIntegerDefaultValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OuterEnumIntegerDefaultValue value
func (v OuterEnumIntegerDefaultValue) Ptr() *OuterEnumIntegerDefaultValue {
	return &v
}

type NullableOuterEnumIntegerDefaultValue struct {
	value *OuterEnumIntegerDefaultValue
	isSet bool
}

func (v NullableOuterEnumIntegerDefaultValue) Get() *OuterEnumIntegerDefaultValue {
	return v.value
}

func (v *NullableOuterEnumIntegerDefaultValue) Set(val *OuterEnumIntegerDefaultValue) {
	v.value = val
	v.isSet = true
}

func (v NullableOuterEnumIntegerDefaultValue) IsSet() bool {
	return v.isSet
}

func (v *NullableOuterEnumIntegerDefaultValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOuterEnumIntegerDefaultValue(val *OuterEnumIntegerDefaultValue) *NullableOuterEnumIntegerDefaultValue {
	return &NullableOuterEnumIntegerDefaultValue{value: val, isSet: true}
}

func (v NullableOuterEnumIntegerDefaultValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOuterEnumIntegerDefaultValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

