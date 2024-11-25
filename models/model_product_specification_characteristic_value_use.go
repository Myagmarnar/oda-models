/*
 * Product Catalog Management
 *
 * Product Catalog API is one of Catalog Management API Family. Product Catalog API goal is to provide a catalog of products.  ### Operations Product Catalog API performs the following operations on the resources : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity - Manage notification of events
 *
 * API version: 5.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

type ProductSpecificationCharacteristicValueUse struct {
	// Name of the associated productSpecificationCharacteristic
	Name string `json:"name,omitempty"`
	// Unique ID for the characteristic
	Id string `json:"id,omitempty"`
	// A narrative that explains in detail what the productSpecificationCharacteristic is
	Description string `json:"description,omitempty"`
	// A kind of value that the characteristic can take on, such as numeric, text and so forth
	ValueType string `json:"valueType,omitempty"`
	// The minimum number of instances a CharacteristicValue can take on. For example, zero to five phone numbers in a group calling plan, where zero is the value for the minCardinality.
	MinCardinality int32 `json:"minCardinality,omitempty"`
	// The maximum number of instances a CharacteristicValue can take on. For example, zero to five phone numbers in a group calling plan, where five is the value for the maxCardinality.
	MaxCardinality int32 `json:"maxCardinality,omitempty"`

	ValidFor *TimePeriod `json:"validFor,omitempty"`
	// A number or text that can be assigned to a ProductSpecificationCharacteristic.
	ProductSpecCharacteristicValue []CharacteristicValueSpecification `json:"productSpecCharacteristicValue,omitempty"`

	ProductSpecification *ProductSpecificationRef `json:"productSpecification,omitempty"`
	// When sub-classing, this defines the sub-class Extensible name
	Type_ string `json:"@type"`
	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`
	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty"`
}
