package handler

import (
	"net/http"
	"strconv"

	"github.com/AnnaRozhnova/blog"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createPost(c *gin.Context) {
	var post blog.Post
	if err := c.BindJSON(&post); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	/*
	session, _ := store.Get(c.Request, post.Username)
	username, ok := session.Values[usernameCtx]
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "Session values error")
		return
	}
	post.Username = username.(string)
*/
	id, err := h.service.Post.Create(post)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id, "username": post.Username})

}

func (h *Handler) getAllPosts(c *gin.Context) {
	posts, err := h.service.Post.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, posts)
}




func (h *Handler) getPostById(c *gin.Context) {
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




func (h *Handler) getPostByUsername(c *gin.Context) {
	username:= c.Param("username")
	posts, err := h.service.Post.GetByUsername(username)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"posts": posts})
}
