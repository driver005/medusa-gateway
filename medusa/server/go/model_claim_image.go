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

// ClaimImage - Represents photo documentation of a claim.
type ClaimImage struct {

	// The claim image's ID
	Id string `json:"id,omitempty"`

	// The ID of the claim item associated with the image
	ClaimItemId string `json:"claim_item_id"`

	// A claim item object. Available if the relation `claim_item` is expanded.
	ClaimItem map[string]interface{} `json:"claim_item,omitempty"`

	// The URL of the image
	Url string `json:"url"`

	// The date with timezone at which the resource was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt time.Time `json:"deleted_at,omitempty"`

	// An optional key-value map with additional details
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}
