package model

import "database/sql"

type Product struct {
	ID          int64  `gorm:"primary_key;not_null;auto_increment"`
	ProductName string `json:"product_name"`
	ProductSku  sql.NullString `gorm:"unique_index;not_null" json:"product_sku"`
	ProductPrice float64 `json:"product_price"`
	ProductDescription string `json:"product_description"`
	ProductImage []ProductImage `gorm:"foreignKey:ImageProductId" json:"product_image"`
	ProductSize []ProductSize `gorm:"foreignKey:SizeProductId" json:"product_size"`
	ProductSeo ProductSeo `gorm:"foreignKey:SeoProductId" json:"product_seo"`
}
