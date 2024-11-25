/*
 * Product Catalog Management
 *
 * Product Catalog API is one of Catalog Management API Family. Product Catalog API goal is to provide a catalog of products.  ### Operations Product Catalog API performs the following operations on the resources : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity - Manage notification of events
 *
 * API version: 5.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

type StringCharacteristicValueSpecificationMvo struct {
	// Value of the characteristic
	Value string `json:"value,omitempty"`
	// A kind of value that the characteristic value can take on, such as numeric, text and so forth
	ValueType string `json:"valueType,omitempty"`
	// If true, the Boolean Indicates if the value is the default value for a characteristic
	IsDefault bool `json:"isDefault,omitempty"`
	// A length, surface, volume, dry measure, liquid measure, money, weight, time, and the like. In general, a determinate quantity or magnitude of the kind designated, taken as a standard of comparison for others of the same kind, in assigning to them numerical values, as 1 foot, 1 yard, 1 mile, 1 square foot.
	UnitOfMeasure string `json:"unitOfMeasure,omitempty"`

	ValidFor *TimePeriod `json:"validFor,omitempty"`
	// The low range value that a characteristic can take on
	ValueFrom int32 `json:"valueFrom,omitempty"`
	// The upper range value that a characteristic can take on
	ValueTo int32 `json:"valueTo,omitempty"`
	// An indicator that specifies the inclusion or exclusion of the valueFrom and valueTo attributes. If applicable, possible values are \"open\", \"closed\", \"closedBottom\" and \"closedTop\".
	RangeInterval string `json:"rangeInterval,omitempty"`
	// A regular expression constraint for given value
	Regex string `json:"regex,omitempty"`
}