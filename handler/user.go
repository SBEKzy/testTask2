package handler

import (
	"github.com/SBEKzy/testTask2/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "OK!")
}

func (h *handler) GetUser(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, "email is required")
		return
	}

	user, err := h.service.GetUser(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *handler) UpdateUser(c *gin.Context) {
	var user model.UpdateUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.UpdateUser(user.Email, &user.Replacement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "OK!")
}

func (h *handler) DeleteUser(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, "email is required")
		return
	}

	err := h.service.DeleteUser(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "OK!")
}
