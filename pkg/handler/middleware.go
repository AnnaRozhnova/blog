package handler

import "github.com/gin-gonic/gin"


func (h *Handler) addHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
}
