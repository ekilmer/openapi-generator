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

// EnumQueryDouble the model 'EnumQueryDouble'
type EnumQueryDouble float64

// List of enum_query_double
const (
	ENUM_QUERY_DOUBLE__1_DOT_1 EnumQueryDouble = 1.1
	ENUM_QUERY_DOUBLE__MINUS_1_DOT_2 EnumQueryDouble = -1.2
)

func (v *EnumQueryDouble) UnmarshalJSON(src []byte) error {
	var value float64
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumQueryDouble(value)
	for _, existing := range []EnumQueryDouble{ 1.1, -1.2,   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumQueryDouble", value)
}

// Ptr returns reference to enum_query_double value
func (v EnumQueryDouble) Ptr() *EnumQueryDouble {
	return &v
}

type NullableEnumQueryDouble struct {
	value *EnumQueryDouble
	isSet bool
}

func (v NullableEnumQueryDouble) Get() *EnumQueryDouble {
	return v.value
}

func (v *NullableEnumQueryDouble) Set(val *EnumQueryDouble) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumQueryDouble) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumQueryDouble) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumQueryDouble(val *EnumQueryDouble) *NullableEnumQueryDouble {
	return &NullableEnumQueryDouble{value: val, isSet: true}
}

func (v NullableEnumQueryDouble) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumQueryDouble) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

