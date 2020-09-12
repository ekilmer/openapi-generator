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

// ArrayOfNumberOnly struct for ArrayOfNumberOnly
type ArrayOfNumberOnly struct {
	ArrayNumber *[]float32 `json:"ArrayNumber,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ArrayOfNumberOnly ArrayOfNumberOnly

// NewArrayOfNumberOnly instantiates a new ArrayOfNumberOnly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArrayOfNumberOnly() *ArrayOfNumberOnly {
	this := ArrayOfNumberOnly{}
	return &this
}

// NewArrayOfNumberOnlyWithDefaults instantiates a new ArrayOfNumberOnly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArrayOfNumberOnlyWithDefaults() *ArrayOfNumberOnly {
	this := ArrayOfNumberOnly{}
	return &this
}

// GetArrayNumber returns the ArrayNumber field value if set, zero value otherwise.
func (o *ArrayOfNumberOnly) GetArrayNumber()  {
	if o == nil || o.ArrayNumber == nil {
		var ret 
		return ret
	}
	return *o.ArrayNumber
}

// GetArrayNumberOk returns a tuple with the ArrayNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArrayOfNumberOnly) GetArrayNumberOk() (*, bool) {
	if o == nil || o.ArrayNumber == nil {
		return nil, false
	}
	return o.ArrayNumber, true
}

// HasArrayNumber returns a boolean if a field has been set.
func (o *ArrayOfNumberOnly) HasArrayNumber() bool {
	if o != nil && o.ArrayNumber != nil {
		return true
	}

	return false
}

// SetArrayNumber gets a reference to the given []float32 and assigns it to the ArrayNumber field.
func (o *ArrayOfNumberOnly) SetArrayNumber(v ) {
	o.ArrayNumber = &v
}

func (o ArrayOfNumberOnly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArrayNumber != nil {
		toSerialize["ArrayNumber"] = o.ArrayNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ArrayOfNumberOnly) UnmarshalJSON(bytes []byte) (err error) {
	varArrayOfNumberOnly := _ArrayOfNumberOnly{}

	if err = json.Unmarshal(bytes, &varArrayOfNumberOnly); err == nil {
		*o = ArrayOfNumberOnly(varArrayOfNumberOnly)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ArrayNumber")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableArrayOfNumberOnly struct {
	value *ArrayOfNumberOnly
	isSet bool
}

func (v NullableArrayOfNumberOnly) Get() *ArrayOfNumberOnly {
	return v.value
}

func (v *NullableArrayOfNumberOnly) Set(val *ArrayOfNumberOnly) {
	v.value = val
	v.isSet = true
}

func (v NullableArrayOfNumberOnly) IsSet() bool {
	return v.isSet
}

func (v *NullableArrayOfNumberOnly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArrayOfNumberOnly(val *ArrayOfNumberOnly) *NullableArrayOfNumberOnly {
	return &NullableArrayOfNumberOnly{value: val, isSet: true}
}

func (v NullableArrayOfNumberOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArrayOfNumberOnly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


