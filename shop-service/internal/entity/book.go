package entity

type BookCreate struct {
	Title       string  
	Description string  
	Price       float32 
	PictureUrls []string
}

type BookUpt struct {
	Title       string 
	Description string 
	Price       float32
}

type BookUpdate struct {
	Id   string 
	Body BookUpt
}

type BookPicture struct {
	BookId     string
	PictureUrl string
}

type BookGet struct {
	Id          string  
	Title       string  
	Description string  
	Price       float32 
	PictureUrls []string
	CreatedAt   string  
}

type BookFilter struct {
	Title     string 
	PriceFrom float32
	PriceTo   float32
	Pagination Pagination 
}

type BookList struct {
	Books      []*BookGet
	TotalCount int64    
}
