/*
 * Product Catalog Management
 *
 * Product Catalog API is one of Catalog Management API Family. Product Catalog API goal is to provide a catalog of products.  ### Operations Product Catalog API performs the following operations on the resources : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity - Manage notification of events
 *
 * API version: 5.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

type CategoryRefMvo struct {
	// Version of the category
	Version string `json:"version,omitempty"`
	// The URI of the referred entity.
	Href string `json:"href,omitempty"`
	// The identifier of the referred entity.
	Id string `json:"id"`
	// Name of the referred entity.
	Name string `json:"name,omitempty"`
	// The actual type of the target instance when needed for disambiguation.
	ReferredType   string `json:"@referredType,omitempty"`
	Type           string `json:"@type,omitempty"`
	SchemaLocation string `json:"@schemaLocation,omitempty"`
	BaseType       string `json:"@baseType,omitempty"`
}
