package handler

import (
	"net/http"
	"strings"

	"github.com/AnnaRozhnova/blog"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

// NewCookieStore returns a new CookieStore
var store = sessions.NewCookieStore([]byte("secret-key"))

// signUp creates new user
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

	// create a new session
	session, err := store.New(c.Request, user.Username)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// add session to the response
	session.Save(c.Request, c.Writer)

	c.JSON(http.StatusOK, map[string]interface{}{"username": user.Username})
}


// sign-in handler
func (h *Handler) signIn(c *gin.Context) {
	var input blog.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.service.Authorization.GetUser(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// create a new session
	session, err := store.New(c.Request, user.Username)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// add session to the response
	session.Save(c.Request, c.Writer)


	c.JSON(http.StatusOK, map[string]interface{}{"username": user.Username})
}


// sign-out handler
func(h *Handler) signOut(c *gin.Context) {
	cookie := c.Request.Header["Cookie"][0]

	// get username from the cookie
	username := strings.Split(cookie, "=")
	session, _ := store.Get(c.Request, username[0])
	// set MaxAge to -1 to delete cookie from the store
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)

	c.JSON(http.StatusOK, map[string]interface{}{"max_age": session.Options.MaxAge})
	
}