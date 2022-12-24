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

// CustomerGroup - Represents a customer group
type CustomerGroup struct {

	// The customer group's ID
	Id string `json:"id,omitempty"`

	// The name of the customer group
	Name string `json:"name"`

	// The customers that belong to the customer group. Available if the relation `customers` is expanded.
	Customers []map[string]interface{} `json:"customers,omitempty"`

	// The price lists that are associated with the customer group. Available if the relation `price_lists` is expanded.
	PriceLists []PriceList `json:"price_lists,omitempty"`

	// The date with timezone at which the resource was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt time.Time `json:"deleted_at,omitempty"`

	// An optional key-value map with additional details
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}