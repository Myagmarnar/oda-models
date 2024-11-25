/*
 * Product Catalog Management
 *
 * Product Catalog API is one of Catalog Management API Family. Product Catalog API goal is to provide a catalog of products.  ### Operations Product Catalog API performs the following operations on the resources : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity - Manage notification of events
 *
 * API version: 5.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

type AttachmentFvo struct {
	// The name of the attachment
	Name string `json:"name,omitempty"`
	// A narrative text describing the content of the attachment
	Description string `json:"description,omitempty"`
	// Uniform Resource Locator, is a web page address (a subset of URI)
	Url string `json:"url,omitempty"`
	// The actual contents of the attachment object, if embedded, encoded as base64
	Content string `json:"content,omitempty"`

	Size *Quantity `json:"size,omitempty"`

	ValidFor *TimePeriod `json:"validFor,omitempty"`
	// a business characterization of the purpose of the attachment, for example logo, instructionManual, contractCopy
	AttachmentType string `json:"attachmentType"`
	// a technical characterization of the attachment content format using IETF Mime Types
	MimeType string `json:"mimeType"`
	// unique identifier
	Id string `json:"id,omitempty"`
}
