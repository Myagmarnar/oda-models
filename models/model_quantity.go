/*
 * Product Catalog Management
 *
 * Product Catalog API is one of Catalog Management API Family. Product Catalog API goal is to provide a catalog of products.  ### Operations Product Catalog API performs the following operations on the resources : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity - Manage notification of events
 *
 * API version: 5.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

// An amount in a given unit
type Quantity struct {
	// Numeric value in a given unit
	Amount float32 `json:"amount,omitempty"`
	// Unit
	Units          string `json:"units,omitempty"`
	Type           string `json:"@type,omitempty"`
	SchemaLocation string `json:"@schemaLocation,omitempty"`
	BaseType       string `json:"@baseType,omitempty"`
}
