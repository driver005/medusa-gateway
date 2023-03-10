/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PostCollectionsCollectionRequest struct {

	// The title to identify the Collection by.
	Title string `json:"title,omitempty"`

	// An optional handle to be used in slugs, if none is provided we will kebab-case the title.
	Handle string `json:"handle,omitempty"`

	// An optional set of key-value pairs to hold additional information.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}
