/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PostOrdersOrderReturnsRequest struct {

	// The Line Items that will be returned.
	Items []PostOrdersOrderReturnsRequestItemsInner `json:"items"`

	ReturnShipping PostOrdersOrderReturnsRequestReturnShipping `json:"return_shipping,omitempty"`

	// An optional note with information about the Return.
	Note string `json:"note,omitempty"`

	// A flag to indicate if the Return should be registerd as received immediately.
	ReceiveNow bool `json:"receive_now,omitempty"`

	// A flag to indicate if no notifications should be emitted related to the requested Return.
	NoNotification bool `json:"no_notification,omitempty"`

	// The amount to refund.
	Refund int32 `json:"refund,omitempty"`
}
