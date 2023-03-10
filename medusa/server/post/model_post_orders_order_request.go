/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PostOrdersOrderRequest struct {

	// the email for the order
	Email string `json:"email,omitempty"`

	BillingAddress PostOrdersOrderRequestBillingAddress `json:"billing_address,omitempty"`

	ShippingAddress PostOrdersOrderRequestShippingAddress `json:"shipping_address,omitempty"`

	// The Line Items for the order
	Items []LineItem `json:"items,omitempty"`

	// ID of the region where the order belongs
	Region string `json:"region,omitempty"`

	// Discounts applied to the order
	Discounts []Discount `json:"discounts,omitempty"`

	// ID of the customer
	CustomerId string `json:"customer_id,omitempty"`

	PaymentMethod PostOrdersOrderRequestPaymentMethod `json:"payment_method,omitempty"`

	ShippingMethod PostOrdersOrderRequestShippingMethod `json:"shipping_method,omitempty"`

	// A flag to indicate if no notifications should be emitted related to the updated order.
	NoNotification bool `json:"no_notification,omitempty"`
}
