package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/go-project/handler"
)

func DefineRoutes(r gin.IRouter) {
	post := r.Group("/api/post")
	{
		post.POST("/get", handler.GetPost())
		post.POST("/create", handler.CreatePost())
		post.POST("/update", handler.UpdatePost())
		post.POST("/delete", handler.DeletePost())
		post.POST("/list", handler.ListPost())
	}
	comment := r.Group("/api/comment")
	{
		comment.POST("/get", handler.GetComment())
		comment.POST("/create", handler.CreateComment())
		comment.POST("/update", handler.UpdateComment())
		comment.POST("/delete", handler.DeleteComment())
	}
	user := r.Group("/api/user")
	{
		user.POST("/get", handler.GetUser())
		user.POST("/create", handler.CreateUser())
		user.POST("/update", handler.UpdateUser())
		user.POST("/delete", handler.DeleteUser())
	}

}
