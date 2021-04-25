package post

import (
	"net/http"

	"github.com/ipe-dev/go-project/packages/comment"
)

type Post struct {
	ID       int               `json:"id"`
	Title    string            `json:"title"`
	UserName string            `json:"user_name"`
	Content  string            `json:"content"`
	Comment  []comment.Comment `json:"comments"`
}

func New() *Post {
	return &Post{}
}
func Create(r *http.Request) error {

}

func Update() {

}

func Get() {

}

func Delete() {

}
