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
)

// TestEnumParametersBodyEnumFormStringArrayItems the model 'TestEnumParametersBodyEnumFormStringArrayItems'
type TestEnumParametersBodyEnumFormStringArrayItems string

// List of testEnumParametersBodyEnumFormStringArrayItems
const (
	TEST_ENUM_PARAMETERS_BODY_ENUM_FORM_STRING_ARRAY_ITEMS_GREATER_THAN TestEnumParametersBodyEnumFormStringArrayItems = ">"
	TEST_ENUM_PARAMETERS_BODY_ENUM_FORM_STRING_ARRAY_ITEMS_DOLLAR TestEnumParametersBodyEnumFormStringArrayItems = "$"
)

func (v *TestEnumParametersBodyEnumFormStringArrayItems) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TestEnumParametersBodyEnumFormStringArrayItems(value)
	for _, existing := range []TestEnumParametersBodyEnumFormStringArrayItems{ ">", "$",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TestEnumParametersBodyEnumFormStringArrayItems", value)
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

