/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ShippingMethod - Shipping Methods represent a way in which an Order or Return can be shipped. Shipping Methods are built from a Shipping Option, but may contain additional details, that can be necessary for the Fulfillment Provider to handle the shipment.
type ShippingMethod struct {

	// The shipping method's ID
	Id string `json:"id,omitempty"`

	// The id of the Shipping Option that the Shipping Method is built from.
	ShippingOptionId string `json:"shipping_option_id"`

	ShippingOption ShippingOption `json:"shipping_option,omitempty"`

	// The id of the Order that the Shipping Method is used on.
	OrderId string `json:"order_id,omitempty"`

	// An order object. Available if the relation `order` is expanded.
	Order map[string]interface{} `json:"order,omitempty"`

	// The id of the Return that the Shipping Method is used on.
	ReturnId string `json:"return_id,omitempty"`

	// A return object. Available if the relation `return_order` is expanded.
	ReturnOrder map[string]interface{} `json:"return_order,omitempty"`

	// The id of the Swap that the Shipping Method is used on.
	SwapId string `json:"swap_id,omitempty"`

	// A swap object. Available if the relation `swap` is expanded.
	Swap map[string]interface{} `json:"swap,omitempty"`

	// The id of the Cart that the Shipping Method is used on.
	CartId string `json:"cart_id,omitempty"`

	// A cart object. Available if the relation `cart` is expanded.
	Cart map[string]interface{} `json:"cart,omitempty"`

	// The id of the Claim that the Shipping Method is used on.
	ClaimOrderId string `json:"claim_order_id,omitempty"`

	// A claim order object. Available if the relation `claim_order` is expanded.
	ClaimOrder map[string]interface{} `json:"claim_order,omitempty"`

	// Available if the relation `tax_lines` is expanded.
	TaxLines []ShippingMethodTaxLine `json:"tax_lines,omitempty"`

	// The amount to charge for the Shipping Method. The currency of the price is defined by the Region that the Order that the Shipping Method belongs to is a part of.
	Price int32 `json:"price"`

	// Additional data that the Fulfillment Provider needs to fulfill the shipment. This is used in combination with the Shipping Options data, and may contain information such as a drop point id.
	Data map[string]interface{} `json:"data,omitempty"`

	// [EXPERIMENTAL] Indicates if the shipping method price include tax
	IncludesTax bool `json:"includes_tax,omitempty"`
}