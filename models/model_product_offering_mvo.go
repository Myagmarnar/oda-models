/*
 * Product Catalog Management
 *
 * Product Catalog API is one of Catalog Management API Family. Product Catalog API goal is to provide a catalog of products.  ### Operations Product Catalog API performs the following operations on the resources : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity - Manage notification of events
 *
 * API version: 5.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

type ProductOfferingMvo struct {
	// Description of the productOffering
	Description string `json:"description,omitempty"`
	// isBundle determines whether a productOffering represents a single productOffering (false), or a bundle of productOfferings (true).
	IsBundle bool `json:"isBundle,omitempty"`
	// A flag indicating if this product offer can be sold stand-alone for sale or not. If this flag is false it indicates that the offer can only be sold within a bundle.
	IsSellable bool `json:"isSellable,omitempty"`
	// A string providing a complementary information on the value of the lifecycle status attribute.
	StatusReason string `json:"statusReason,omitempty"`

	ValidFor *TimePeriod `json:"validFor,omitempty"`
	// ProductOffering version
	Version string `json:"version,omitempty"`
	// Place defines the places where the products are sold or delivered.
	Place []PlaceRefMvo `json:"place,omitempty"`

	ServiceLevelAgreement *SlaRefMvo `json:"serviceLevelAgreement,omitempty"`
	// The channel defines the channel for selling product offerings.
	Channel []ChannelRefMvo `json:"channel,omitempty"`

	ServiceCandidate *ServiceCandidateRefMvo `json:"serviceCandidate,omitempty"`
	// The category resource is used to group product offerings, service and resource candidates in logical containers. Categories can contain other categories and/or product offerings, resource or service candidates.
	Category []CategoryRefMvo `json:"category,omitempty"`

	ResourceCandidate *ResourceCandidateRefMvo `json:"resourceCandidate,omitempty"`
	// A condition under which a ProductOffering is made available to Customers. For instance, a productOffering can be offered with multiple commitment periods.
	ProductOfferingTerm []ProductOfferingTermMvo `json:"productOfferingTerm,omitempty"`
	// An amount, usually of money, that is asked for or allowed when a ProductOffering is bought, rented, or leased. The price is valid for a defined period of time and may not represent the actual price paid by a customer.
	ProductOfferingPrice []ProductOfferingPriceRefOrValueMvo `json:"productOfferingPrice,omitempty"`
	// An agreement represents a contract or arrangement, either written or verbal and sometimes enforceable by law, such as a service level agreement or a customer price agreement. An agreement involves a number of other business entities, such as products, services, and resources and/or their specifications.
	Agreement []AgreementRefMvo `json:"agreement,omitempty"`
	// A type of ProductOffering that belongs to a grouping of ProductOfferings made available to the market. It inherits of all attributes of ProductOffering.
	BundledProductOffering []BundledProductOfferingMvo `json:"bundledProductOffering,omitempty"`
	// A group of ProductOfferings that can be selected for instantiation, e.g. between 2 and 7 from a list of 15 channel packs.
	BundledGroupProductOffering []BundledGroupProductOfferingMvo `json:"bundledGroupProductOffering,omitempty"`
	// Complements the description of an element (for instance a product) through video, pictures...
	Attachment []AttachmentRefOrValueMvo `json:"attachment,omitempty"`
	// provides references to the corresponding market segment as target of product offerings. A market segment is grouping of Parties, GeographicAreas, SalesChannels, and so forth.
	MarketSegment []MarketSegmentRefMvo `json:"marketSegment,omitempty"`
	// A relationship between this product offering and other product offerings.
	ProductOfferingRelationship []ProductOfferingRelationshipMvo `json:"productOfferingRelationship,omitempty"`
	// A characteristic quality or distinctive feature of a ProductOffering. The characteristic can take on a discrete value fixed at design (catalog authoring) time, such as Mobile Plan Rank, and is not generally modifiable at inventory level.
	ProductOfferingCharacteristic []CharacteristicSpecificationMvo `json:"productOfferingCharacteristic,omitempty"`
	// A use of the ProductSpecificationCharacteristicValue by a ProductOffering to which additional properties (attributes) apply or override the properties of similar properties contained in ProductSpecificationCharacteristicValue. It should be noted that characteristics which their value(s) addressed by this object must exist in corresponding product specification. The available characteristic values for a ProductSpecificationCharacteristic in a Product specification can be modified at the ProductOffering level. For example, a characteristic 'Color' might have values White, Blue, Green, and Red. But, the list of values can be restricted to e.g. White and Blue in an associated product offering. It should be noted that the list of values in 'ProductSpecificationCharacteristicValueUse' is a strict subset of the list of values as defined in the corresponding product specification characteristics.
	ProdSpecCharValueUse []ProductSpecificationCharacteristicValueUseMvo `json:"prodSpecCharValueUse,omitempty"`
	// The Policy resource represents a policy/rule applied to ProductOffering.
	Policy []PolicyRefMvo `json:"policy,omitempty"`
	// List of actions that can be executed (in context of a product order) on products instantiated from this offering
	AllowedAction []AllowedProductActionMvo `json:"allowedAction,omitempty"`
	// Used to indicate the current lifecycle status
	LifecycleStatus string `json:"lifecycleStatus,omitempty"`
	// Name of the productOffering
	Name string `json:"name,omitempty"`

	ProductSpecification *ProductSpecificationRefMvo `json:"productSpecification,omitempty"`
	// List of external identifieers for the offering, e.g. identifier in source catalog
	ExternalIdentifier []ExternalIdentifierMvo `json:"externalIdentifier,omitempty"`
	Type               string                  `json:"@type,omitempty"`
	SchemaLocation     string                  `json:"@schemaLocation,omitempty"`
	BaseType           string                  `json:"@baseType,omitempty"`
}
