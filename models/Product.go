package models

import "time"

type Product struct {
	// A product collection object. Available if the relation `collection` is expanded.
	Collection *map[string]interface{} `json:"collection,omitempty"`

	// The Product Collection that the Product belongs to
	CollectionId *string `json:"collection_id,omitempty"`

	// The date with timezone at which the resource was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	// A short description of the Product.
	Description *string `json:"description,omitempty"`

	// Whether the Product can be discounted. Discounts will not apply to Line Items of this Product when this flag is set to `false`.
	Discountable *bool `json:"discountable,omitempty"`

	// The external ID of the product
	ExternalId *string `json:"external_id,omitempty"`

	// A unique identifier for the Product (e.g. for slug structure).
	Handle *string `json:"handle,omitempty"`

	// The height of the Product Variant. May be used in shipping rate calculations.
	Height *float32 `json:"height,omitempty"`

	// The Harmonized System code of the Product Variant. May be used by Fulfillment Providers to pass customs information to shipping carriers.
	HsCode *string `json:"hs_code,omitempty"`

	// The product's ID
	Id *string `json:"id,omitempty"`

	// Images of the Product. Available if the relation `images` is expanded.
	Images *[]Image `json:"images,omitempty"`

	// Whether the Product represents a Gift Card. Products that represent Gift Cards will automatically generate a redeemable Gift Card code once they are purchased.
	IsGiftcard *bool `json:"is_giftcard,omitempty"`

	// The length of the Product Variant. May be used in shipping rate calculations.
	Length *float32 `json:"length,omitempty"`

	// The material and composition that the Product Variant is made of, May be used by Fulfillment Providers to pass customs information to shipping carriers.
	Material *string `json:"material,omitempty"`

	// An optional key-value map with additional details
	Metadata *map[string]interface{} `json:"metadata,omitempty"`

	// The Manufacturers Identification code that identifies the manufacturer of the Product Variant. May be used by Fulfillment Providers to pass customs information to shipping carriers.
	MidCode *string `json:"mid_code,omitempty"`

	// The Product Options that are defined for the Product. Product Variants of the Product will have a unique combination of Product Option Values. Available if the relation `options` is expanded.
	Options *[]ProductOption `json:"options,omitempty"`

	// The country in which the Product Variant was produced. May be used by Fulfillment Providers to pass customs information to shipping carriers.
	OriginCountry *string `json:"origin_country,omitempty"`

	// Shipping Profiles have a set of defined Shipping Options that can be used to fulfill a given set of Products.
	Profile *ShippingProfile `json:"profile,omitempty"`

	// The ID of the Shipping Profile that the Product belongs to. Shipping Profiles have a set of defined Shipping Options that can be used to Fulfill a given set of Products.
	ProfileId string `json:"profile_id"`

	// The sales channels the product is associated with. Available if the relation `sales_channels` is expanded.
	SalesChannels *[]map[string]interface{} `json:"sales_channels,omitempty"`

	// The status of the product
	Status *ProductStatus `json:"status,omitempty"`

	// An optional subtitle that can be used to further specify the Product.
	Subtitle *string `json:"subtitle,omitempty"`

	// The Product Tags assigned to the Product. Available if the relation `tags` is expanded.
	Tags *[]ProductTag `json:"tags,omitempty"`

	// A URL to an image file that can be used to identify the Product.
	Thumbnail *string `json:"thumbnail,omitempty"`

	// A title that can be displayed for easy identification of the Product.
	Title string `json:"title"`

	// Product Type can be added to Products for filtering and reporting purposes.
	Type *ProductType `json:"type,omitempty"`

	// The Product type that the Product belongs to
	TypeId *string `json:"type_id,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// The Product Variants that belong to the Product. Each will have a unique combination of Product Option Values. Available if the relation `variants` is expanded.
	Variants *[]ProductVariant `json:"variants,omitempty"`

	// The weight of the Product Variant. May be used in shipping rate calculations.
	Weight *float32 `json:"weight,omitempty"`

	// The width of the Product Variant. May be used in shipping rate calculations.
	Width *float32 `json:"width,omitempty"`
}