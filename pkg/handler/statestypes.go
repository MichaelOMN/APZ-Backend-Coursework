package handler

import (
	"net/http"
	entity "sport_app"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createST(c *gin.Context) {
	_, err := getUserId(c, true)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var input entity.StatesTypes
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.StatesTypes.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, map[string]any{
		"id": id,
	})
}

func (h *Handler) getSTById(c *gin.Context) {
	_, err := getUserId(c, true)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	id_param, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	st, err := h.services.StatesTypes.GetById(id_param)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, st)
}

func (h *Handler) updateSTById(c *gin.Context) {
	_, err := getUserId(c, true)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	id_param, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var update entity.StatesTypesUpdateForm
	if err := c.BindJSON(&update); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "missing some fields")
		return
	}

	err = h.services.StatesTypes.Update(id_param, update)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, statusResponse{Status: "ok"})
}
