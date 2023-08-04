package handler

import (
	"net/http"
	entity "sport_app"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createActUsage(c *gin.Context) {
	visitorId, err := getUserId(c, false)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var input entity.ActivityUsage
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.VisitorId = visitorId

	id, err := h.services.ActivityUsage.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, map[string]any{
		"id": id,
	})
}

func (h *Handler) getActUsageByIdAndVisitorId(c *gin.Context) {
	visitorId, err := getUserId(c, false)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	name_param := c.Query("name")

	activity_usage, err := h.services.ActivityUsage.GetById(visitorId, name_param)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, activity_usage)
}

func (h *Handler) deleteActUsageByVisitorId(c *gin.Context){
	visitorId, err := getUserId(c, false)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	id_param, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.ActivityUsage.Delete(visitorId, id_param)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, statusResponse{Status: "ok"})
}
