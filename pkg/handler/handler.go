package handler

import (
	"github.com/AnnaRozhnova/blog/pkg/service"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	//"github.com/swaggo/gin-swagger/swaggerFiles"

	//"github.com/swaggo/gin-swagger/swaggerFiles"

	//"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/AnnaRozhnova/blog/docs"
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
	
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		//auth.GET("/sign-out", h.signOut)
	}
	
	users := router.Group("/users", h.userIdentity) 
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

	comments := router.Group("/comments", h.userIdentity) 
	{
		comments.POST("/create", h.createComment)
		comments.GET("/:id", h.getComments)
	}
	return router
}
