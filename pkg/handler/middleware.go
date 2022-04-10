package handler

import "github.com/gin-gonic/gin"


func addHeader(c *gin.Context) (int, error) {
	c.Header("Access-Control-Allow-Origin", "*")
}
