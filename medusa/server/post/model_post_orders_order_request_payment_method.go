/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PostOrdersOrderRequestPaymentMethod - payment method chosen for the order
type PostOrdersOrderRequestPaymentMethod struct {

	// ID of the payment provider
	ProviderId string `json:"provider_id,omitempty"`

	// Data relevant for the given payment method
	Data map[string]interface{} `json:"data,omitempty"`
}
