package handler

import (
	"net/http"
	entity "sport_app"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUpVisitor(c *gin.Context) {
	var input entity.Visitor

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authorization.CreateVisitor(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signUpActivity(c *gin.Context) {
	var input entity.Activity

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Activity.Create(input)
	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"id": id,
	})
}

type visitorSignInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signInVisitor(c *gin.Context) {
	var input visitorSignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Login, input.Password, false)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}

func (h *Handler) signInActivity(c *gin.Context) {
	var input entity.Activity

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateTokenForActivity(input.Name)

	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func (h *Handler) signInCoach(c *gin.Context) {
	var input visitorSignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Login, input.Password, true)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}

func (h *Handler) signUpCoach(c *gin.Context) {
	var input entity.Coach

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authorization.CreateCoach(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
