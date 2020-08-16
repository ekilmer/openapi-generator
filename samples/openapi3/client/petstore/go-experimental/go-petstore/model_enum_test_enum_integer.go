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

// EnumTestEnumInteger the model 'EnumTestEnumInteger'
type EnumTestEnumInteger int32

// List of Enum_TestEnumInteger
const (
	ENUMTESTENUMINTEGER_ENUM_TEST_ENUM_INTEGER__1 EnumTestEnumInteger = 1
	ENUMTESTENUMINTEGER_ENUM_TEST_ENUM_INTEGER__MINUS_1 EnumTestEnumInteger = -1
)

func (v *EnumTestEnumInteger) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumTestEnumInteger(value)
	for _, existing := range []EnumTestEnumInteger{ 1, -1,   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumTestEnumInteger", value)
}

// Ptr returns reference to Enum_TestEnumInteger value
func (v EnumTestEnumInteger) Ptr() *EnumTestEnumInteger {
	return &v
}

type NullableEnumTestEnumInteger struct {
	value *EnumTestEnumInteger
	isSet bool
}

func (v NullableEnumTestEnumInteger) Get() *EnumTestEnumInteger {
	return v.value
}

func (v *NullableEnumTestEnumInteger) Set(val *EnumTestEnumInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumTestEnumInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumTestEnumInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumTestEnumInteger(val *EnumTestEnumInteger) *NullableEnumTestEnumInteger {
	return &NullableEnumTestEnumInteger{value: val, isSet: true}
}

func (v NullableEnumTestEnumInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumTestEnumInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

