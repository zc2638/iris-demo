package model

// 产品
type Product struct {
	AutoID
	Name string `json:"name"` // 产品名称
	Timestamps
}

// 产品型号
type ProductModel struct {
	AutoID
	ProductID uint    `json:"product_id"` // 产品id
	Name      string  `json:"name"`       // 型号名称
	Product   Product `gorm:"ForeignKey:ID;AssociationForeignKey:ProductID"`
	Timestamps
}
