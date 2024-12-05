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

type ProductSpecificationAttributeValueChangeEvent struct {
	Event *ProductSpecificationAttributeValueChangeEventPayload `json:"event,omitempty"`
	// Hyperlink reference
	Href string `json:"href,omitempty"`
	// unique identifier
	Id string `json:"id,omitempty"`
	// The correlation id for this event.
	CorrelationId string `json:"correlationId,omitempty"`
	// The domain of the event.
	Domain string `json:"domain,omitempty"`
	// The title of the event.
	Title string `json:"title,omitempty"`
	// An explnatory of the event.
	Description string `json:"description,omitempty"`
	// A priority.
	Priority string `json:"priority,omitempty"`
	// The time the event occurred.
	TimeOccurred time.Time `json:"timeOccurred,omitempty"`

	Source *EntityRef `json:"source,omitempty"`

	ReportingSystem *EntityRef `json:"reportingSystem,omitempty"`

	RelatedParty []RelatedPartyRefOrPartyRoleRef `json:"relatedParty,omitempty"`

	AnalyticCharacteristic []Characteristic `json:"analyticCharacteristic,omitempty"`
	// The identifier of the notification.
	EventId string `json:"eventId,omitempty"`
	// Time of the event occurrence.
	EventTime time.Time `json:"eventTime,omitempty"`
	// The type of the notification.
	EventType      string `json:"eventType,omitempty"`
	Type           string `json:"@type,omitempty"`
	SchemaLocation string `json:"@schemaLocation,omitempty"`
	BaseType       string `json:"@baseType,omitempty"`
}
