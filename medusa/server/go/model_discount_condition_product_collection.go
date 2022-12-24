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

// DiscountConditionProductCollection - Associates a discount condition with a product collection
type DiscountConditionProductCollection struct {

	// The ID of the Product Collection
	ProductCollectionId string `json:"product_collection_id"`

	// The ID of the Discount Condition
	ConditionId string `json:"condition_id"`

	ProductCollection ProductCollection `json:"product_collection,omitempty"`

	DiscountCondition DiscountCondition `json:"discount_condition,omitempty"`

	// The date with timezone at which the resource was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// An optional key-value map with additional details
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}