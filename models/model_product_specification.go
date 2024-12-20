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

type ProductSpecification struct {
	// The manufacturer or trademark of the specification
	Brand string `json:"brand,omitempty"`
	// A narrative that explains in detail what the product specification is
	Description string `json:"description,omitempty"`
	// isBundle determines whether a productSpecification represents a single productSpecification (false), or a bundle of productSpecification (true).
	IsBundle bool `json:"isBundle,omitempty"`
	// An identification number assigned to uniquely identity the specification
	ProductNumber string `json:"productNumber,omitempty"`
	// The category resource is used to group product specifications in logical containers. Categories can contain other categories and/or other catalog entries.
	Category []CategoryRef `json:"category,omitempty"`

	ValidFor *TimePeriod `json:"validFor,omitempty"`
	// Product specification version
	Version string `json:"version,omitempty"`
	// A related party defines party or party role linked to a specific entity.
	RelatedParty []RelatedPartyRefOrPartyRoleRef `json:"relatedParty,omitempty"`
	// A characteristic quality or distinctive feature of a ProductSpecification.  The characteristic can be take on a discrete value, such as color, can take on a range of values, (for example, sensitivity of 100-240 mV), or can be derived from a formula (for example, usage time (hrs) = 30 - talk time *3). Certain characteristics, such as color, may be configured during the ordering or some other process.
	ProductSpecCharacteristic []CharacteristicSpecification `json:"productSpecCharacteristic,omitempty"`
	// ServiceSpecification(s) required to realize a ProductSpecification.
	ServiceSpecification []ServiceSpecificationRef `json:"serviceSpecification,omitempty"`
	// A type of ProductSpecification that belongs to a grouping of ProductSpecifications made available to the market. It inherits of all attributes of ProductSpecification.
	BundledProductSpecification []BundledProductSpecification `json:"bundledProductSpecification,omitempty"`
	// A migration, substitution, dependency or exclusivity relationship between/among product specifications.
	ProductSpecificationRelationship []ProductSpecificationRelationship `json:"productSpecificationRelationship,omitempty"`
	// The ResourceSpecification is required to realize a ProductSpecification.
	ResourceSpecification []ResourceSpecificationRef `json:"resourceSpecification,omitempty"`
	// Complements the description of an element (for instance a product) through video, pictures...
	Attachment []AttachmentRefOrValue `json:"attachment,omitempty"`
	// The Policy resource represents a policy/rule applied to ProductSpecification.
	Policy []PolicyRef `json:"policy,omitempty"`

	TargetProductSchema *TargetProductSchema `json:"targetProductSchema,omitempty"`

	IntentSpecification *IntentSpecificationRef `json:"intentSpecification,omitempty"`
	// Date and time of the last update
	LastUpdate time.Time `json:"lastUpdate,omitempty"`
	// Used to indicate the current lifecycle status
	LifecycleStatus string `json:"lifecycleStatus,omitempty"`
	// Name of the product specification
	Name string `json:"name,omitempty"`
	// List of external identifieers for the specification, e.g. identifier in source catalog
	ExternalIdentifier []ExternalIdentifier `json:"externalIdentifier,omitempty"`
	// Hyperlink reference
	Href string `json:"href,omitempty"`
	// unique identifier
	Id             string `json:"id,omitempty"`
	Type           string `json:"@type,omitempty"`
	SchemaLocation string `json:"@schemaLocation,omitempty"`
	BaseType       string `json:"@baseType,omitempty"`
}
