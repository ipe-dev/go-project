package post

type GetRequest struct {
	ID int `json:"id"`
}

type ListRequest struct {
	Word  string `json:"Word"`
	TagID int    `json:"tag_id"`
}
