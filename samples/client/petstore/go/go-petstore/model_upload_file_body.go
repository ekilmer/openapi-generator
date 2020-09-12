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
	"os"
)

// UploadFileBody struct for UploadFileBody
type UploadFileBody struct {
	// Additional data to pass to server
	AdditionalMetadata *string `json:"additionalMetadata,omitempty"`
	// file to upload
	File **os.File `json:"file,omitempty"`
}

// NewUploadFileBody instantiates a new UploadFileBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadFileBody() *UploadFileBody {
	this := UploadFileBody{}
	return &this
}

// NewUploadFileBodyWithDefaults instantiates a new UploadFileBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadFileBodyWithDefaults() *UploadFileBody {
	this := UploadFileBody{}
	return &this
}

// GetAdditionalMetadata returns the AdditionalMetadata field value if set, zero value otherwise.
func (o *UploadFileBody) GetAdditionalMetadata()  {
	if o == nil || o.AdditionalMetadata == nil {
		var ret 
		return ret
	}
	return *o.AdditionalMetadata
}

// GetAdditionalMetadataOk returns a tuple with the AdditionalMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadFileBody) GetAdditionalMetadataOk() (*, bool) {
	if o == nil || o.AdditionalMetadata == nil {
		return nil, false
	}
	return o.AdditionalMetadata, true
}

// HasAdditionalMetadata returns a boolean if a field has been set.
func (o *UploadFileBody) HasAdditionalMetadata() bool {
	if o != nil && o.AdditionalMetadata != nil {
		return true
	}

	return false
}

// SetAdditionalMetadata gets a reference to the given string and assigns it to the AdditionalMetadata field.
func (o *UploadFileBody) SetAdditionalMetadata(v ) {
	o.AdditionalMetadata = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *UploadFileBody) GetFile()  {
	if o == nil || o.File == nil {
		var ret 
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadFileBody) GetFileOk() (*, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *UploadFileBody) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given *os.File and assigns it to the File field.
func (o *UploadFileBody) SetFile(v ) {
	o.File = &v
}

func (o UploadFileBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalMetadata != nil {
		toSerialize["additionalMetadata"] = o.AdditionalMetadata
	}
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	return json.Marshal(toSerialize)
}

type NullableUploadFileBody struct {
	value *UploadFileBody
	isSet bool
}

func (v NullableUploadFileBody) Get() *UploadFileBody {
	return v.value
}

func (v *NullableUploadFileBody) Set(val *UploadFileBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadFileBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadFileBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadFileBody(val *UploadFileBody) *NullableUploadFileBody {
	return &NullableUploadFileBody{value: val, isSet: true}
}

func (v NullableUploadFileBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadFileBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


