package entity

type PostCreate struct {
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	PictureUrls []string `json:"picture_urls"`
}

type PostGet struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	PictureUrls []string `json:"picture_urls"`
	CreatedAt   string   `json:"created_at"`
}

type PostUpt struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostUpdate struct {
	ID   string  `json:"id"`
	Body PostUpt `json:"body"`
}

type PostList struct {
	Posts      []*PostGet `json:"posts"`
	TotalCount int32     `json:"total_count"`
	Pagination Pagination `json:"pagination"`
}

type PostPicture struct {
	PostID     string `json:"post_id"`
	PictureUrl string `json:"picture_url"`
}

type Pagination struct {
	Limit     int32 `json:"limit"`
	Offset int32 `json:"offset"`
}
