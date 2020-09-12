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

// TestJsonFormDataBody struct for TestJsonFormDataBody
type TestJsonFormDataBody struct {
	// field1
	Param string `json:"param"`
	// field2
	Param2 string `json:"param2"`
	AdditionalProperties map[string]interface{}
}

type _TestJsonFormDataBody TestJsonFormDataBody

// NewTestJsonFormDataBody instantiates a new TestJsonFormDataBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestJsonFormDataBody(Param string, Param2 string, ) *TestJsonFormDataBody {
	this := TestJsonFormDataBody{}
	this.Param = Param
	this.Param2 = Param2
	return &this
}

// NewTestJsonFormDataBodyWithDefaults instantiates a new TestJsonFormDataBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestJsonFormDataBodyWithDefaults() *TestJsonFormDataBody {
	this := TestJsonFormDataBody{}
	return &this
}

// GetParam returns the Param field value
func (o *TestJsonFormDataBody) GetParam()  {
	if o == nil  {
		var ret 
		return ret
	}

	return o.Param
}

// GetParamOk returns a tuple with the Param field value
// and a boolean to check if the value has been set.
func (o *TestJsonFormDataBody) GetParamOk() (*, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Param, true
}

// SetParam sets field value
func (o *TestJsonFormDataBody) SetParam(v ) {
	o.Param = v
}

// GetParam2 returns the Param2 field value
func (o *TestJsonFormDataBody) GetParam2()  {
	if o == nil  {
		var ret 
		return ret
	}

	return o.Param2
}

// GetParam2Ok returns a tuple with the Param2 field value
// and a boolean to check if the value has been set.
func (o *TestJsonFormDataBody) GetParam2Ok() (*, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Param2, true
}

// SetParam2 sets field value
func (o *TestJsonFormDataBody) SetParam2(v ) {
	o.Param2 = v
}

func (o TestJsonFormDataBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["param"] = o.Param
	}
	if true {
		toSerialize["param2"] = o.Param2
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TestJsonFormDataBody) UnmarshalJSON(bytes []byte) (err error) {
	varTestJsonFormDataBody := _TestJsonFormDataBody{}

	if err = json.Unmarshal(bytes, &varTestJsonFormDataBody); err == nil {
		*o = TestJsonFormDataBody(varTestJsonFormDataBody)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "param")
		delete(additionalProperties, "param2")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestJsonFormDataBody struct {
	value *TestJsonFormDataBody
	isSet bool
}

func (v NullableTestJsonFormDataBody) Get() *TestJsonFormDataBody {
	return v.value
}

func (v *NullableTestJsonFormDataBody) Set(val *TestJsonFormDataBody) {
	v.value = val
	v.isSet = true
}

func (v NullableTestJsonFormDataBody) IsSet() bool {
	return v.isSet
}

func (v *NullableTestJsonFormDataBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestJsonFormDataBody(val *TestJsonFormDataBody) *NullableTestJsonFormDataBody {
	return &NullableTestJsonFormDataBody{value: val, isSet: true}
}

func (v NullableTestJsonFormDataBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestJsonFormDataBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


