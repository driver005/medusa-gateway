/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type DeleteDiscountsDiscount200Response struct {

	// The ID of the deleted Discount
	Id string `json:"id,omitempty"`

	// The type of the object that was deleted.
	Object string `json:"object,omitempty"`

	// Whether the discount was deleted successfully or not.
	Deleted bool `json:"deleted,omitempty"`
}
