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
	"os"
)

// UploadFileWithRequiredFileBody struct for UploadFileWithRequiredFileBody
type UploadFileWithRequiredFileBody struct {
	// Additional data to pass to server
	AdditionalMetadata *string `json:"additionalMetadata,omitempty"`
	// file to upload
	RequiredFile *os.File `json:"requiredFile"`
	AdditionalProperties map[string]interface{}
}

type _UploadFileWithRequiredFileBody UploadFileWithRequiredFileBody

// NewUploadFileWithRequiredFileBody instantiates a new UploadFileWithRequiredFileBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadFileWithRequiredFileBody(requiredFile *os.File, ) *UploadFileWithRequiredFileBody {
	this := UploadFileWithRequiredFileBody{}
	this.RequiredFile = requiredFile
	return &this
}

// NewUploadFileWithRequiredFileBodyWithDefaults instantiates a new UploadFileWithRequiredFileBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadFileWithRequiredFileBodyWithDefaults() *UploadFileWithRequiredFileBody {
	this := UploadFileWithRequiredFileBody{}
	return &this
}

// GetAdditionalMetadata returns the AdditionalMetadata field value if set, zero value otherwise.
func (o *UploadFileWithRequiredFileBody) GetAdditionalMetadata() string {
	if o == nil || o.AdditionalMetadata == nil {
		var ret string
		return ret
	}
	return *o.AdditionalMetadata
}

// GetAdditionalMetadataOk returns a tuple with the AdditionalMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadFileWithRequiredFileBody) GetAdditionalMetadataOk() (*string, bool) {
	if o == nil || o.AdditionalMetadata == nil {
		return nil, false
	}
	return o.AdditionalMetadata, true
}

// HasAdditionalMetadata returns a boolean if a field has been set.
func (o *UploadFileWithRequiredFileBody) HasAdditionalMetadata() bool {
	if o != nil && o.AdditionalMetadata != nil {
		return true
	}

	return false
}

// SetAdditionalMetadata gets a reference to the given string and assigns it to the AdditionalMetadata field.
func (o *UploadFileWithRequiredFileBody) SetAdditionalMetadata(v string) {
	o.AdditionalMetadata = &v
}

// GetRequiredFile returns the RequiredFile field value
func (o *UploadFileWithRequiredFileBody) GetRequiredFile() *os.File {
	if o == nil  {
		var ret *os.File
		return ret
	}

	return o.RequiredFile
}

// GetRequiredFileOk returns a tuple with the RequiredFile field value
// and a boolean to check if the value has been set.
func (o *UploadFileWithRequiredFileBody) GetRequiredFileOk() (**os.File, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequiredFile, true
}

// SetRequiredFile sets field value
func (o *UploadFileWithRequiredFileBody) SetRequiredFile(v *os.File) {
	o.RequiredFile = v
}

func (o UploadFileWithRequiredFileBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalMetadata != nil {
		toSerialize["additionalMetadata"] = o.AdditionalMetadata
	}
	if true {
		toSerialize["requiredFile"] = o.RequiredFile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UploadFileWithRequiredFileBody) UnmarshalJSON(bytes []byte) (err error) {
	varUploadFileWithRequiredFileBody := _UploadFileWithRequiredFileBody{}

	if err = json.Unmarshal(bytes, &varUploadFileWithRequiredFileBody); err == nil {
		*o = UploadFileWithRequiredFileBody(varUploadFileWithRequiredFileBody)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "additionalMetadata")
		delete(additionalProperties, "requiredFile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUploadFileWithRequiredFileBody struct {
	value *UploadFileWithRequiredFileBody
	isSet bool
}

func (v NullableUploadFileWithRequiredFileBody) Get() *UploadFileWithRequiredFileBody {
	return v.value
}

func (v *NullableUploadFileWithRequiredFileBody) Set(val *UploadFileWithRequiredFileBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadFileWithRequiredFileBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadFileWithRequiredFileBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadFileWithRequiredFileBody(val *UploadFileWithRequiredFileBody) *NullableUploadFileWithRequiredFileBody {
	return &NullableUploadFileWithRequiredFileBody{value: val, isSet: true}
}

func (v NullableUploadFileWithRequiredFileBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadFileWithRequiredFileBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


