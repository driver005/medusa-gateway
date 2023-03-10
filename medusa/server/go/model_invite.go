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

// Invite - Represents an invite
type Invite struct {

	// The invite's ID
	Id string `json:"id,omitempty"`

	// The email of the user being invited.
	UserEmail string `json:"user_email"`

	// The user's role.
	Role string `json:"role,omitempty"`

	// Whether the invite was accepted or not.
	Accepted bool `json:"accepted,omitempty"`

	// The token used to accept the invite.
	Token string `json:"token,omitempty"`

	// The date the invite expires at.
	ExporesAt time.Time `json:"expores_at,omitempty"`

	// The date with timezone at which the resource was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt time.Time `json:"deleted_at,omitempty"`

	// An optional key-value map with additional details
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}
