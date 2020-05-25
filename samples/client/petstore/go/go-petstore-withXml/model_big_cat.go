/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore
// BigCat struct for BigCat
type BigCat struct {
	ClassName string `json:"className" xml:"className"`
	Color string `json:"color,omitempty" xml:"color"`
	Declawed bool `json:"declawed,omitempty" xml:"declawed"`
	Kind BigCatAllOfKind `json:"kind,omitempty" xml:"kind"`
}
