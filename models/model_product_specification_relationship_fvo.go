/*
 * Product Catalog Management
 *
 * Product Catalog API is one of Catalog Management API Family. Product Catalog API goal is to provide a catalog of products.  ### Operations Product Catalog API performs the following operations on the resources : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity - Manage notification of events
 *
 * API version: 5.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

type ProductSpecificationRelationshipFvo struct {
	// A characteristic that refines the relationship. For example, consider the relationship between broadband and TV. For a 4k TV it is important to know the minimum bandwidth to support 4k, so a characteristic Resolution might be defined on the relationship to allow capturing of this in the inventory
	Characteristic []CharacteristicSpecificationFvo `json:"characteristic,omitempty"`

	ValidFor *TimePeriod `json:"validFor,omitempty"`
	// type of the relationship, for example override, discount, etc.
	RelationshipType string `json:"relationshipType"`
	// Version of the referred product specification.
	Version string `json:"version,omitempty"`
	// The identifier of the referred entity.
	Id string `json:"id"`
	// The URI of the referred entity.
	Href string `json:"href,omitempty"`
	// Name of the referred entity.
	Name string `json:"name,omitempty"`
	// The actual type of the target instance when needed for disambiguation.
	ReferredType   string `json:"@referredType,omitempty"`
	Type           string `json:"@type,omitempty"`
	SchemaLocation string `json:"@schemaLocation,omitempty"`
	BaseType       string `json:"@baseType,omitempty"`
}
