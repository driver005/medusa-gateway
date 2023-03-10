/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GetBatchJobsConfirmedAtParameter struct {

	// filter by dates less than this date
	Lt string `json:"lt,omitempty"`

	// filter by dates greater than this date
	Gt string `json:"gt,omitempty"`

	// filter by dates less than or equal to this date
	Lte string `json:"lte,omitempty"`

	// filter by dates greater than or equal to this date
	Gte string `json:"gte,omitempty"`
}
