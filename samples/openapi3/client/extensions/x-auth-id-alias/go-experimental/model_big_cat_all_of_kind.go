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

// BigCatAllOfKind the model 'BigCatAllOfKind'
type BigCatAllOfKind string

// List of BigCatAllOfKind
const (
	BIG_CAT_ALL_OF_KIND_LIONS BigCatAllOfKind = "lions"
	BIG_CAT_ALL_OF_KIND_TIGERS BigCatAllOfKind = "tigers"
	BIG_CAT_ALL_OF_KIND_LEOPARDS BigCatAllOfKind = "leopards"
	BIG_CAT_ALL_OF_KIND_JAGUARS BigCatAllOfKind = "jaguars"
)

func (v *BigCatAllOfKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BigCatAllOfKind(value)
	for _, existing := range []BigCatAllOfKind{ "lions", "tigers", "leopards", "jaguars",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BigCatAllOfKind", value)
}

// Ptr returns reference to BigCatAllOfKind value
func (v BigCatAllOfKind) Ptr() *BigCatAllOfKind {
	return &v
}

type NullableBigCatAllOfKind struct {
	value *BigCatAllOfKind
	isSet bool
}

func (v NullableBigCatAllOfKind) Get() *BigCatAllOfKind {
	return v.value
}

func (v *NullableBigCatAllOfKind) Set(val *BigCatAllOfKind) {
	v.value = val
	v.isSet = true
}

func (v NullableBigCatAllOfKind) IsSet() bool {
	return v.isSet
}

func (v *NullableBigCatAllOfKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBigCatAllOfKind(val *BigCatAllOfKind) *NullableBigCatAllOfKind {
	return &NullableBigCatAllOfKind{value: val, isSet: true}
}

func (v NullableBigCatAllOfKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBigCatAllOfKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
