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

// ProductOptionValue - A value given to a Product Variant's option set. Product Variant have a Product Option Value for each of the Product Options defined on the Product.
type ProductOptionValue struct {

	// The product option value's ID
	Id string `json:"id,omitempty"`

	// The value that the Product Variant has defined for the specific Product Option (e.g. if the Product Option is \"Size\" this value could be \"Small\", \"Medium\" or \"Large\").
	Value string `json:"value"`

	// The ID of the Product Option that the Product Option Value is defined for.
	OptionId string `json:"option_id"`

	Option ProductOption `json:"option,omitempty"`

	// The ID of the Product Variant that the Product Option Value is defined for.
	VariantId string `json:"variant_id"`

	Variant ProductVariant `json:"variant,omitempty"`

	// The date with timezone at which the resource was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt time.Time `json:"deleted_at,omitempty"`

	// An optional key-value map with additional details
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}
