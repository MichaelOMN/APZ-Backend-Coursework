package handler

import (
	"net/http"
	entity "sport_app"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createPhysicalInfo(c *gin.Context) {
	visitorId, err := getUserId(c, false)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var physicalInfo entity.PhysicalInfo
	if err := c.BindJSON(&physicalInfo); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	physicalInfo.VisitorId = visitorId

	id, err := h.services.PhysicalInfo.Create(physicalInfo)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": id,
	})
}

func (h *Handler) getPhysicalInfoByVisitorId(c *gin.Context) {
	userId, err := getUserId(c, false)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	// 	return
	// }

	pinfo, err := h.services.PhysicalInfo.GetByVisitorId(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, pinfo)
}

func (h *Handler) updatePhysicalInfoByVisitorId(c *gin.Context) {
	userId, err := getUserId(c, false)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// id_param, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	// 	return
	// }

	var input entity.PhysicalInfoUpdateForm
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}

	if err := input.Validate(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: err.Error()})
	}

	if err := h.services.PhysicalInfo.Update(userId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, statusResponse{Status: "ok"})
}
