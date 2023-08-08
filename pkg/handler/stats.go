package handler

import (
	"net/http"
	entity "sport_app"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getActivityUsageStats(c *gin.Context) {
	visitorId, err := getUserId(c, false)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var activityName string = c.Query("activityName")
	if activityName == "" {
		newErrorResponse(c, http.StatusBadRequest, "activityName is not in query")
	}
	println(activityName)
	println(visitorId)

	var result []entity.ActivityState
	result, err = h.services.Stats.GetActivityStateStats(visitorId, activityName)
	if err != nil {
		newErrorResponse(c, 500, err.Error())
	}
	println(len(result))
	c.JSON(200, result)
}
