package handler

import (
	"github.com/AnnaRozhnova/blog/pkg/service"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)


type Handler struct {
	service *service.Service
}

// NewHandler creates new Handler instance
func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

// InitRoutes creates a new router group
func (h *Handler) InitRoutes() *gin.Engine {

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "https://rozhnova-client.herokuapp.com"},
		AllowCredentials: true,
	})
	
	router := gin.New()
	router.Use(c)


	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/sign-out", h.signOut)
	}
	
	users := router.Group("/users") 
	{
		users.GET("/", h.getAllUsers)
		users.GET("/:username", h.getUserByUsername)
	}

	posts := router.Group("/posts")
	{
		posts.POST("/create", h.createPost)
		posts.GET("/", h.getAllPosts)
		posts.GET("/post/:id", h.getPostById)
		posts.GET("/:username", h.getPostByUsername)
	}

	comments := router.Group("/comments") 
	{
		comments.POST("/create", h.createComment)
		comments.GET("/:id", h.getComments)
	}
	return router
}
