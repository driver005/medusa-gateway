/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PostOrdersOrderClaimsRequest struct {

	// The type of the Claim. This will determine how the Claim is treated: `replace` Claims will result in a Fulfillment with new items being created, while a `refund` Claim will refund the amount paid for the claimed items.
	Type string `json:"type"`

	// The Claim Items that the Claim will consist of.
	ClaimItems []PostOrdersOrderClaimsRequestClaimItemsInner `json:"claim_items"`

	ReturnShipping PostOrdersOrderClaimsRequestReturnShipping `json:"return_shipping,omitempty"`

	// The new items to send to the Customer when the Claim type is Replace.
	AdditionalItems []PostOrdersOrderClaimsRequestAdditionalItemsInner `json:"additional_items,omitempty"`

	// The Shipping Methods to send the additional Line Items with.
	ShippingMethods []PostOrdersOrderClaimsRequestShippingMethodsInner `json:"shipping_methods,omitempty"`

	ShippingAddress Address `json:"shipping_address,omitempty"`

	// The amount to refund the Customer when the Claim type is `refund`.
	RefundAmount int32 `json:"refund_amount,omitempty"`

	// If set to true no notification will be send related to this Claim.
	NoNotification bool `json:"no_notification,omitempty"`

	// An optional set of key-value pairs to hold additional information.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}