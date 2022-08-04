package handler

import (
	"net/http"

	"strconv"

	"github.com/AnnaRozhnova/blog"
	"github.com/gin-gonic/gin"
)

// getComments handles /comments/:id requests
func (h *Handler) getComments(c *gin.Context) {
	// get id of the post from URL and convert it to int
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	comments, err := h.service.Comment.GetByPostId(postId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, comments)
}



// createComment adds comment to the post
func (h *Handler) createComment(c *gin.Context) {
	// check if the user logged in
	if c.Request.Header["Cookie"][0] == "" {
		newErrorResponse(c, http.StatusBadRequest, "Loged out")
		return
	}

	var comment blog.Comment
	if err := c.BindJSON(&comment); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	err := h.service.Comment.Create(comment)
	if err!= nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK,  map[string]interface{}{"isOk": "ok"})
}


