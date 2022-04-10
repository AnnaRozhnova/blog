package handler

import (
	"github.com/AnnaRozhnova/blog/pkg/service"
	"github.com/gin-gonic/gin"
)

const usernameCtx = "userId"

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth", h.addHeader)
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-out", h.signOut)
	}

	posts := router.Group("/posts", h.addHeader)
	{
		posts.POST("/create", h.createPost)
		posts.GET("/", h.getAllPosts)
		posts.GET("/post/:id", h.getPostById)
		posts.GET("/:username", h.getPostByUsername)
	}

	comments := router.Group("/comments", h.addHeader) 
	{
		comments.POST("/create", h.createComment)
		comments.GET("/:id", h.getComments)
	}
	return router
}
