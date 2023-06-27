package handler

import (
	"net/http"
	entity "sport_app"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO: Тренировку НЕ МОЖЕТ создать кто попало
func (h *Handler) createTraining(c *gin.Context) {
	coachId, err := getUserId(c, true)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var input entity.Training
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.CoachId = coachId

	id, err := h.services.Training.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(200, map[string]any{
		"id": id,
	})
}

func (h *Handler) getTrainingById(c *gin.Context) {
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

	training, err := h.services.Training.GetById(id_param)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(200, training)
}
