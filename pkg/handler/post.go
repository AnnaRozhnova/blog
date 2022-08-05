package handler

import (
	"net/http"
	"strconv"

	"github.com/AnnaRozhnova/blog"
	"github.com/gin-gonic/gin"
)

// creates new post
func (h *Handler) createPost(c *gin.Context) {
	// check if the user logged in
	if c.Request.Header["Cookie"][0] == "" {
		newErrorResponse(c, http.StatusBadRequest, "Loged out")
		return
	}
	
	var post blog.Post
	if err := c.BindJSON(&post); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Post.Create(post)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id, "username": post.Username})

}

// getAllPosts handles /posts/ requests
func (h *Handler) getAllPosts(c *gin.Context) {
	posts, err := h.service.Post.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, posts)
}



// getPostById handles /posts/:id requests
func (h *Handler) getPostById(c *gin.Context) {
	// get id of the post from URL and convert it to int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	post, err := h.service.Post.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, map[string]interface{}{"post": post})
}



// getPostByUsername handles /posts/:username requests
func (h *Handler) getPostByUsername(c *gin.Context) {
	username:= c.Param("username")
	posts, err := h.service.Post.GetByUsername(username)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"posts": posts})
}
