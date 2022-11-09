package models

import "time"

// Product Variants represent a Product with a specific set of Product Option configurations. The maximum number of Product Variants that a Product can have is given by the number of available Product Option combinations.
type ProductVariant struct {
	// Whether the Product Variant should be purchasable when `inventory_quantity` is 0.
	AllowBackorder *bool `json:"allow_backorder,omitempty"`

	// A generic field for a GTIN number that can be used to identify the Product Variant.
	Barcode *string `json:"barcode,omitempty"`

	// An EAN barcode number that can be used to identify the Product Variant.
	Ean *string `json:"ean,omitempty"`

	// The height of the Product Variant. May be used in shipping rate calculations.
	Height *float32 `json:"height,omitempty"`

	// The Harmonized System code of the Product Variant. May be used by Fulfillment Providers to pass customs information to shipping carriers.
	HsCode *string `json:"hs_code,omitempty"`

	// The product variant's ID
	Id *string `json:"id,omitempty"`

	// The current quantity of the item that is stocked.
	InventoryQuantity int `json:"inventory_quantity"`

	// The length of the Product Variant. May be used in shipping rate calculations.
	Length *float32 `json:"length,omitempty"`

	// Whether Medusa should manage inventory for the Product Variant.
	ManageInventory *bool `json:"manage_inventory,omitempty"`

	// The material and composition that the Product Variant is made of, May be used by Fulfillment Providers to pass customs information to shipping carriers.
	Material *string `json:"material,omitempty"`

	// An optional key-value map with additional details
	Metadata *map[string]interface{} `json:"metadata,omitempty"`

	// The Manufacturers Identification code that identifies the manufacturer of the Product Variant. May be used by Fulfillment Providers to pass customs information to shipping carriers.
	MidCode *string `json:"mid_code,omitempty"`

	// The Product Option Values specified for the Product Variant. Available if the relation `options` is expanded.
	Options *[]ProductOptionValue `json:"options,omitempty"`

	// The country in which the Product Variant was produced. May be used by Fulfillment Providers to pass customs information to shipping carriers.
	OriginCountry *string `json:"origin_country,omitempty"`

	// The Money Amounts defined for the Product Variant. Each Money Amount represents a price in a given currency or a price in a specific Region. Available if the relation `prices` is expanded.
	Prices *[]MoneyAmount `json:"prices,omitempty"`

	// A product object. Available if the relation `product` is expanded.
	Product *map[string]interface{} `json:"product,omitempty"`

	// The ID of the Product that the Product Variant belongs to.
	ProductId string `json:"product_id"`

	// The unique stock keeping unit used to identify the Product Variant. This will usually be a unqiue identifer for the item that is to be shipped, and can be referenced across multiple systems.
	Sku *string `json:"sku,omitempty"`

	// A title that can be displayed for easy identification of the Product Variant.
	Title string `json:"title"`

	// A UPC barcode number that can be used to identify the Product Variant.
	Upc *string `json:"upc,omitempty"`

	// The ranking of this variant
	VariantRank *float32 `json:"variant_rank,omitempty"`

	// The weight of the Product Variant. May be used in shipping rate calculations.
	Weight *float32 `json:"weight,omitempty"`

	// The width of the Product Variant. May be used in shipping rate calculations.
	Width *float32 `json:"width,omitempty"`

	// The date with timezone at which the resource was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
