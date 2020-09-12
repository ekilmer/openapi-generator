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
	"time"
)

// FormatTest struct for FormatTest
type FormatTest struct {
	Integer *int32 `json:"integer,omitempty"`
	Int32 *int32 `json:"int32,omitempty"`
	Int64 *int64 `json:"int64,omitempty"`
	Number float32 `json:"number"`
	Float *float32 `json:"float,omitempty"`
	Double *float64 `json:"double,omitempty"`
	String *string `json:"string,omitempty"`
	Byte string `json:"byte"`
	Binary **os.File `json:"binary,omitempty"`
	Date string `json:"date"`
	DateTime *time.Time `json:"dateTime,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Password string `json:"password"`
	// A string that is a 10 digit number. Can have leading zeros.
	PatternWithDigits *string `json:"pattern_with_digits,omitempty"`
	// A string starting with 'image_' (case insensitive) and one to three digits following i.e. Image_01.
	PatternWithDigitsAndDelimiter *string `json:"pattern_with_digits_and_delimiter,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FormatTest FormatTest

// NewFormatTest instantiates a new FormatTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormatTest(Number float32, Byte string, Date string, Password string, ) *FormatTest {
	this := FormatTest{}
	this.Number = Number
	this.Byte = Byte
	this.Date = Date
	this.Password = Password
	return &this
}

// NewFormatTestWithDefaults instantiates a new FormatTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormatTestWithDefaults() *FormatTest {
	this := FormatTest{}
	return &this
}

// GetInteger returns the Integer field value if set, zero value otherwise.
func (o *FormatTest) GetInteger()  {
	if o == nil || o.Integer == nil {
		var ret 
		return ret
	}
	return *o.Integer
}

// GetIntegerOk returns a tuple with the Integer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetIntegerOk() (*, bool) {
	if o == nil || o.Integer == nil {
		return nil, false
	}
	return o.Integer, true
}

// HasInteger returns a boolean if a field has been set.
func (o *FormatTest) HasInteger() bool {
	if o != nil && o.Integer != nil {
		return true
	}

	return false
}

// SetInteger gets a reference to the given int32 and assigns it to the Integer field.
func (o *FormatTest) SetInteger(v ) {
	o.Integer = &v
}

// GetInt32 returns the Int32 field value if set, zero value otherwise.
func (o *FormatTest) GetInt32()  {
	if o == nil || o.Int32 == nil {
		var ret 
		return ret
	}
	return *o.Int32
}

// GetInt32Ok returns a tuple with the Int32 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetInt32Ok() (*, bool) {
	if o == nil || o.Int32 == nil {
		return nil, false
	}
	return o.Int32, true
}

// HasInt32 returns a boolean if a field has been set.
func (o *FormatTest) HasInt32() bool {
	if o != nil && o.Int32 != nil {
		return true
	}

	return false
}

// SetInt32 gets a reference to the given int32 and assigns it to the Int32 field.
func (o *FormatTest) SetInt32(v ) {
	o.Int32 = &v
}

// GetInt64 returns the Int64 field value if set, zero value otherwise.
func (o *FormatTest) GetInt64()  {
	if o == nil || o.Int64 == nil {
		var ret 
		return ret
	}
	return *o.Int64
}

// GetInt64Ok returns a tuple with the Int64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetInt64Ok() (*, bool) {
	if o == nil || o.Int64 == nil {
		return nil, false
	}
	return o.Int64, true
}

// HasInt64 returns a boolean if a field has been set.
func (o *FormatTest) HasInt64() bool {
	if o != nil && o.Int64 != nil {
		return true
	}

	return false
}

// SetInt64 gets a reference to the given int64 and assigns it to the Int64 field.
func (o *FormatTest) SetInt64(v ) {
	o.Int64 = &v
}

// GetNumber returns the Number field value
func (o *FormatTest) GetNumber()  {
	if o == nil  {
		var ret 
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *FormatTest) GetNumberOk() (*, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *FormatTest) SetNumber(v ) {
	o.Number = v
}

// GetFloat returns the Float field value if set, zero value otherwise.
func (o *FormatTest) GetFloat()  {
	if o == nil || o.Float == nil {
		var ret 
		return ret
	}
	return *o.Float
}

// GetFloatOk returns a tuple with the Float field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetFloatOk() (*, bool) {
	if o == nil || o.Float == nil {
		return nil, false
	}
	return o.Float, true
}

// HasFloat returns a boolean if a field has been set.
func (o *FormatTest) HasFloat() bool {
	if o != nil && o.Float != nil {
		return true
	}

	return false
}

// SetFloat gets a reference to the given float32 and assigns it to the Float field.
func (o *FormatTest) SetFloat(v ) {
	o.Float = &v
}

// GetDouble returns the Double field value if set, zero value otherwise.
func (o *FormatTest) GetDouble()  {
	if o == nil || o.Double == nil {
		var ret 
		return ret
	}
	return *o.Double
}

// GetDoubleOk returns a tuple with the Double field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetDoubleOk() (*, bool) {
	if o == nil || o.Double == nil {
		return nil, false
	}
	return o.Double, true
}

// HasDouble returns a boolean if a field has been set.
func (o *FormatTest) HasDouble() bool {
	if o != nil && o.Double != nil {
		return true
	}

	return false
}

// SetDouble gets a reference to the given float64 and assigns it to the Double field.
func (o *FormatTest) SetDouble(v ) {
	o.Double = &v
}

// GetString returns the String field value if set, zero value otherwise.
func (o *FormatTest) GetString()  {
	if o == nil || o.String == nil {
		var ret 
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetStringOk() (*, bool) {
	if o == nil || o.String == nil {
		return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *FormatTest) HasString() bool {
	if o != nil && o.String != nil {
		return true
	}

	return false
}

// SetString gets a reference to the given string and assigns it to the String field.
func (o *FormatTest) SetString(v ) {
	o.String = &v
}

// GetByte returns the Byte field value
func (o *FormatTest) GetByte()  {
	if o == nil  {
		var ret 
		return ret
	}

	return o.Byte
}

// GetByteOk returns a tuple with the Byte field value
// and a boolean to check if the value has been set.
func (o *FormatTest) GetByteOk() (*, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Byte, true
}

// SetByte sets field value
func (o *FormatTest) SetByte(v ) {
	o.Byte = v
}

// GetBinary returns the Binary field value if set, zero value otherwise.
func (o *FormatTest) GetBinary()  {
	if o == nil || o.Binary == nil {
		var ret 
		return ret
	}
	return *o.Binary
}

// GetBinaryOk returns a tuple with the Binary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetBinaryOk() (*, bool) {
	if o == nil || o.Binary == nil {
		return nil, false
	}
	return o.Binary, true
}

// HasBinary returns a boolean if a field has been set.
func (o *FormatTest) HasBinary() bool {
	if o != nil && o.Binary != nil {
		return true
	}

	return false
}

// SetBinary gets a reference to the given *os.File and assigns it to the Binary field.
func (o *FormatTest) SetBinary(v ) {
	o.Binary = &v
}

// GetDate returns the Date field value
func (o *FormatTest) GetDate()  {
	if o == nil  {
		var ret 
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *FormatTest) GetDateOk() (*, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *FormatTest) SetDate(v ) {
	o.Date = v
}

// GetDateTime returns the DateTime field value if set, zero value otherwise.
func (o *FormatTest) GetDateTime()  {
	if o == nil || o.DateTime == nil {
		var ret 
		return ret
	}
	return *o.DateTime
}

// GetDateTimeOk returns a tuple with the DateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetDateTimeOk() (*, bool) {
	if o == nil || o.DateTime == nil {
		return nil, false
	}
	return o.DateTime, true
}

// HasDateTime returns a boolean if a field has been set.
func (o *FormatTest) HasDateTime() bool {
	if o != nil && o.DateTime != nil {
		return true
	}

	return false
}

// SetDateTime gets a reference to the given time.Time and assigns it to the DateTime field.
func (o *FormatTest) SetDateTime(v ) {
	o.DateTime = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *FormatTest) GetUuid()  {
	if o == nil || o.Uuid == nil {
		var ret 
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetUuidOk() (*, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *FormatTest) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *FormatTest) SetUuid(v ) {
	o.Uuid = &v
}

// GetPassword returns the Password field value
func (o *FormatTest) GetPassword()  {
	if o == nil  {
		var ret 
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *FormatTest) GetPasswordOk() (*, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *FormatTest) SetPassword(v ) {
	o.Password = v
}

// GetPatternWithDigits returns the PatternWithDigits field value if set, zero value otherwise.
func (o *FormatTest) GetPatternWithDigits()  {
	if o == nil || o.PatternWithDigits == nil {
		var ret 
		return ret
	}
	return *o.PatternWithDigits
}

// GetPatternWithDigitsOk returns a tuple with the PatternWithDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetPatternWithDigitsOk() (*, bool) {
	if o == nil || o.PatternWithDigits == nil {
		return nil, false
	}
	return o.PatternWithDigits, true
}

// HasPatternWithDigits returns a boolean if a field has been set.
func (o *FormatTest) HasPatternWithDigits() bool {
	if o != nil && o.PatternWithDigits != nil {
		return true
	}

	return false
}

// SetPatternWithDigits gets a reference to the given string and assigns it to the PatternWithDigits field.
func (o *FormatTest) SetPatternWithDigits(v ) {
	o.PatternWithDigits = &v
}

// GetPatternWithDigitsAndDelimiter returns the PatternWithDigitsAndDelimiter field value if set, zero value otherwise.
func (o *FormatTest) GetPatternWithDigitsAndDelimiter()  {
	if o == nil || o.PatternWithDigitsAndDelimiter == nil {
		var ret 
		return ret
	}
	return *o.PatternWithDigitsAndDelimiter
}

// GetPatternWithDigitsAndDelimiterOk returns a tuple with the PatternWithDigitsAndDelimiter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetPatternWithDigitsAndDelimiterOk() (*, bool) {
	if o == nil || o.PatternWithDigitsAndDelimiter == nil {
		return nil, false
	}
	return o.PatternWithDigitsAndDelimiter, true
}

// HasPatternWithDigitsAndDelimiter returns a boolean if a field has been set.
func (o *FormatTest) HasPatternWithDigitsAndDelimiter() bool {
	if o != nil && o.PatternWithDigitsAndDelimiter != nil {
		return true
	}

	return false
}

// SetPatternWithDigitsAndDelimiter gets a reference to the given string and assigns it to the PatternWithDigitsAndDelimiter field.
func (o *FormatTest) SetPatternWithDigitsAndDelimiter(v ) {
	o.PatternWithDigitsAndDelimiter = &v
}

func (o FormatTest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Integer != nil {
		toSerialize["integer"] = o.Integer
	}
	if o.Int32 != nil {
		toSerialize["int32"] = o.Int32
	}
	if o.Int64 != nil {
		toSerialize["int64"] = o.Int64
	}
	if true {
		toSerialize["number"] = o.Number
	}
	if o.Float != nil {
		toSerialize["float"] = o.Float
	}
	if o.Double != nil {
		toSerialize["double"] = o.Double
	}
	if o.String != nil {
		toSerialize["string"] = o.String
	}
	if true {
		toSerialize["byte"] = o.Byte
	}
	if o.Binary != nil {
		toSerialize["binary"] = o.Binary
	}
	if true {
		toSerialize["date"] = o.Date
	}
	if o.DateTime != nil {
		toSerialize["dateTime"] = o.DateTime
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if o.PatternWithDigits != nil {
		toSerialize["pattern_with_digits"] = o.PatternWithDigits
	}
	if o.PatternWithDigitsAndDelimiter != nil {
		toSerialize["pattern_with_digits_and_delimiter"] = o.PatternWithDigitsAndDelimiter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FormatTest) UnmarshalJSON(bytes []byte) (err error) {
	varFormatTest := _FormatTest{}

	if err = json.Unmarshal(bytes, &varFormatTest); err == nil {
		*o = FormatTest(varFormatTest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "integer")
		delete(additionalProperties, "int32")
		delete(additionalProperties, "int64")
		delete(additionalProperties, "number")
		delete(additionalProperties, "float")
		delete(additionalProperties, "double")
		delete(additionalProperties, "string")
		delete(additionalProperties, "byte")
		delete(additionalProperties, "binary")
		delete(additionalProperties, "date")
		delete(additionalProperties, "dateTime")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "password")
		delete(additionalProperties, "pattern_with_digits")
		delete(additionalProperties, "pattern_with_digits_and_delimiter")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFormatTest struct {
	value *FormatTest
	isSet bool
}

func (v NullableFormatTest) Get() *FormatTest {
	return v.value
}

func (v *NullableFormatTest) Set(val *FormatTest) {
	v.value = val
	v.isSet = true
}

func (v NullableFormatTest) IsSet() bool {
	return v.isSet
}

func (v *NullableFormatTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormatTest(val *FormatTest) *NullableFormatTest {
	return &NullableFormatTest{value: val, isSet: true}
}

func (v NullableFormatTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormatTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

