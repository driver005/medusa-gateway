/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Currency - Currency
type Currency struct {

	// The 3 character ISO code for the currency.
	Code string `json:"code"`

	// The symbol used to indicate the currency.
	Symbol string `json:"symbol"`

	// The native symbol used to indicate the currency.
	SymbolNative string `json:"symbol_native"`

	// The written name of the currency
	Name string `json:"name"`

	// [EXPERIMENTAL] Does the currency prices include tax
	IncludesTax bool `json:"includes_tax,omitempty"`
}