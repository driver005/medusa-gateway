/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PostOrdersOrderClaimsClaimRequest struct {

	// The Claim Items that the Claim will consist of.
	ClaimItems []PostOrdersOrderClaimsClaimRequestClaimItemsInner `json:"claim_items,omitempty"`

	// The Shipping Methods to send the additional Line Items with.
	ShippingMethods []PostOrdersOrderClaimsRequestShippingMethodsInner `json:"shipping_methods,omitempty"`

	// If set to true no notification will be send related to this Swap.
	NoNotification bool `json:"no_notification,omitempty"`

	// An optional set of key-value pairs to hold additional information.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}
