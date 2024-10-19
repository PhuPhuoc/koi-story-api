package marketmodel

type NewPostMarket struct {
	Post
	DetailMarket
	KoiInfo
	ListImage []string `json:"list_image"`
}

type MarketDetails struct {
	ID string `json:"id"`
	Post
	DetailMarket
	KoiInfo
	ListImage []PostImage
}

type Post struct {
	ID        string `db:"id" json:"-"`
	UserID    string `db:"user_id" json:"user_id"`
	PostType  string `db:"post_type" json:"post_type"`
	Title     string `db:"title" json:"title"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

type DetailMarket struct {
	ID            string  `db:"id" json:"-"`
	PostID        string  `db:"post_id" json:"-"`
	ProductName   string  `db:"product_name" json:"product_name"`
	ProductType   string  `db:"product_type" json:"product_type"`
	Price         float64 `db:"price" json:"price"`
	SellerAddress string  `db:"seller_address" json:"seller_address"`
	PhoneNumber   string  `db:"phone_number" json:"phone_number"`
	Description   string  `db:"description" json:"description"`
	CreatedAt     string  `db:"created_at" json:"-"`
}

type KoiInfo struct {
	ID             string `db:"id" json:"-"`
	DetailMarketID string `db:"detail_market_id" json:"-"`
	Color          string `db:"color" json:"color"`
	Size           string `db:"size" json:"size"`
	Old            string `db:"old" json:"old"`
	Type           string `db:"type" json:"type"`
	CreatedAt      string `db:"created_at" json:"-"`
}

type PostImage struct {
	ID         string `db:"id" json:"-"`
	PostID     string `db:"post_id" json:"-"`
	FilePath   string `db:"file_path" json:"file_path"`
	ImageOrder int    `db:"image_order" json:"image_order"`
}
