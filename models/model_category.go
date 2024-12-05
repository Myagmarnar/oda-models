/*
 * Product Catalog Management
 *
 * Product Catalog API is one of Catalog Management API Family. Product Catalog API goal is to provide a catalog of products.  ### Operations Product Catalog API performs the following operations on the resources : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity - Manage notification of events
 *
 * API version: 5.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

import (
	"time"
)

type Category struct {
	// Description of the category
	Description string `json:"description,omitempty"`
	// If true, this Boolean indicates that the category is a root of categories
	IsRoot bool `json:"isRoot,omitempty"`

	Parent *CategoryRef `json:"parent,omitempty"`
	// List of product offerings that are referred to by the category
	ProductOffering []ProductOfferingRef `json:"productOffering,omitempty"`
	// The category resource is used to group product offerings, service and resource candidates in logical containers. Categories can contain other (sub-)categories and/or product offerings.
	SubCategory []CategoryRef `json:"subCategory,omitempty"`

	ValidFor *TimePeriod `json:"validFor,omitempty"`
	// Category version
	Version string `json:"version,omitempty"`
	// Date and time of the last update
	LastUpdate time.Time `json:"lastUpdate,omitempty"`
	// Used to indicate the current lifecycle status
	LifecycleStatus string `json:"lifecycleStatus,omitempty"`
	// Name of the category
	Name string `json:"name,omitempty"`
	// Hyperlink reference
	Href string `json:"href,omitempty"`
	// unique identifier
	Id             string `json:"id,omitempty"`
	Type           string `json:"@type,omitempty"`
	SchemaLocation string `json:"@schemaLocation,omitempty"`
	BaseType       string `json:"@baseType,omitempty"`
}
