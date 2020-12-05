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

// MapTestMapOfEnumStringAddlProps the model 'MapTestMapOfEnumStringAddlProps'
type MapTestMapOfEnumStringAddlProps string

// List of MapTestMapOfEnumStringAddlProps
const (
	MAP_TEST_MAP_OF_ENUM_STRING_ADDL_PROPS_UPPER MapTestMapOfEnumStringAddlProps = "UPPER"
	MAP_TEST_MAP_OF_ENUM_STRING_ADDL_PROPS_LOWER MapTestMapOfEnumStringAddlProps = "lower"
)

func (v *MapTestMapOfEnumStringAddlProps) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MapTestMapOfEnumStringAddlProps(value)
	for _, existing := range []MapTestMapOfEnumStringAddlProps{ "UPPER", "lower",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MapTestMapOfEnumStringAddlProps", value)
}

// Ptr returns reference to MapTestMapOfEnumStringAddlProps value
func (v MapTestMapOfEnumStringAddlProps) Ptr() *MapTestMapOfEnumStringAddlProps {
	return &v
}

type NullableMapTestMapOfEnumStringAddlProps struct {
	value *MapTestMapOfEnumStringAddlProps
	isSet bool
}

func (v NullableMapTestMapOfEnumStringAddlProps) Get() *MapTestMapOfEnumStringAddlProps {
	return v.value
}

func (v *NullableMapTestMapOfEnumStringAddlProps) Set(val *MapTestMapOfEnumStringAddlProps) {
	v.value = val
	v.isSet = true
}

func (v NullableMapTestMapOfEnumStringAddlProps) IsSet() bool {
	return v.isSet
}

func (v *NullableMapTestMapOfEnumStringAddlProps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMapTestMapOfEnumStringAddlProps(val *MapTestMapOfEnumStringAddlProps) *NullableMapTestMapOfEnumStringAddlProps {
	return &NullableMapTestMapOfEnumStringAddlProps{value: val, isSet: true}
}

func (v NullableMapTestMapOfEnumStringAddlProps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMapTestMapOfEnumStringAddlProps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

