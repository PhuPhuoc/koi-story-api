package marketmodel

type PostMarket struct {
	PostId      string  `db:"post_id" json:"post_id"`
	ProductName string  `db:"product_name" json:"product_name"`
	Price       float64 `db:"price" json:"price"`
	FilePath    string  `db:"file_path" json:"file_path"`
}
