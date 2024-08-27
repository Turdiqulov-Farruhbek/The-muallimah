package entity

type ProductCreate struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       float32  `json:"price"`
	PictureUrls []string `json:"picture_urls"`
}

type ProductUpt struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

type ProductUpdate struct {
	Id   string     `json:"id"`
	Body ProductUpt `json:"body"`
}

type ProductPicture struct {
	ProductId  string `json:"product_id"`
	PictureUrl string `json:"picture_url"`
}

type ProductGet struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       float32  `json:"price"`
	PictureUrls []string `json:"picture_urls"`
	CreatedAt   string   `json:"created_at"`
}

type ProductFilter struct {
	Title      string     `json:"title"`
	PriceFrom  float32    `json:"price_from"`
	PriceTo    float32    `json:"price_to"`
	Pagination Pagination `json:"pagination"`
}

type ProductList struct {
	Products   []*ProductGet `json:"products"`
	TotalCount int64        `json:"total_count"`
	Pagination Pagination   `json:"pagination"`
}