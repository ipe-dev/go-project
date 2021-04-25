package post

type CreateRequest struct {
	Title    string `json:"title"`
	UserName string `json:"user_name"`
	Content  string `json:"content"`
}
