package handler

import (
	"net/http"

	"github.com/AnnaRozhnova/blog"
	"github.com/gin-gonic/gin"
)

// NewCookieStore returns a new CookieStore
//var store = sessions.NewCookieStore([]byte("secret-key"))

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body blog.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var user blog.User
	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.Authorization.CreateUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"username": user.Username})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body blog.User true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	token, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}



/*
func(h *Handler) signOut(c *gin.Context) {
	cookie := c.Request.Header["Cookie"][0]

	// get username from the cookie
	username := strings.Split(cookie, "=")
	session, _ := store.Get(c.Request, username[0])
	// set MaxAge to -1 to delete cookie from the store
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)

	c.JSON(http.StatusOK, map[string]interface{}{"max_age": session.Options.MaxAge})
	
}*/