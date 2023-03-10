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

// ShippingMethodTaxLine - Shipping Method Tax Line
type ShippingMethodTaxLine struct {

	// The line item tax line's ID
	Id string `json:"id,omitempty"`

	// The ID of the line item
	ShippingMethodId string `json:"shipping_method_id"`

	ShippingMethod ShippingMethod `json:"shipping_method,omitempty"`

	// A code to identify the tax type by
	Code string `json:"code,omitempty"`

	// A human friendly name for the tax
	Name string `json:"name"`

	// The numeric rate to charge tax by
	Rate float32 `json:"rate"`

	// The date with timezone at which the resource was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// An optional key-value map with additional details
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}
