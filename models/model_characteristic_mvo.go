/*
 * Product Catalog Management
 *
 * Product Catalog API is one of Catalog Management API Family. Product Catalog API goal is to provide a catalog of products.  ### Operations Product Catalog API performs the following operations on the resources : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity - Manage notification of events
 *
 * API version: 5.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

type CharacteristicMvo struct {
	// Unique identifier of the characteristic
	Id string `json:"id,omitempty"`
	// Name of the characteristic
	Name string `json:"name"`
	// Data type of the value of the characteristic
	ValueType string `json:"valueType,omitempty"`

	CharacteristicRelationship []CharacteristicRelationshipMvo `json:"characteristicRelationship,omitempty"`
	// When sub-classing, this defines the sub-class Extensible name
	Type_ string `json:"@type"`
	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`
	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty"`
}
