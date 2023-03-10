/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// ClaimItem - Represents a claimed item along with information about the reasons for the claim.
type ClaimItem struct {

	// The claim item's ID
	Id string `json:"id,omitempty"`

	// Available if the relation `images` is expanded.
	Images []ClaimImage `json:"images,omitempty"`

	// The ID of the claim this item is associated with.
	ClaimOrderId string `json:"claim_order_id"`

	// A claim order object. Available if the relation `claim_order` is expanded.
	ClaimOrder map[string]interface{} `json:"claim_order,omitempty"`

	// The ID of the line item that the claim item refers to.
	ItemId string `json:"item_id"`

	Item LineItem `json:"item,omitempty"`

	// The ID of the product variant that is claimed.
	VariantId string `json:"variant_id"`

	// A variant object. Available if the relation `variant` is expanded.
	Variant map[string]interface{} `json:"variant,omitempty"`

	// The reason for the claim
	Reason string `json:"reason"`

	// An optional note about the claim, for additional information
	Note string `json:"note,omitempty"`

	// The quantity of the item that is being claimed; must be less than or equal to the amount purchased in the original order.
	Quantity int32 `json:"quantity"`

	// User defined tags for easy filtering and grouping. Available if the relation 'tags' is expanded.
	Tags []ClaimTag `json:"tags,omitempty"`

	// The date with timezone at which the resource was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt time.Time `json:"deleted_at,omitempty"`

	// An optional key-value map with additional details
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}
