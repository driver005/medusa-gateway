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

// Swap - Swaps can be created when a Customer wishes to exchange Products that they have purchased to different Products. Swaps consist of a Return of previously purchased Products and a Fulfillment of new Products, the amount paid for the Products being returned will be used towards payment for the new Products. In the case where the amount paid for the the Products being returned exceed the amount to be paid for the new Products, a Refund will be issued for the difference.
type Swap struct {

	// The swap's ID
	Id string `json:"id,omitempty"`

	// The status of the Fulfillment of the Swap.
	FulfillmentStatus string `json:"fulfillment_status"`

	// The status of the Payment of the Swap. The payment may either refer to the refund of an amount or the authorization of a new amount.
	PaymentStatus string `json:"payment_status"`

	// The ID of the Order where the Line Items to be returned where purchased.
	OrderId string `json:"order_id"`

	// An order object. Available if the relation `order` is expanded.
	Order map[string]interface{} `json:"order,omitempty"`

	// The new Line Items to ship to the Customer. Available if the relation `additional_items` is expanded.
	AdditionalItems []LineItem `json:"additional_items,omitempty"`

	// A return order object. The Return that is issued for the return part of the Swap. Available if the relation `return_order` is expanded.
	ReturnOrder map[string]interface{} `json:"return_order,omitempty"`

	// The Fulfillments used to send the new Line Items. Available if the relation `fulfillments` is expanded.
	Fulfillments []Fulfillment `json:"fulfillments,omitempty"`

	Payment Payment `json:"payment,omitempty"`

	// The difference that is paid or refunded as a result of the Swap. May be negative when the amount paid for the returned items exceed the total of the new Products.
	DifferenceDue int32 `json:"difference_due,omitempty"`

	// The Address to send the new Line Items to - in most cases this will be the same as the shipping address on the Order.
	ShippingAddressId string `json:"shipping_address_id,omitempty"`

	ShippingAddress Address `json:"shipping_address,omitempty"`

	// The Shipping Methods used to fulfill the additional items purchased. Available if the relation `shipping_methods` is expanded.
	ShippingMethods []ShippingMethod `json:"shipping_methods,omitempty"`

	// The id of the Cart that the Customer will use to confirm the Swap.
	CartId string `json:"cart_id,omitempty"`

	// A cart object. Available if the relation `cart` is expanded.
	Cart map[string]interface{} `json:"cart,omitempty"`

	// If true, swaps can be completed with items out of stock
	AllowBackorder bool `json:"allow_backorder,omitempty"`

	// Randomly generated key used to continue the completion of the swap in case of failure.
	IdempotencyKey string `json:"idempotency_key,omitempty"`

	// The date with timezone at which the Swap was confirmed by the Customer.
	ConfirmedAt time.Time `json:"confirmed_at,omitempty"`

	// The date with timezone at which the Swap was canceled.
	CanceledAt time.Time `json:"canceled_at,omitempty"`

	// If set to true, no notification will be sent related to this swap
	NoNotification bool `json:"no_notification,omitempty"`

	// The date with timezone at which the resource was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt time.Time `json:"deleted_at,omitempty"`

	// An optional key-value map with additional details
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}