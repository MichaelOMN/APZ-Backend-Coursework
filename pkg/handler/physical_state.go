package handler

import (
	"net/http"
	entity "sport_app"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createPST(c *gin.Context) {
	visitorId, err := getUserId(c, false)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var input entity.PhysicalState
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	act_usage, err := h.services.ActivityUsage.GetByActUsageId(visitorId, input.ActivityUsageId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//additional warning
	if act_usage.VisitorId != visitorId {
		newErrorResponse(c, http.StatusInternalServerError, "activity_usage not belong to visitor_id")
		return
	}

	id, err := h.services.PhysicalState.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, map[string]any{
		"id": id,
	})
}

func (h *Handler) getPSTByVisitorId(c *gin.Context) {
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

	pst, err := h.services.PhysicalState.GetByVisitorId(visitorId, id_param)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, pst)
}
