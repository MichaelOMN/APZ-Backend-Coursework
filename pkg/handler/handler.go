package handler

import (
	"sport_app/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) pingRequest(c *gin.Context) {
	logrus.Info("Server was pinged!")
	c.String(200, "Pong from web-server!")
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	router.GET("/ping", h.pingRequest)

	auth := router.Group("/auth")
	{
		coachAuth := auth.Group("/coach")
		{
			coachAuth.POST("/sign-up", h.signUpCoach)
			coachAuth.POST("/sign-in", h.signInCoach)
		}
		visitorAuth := auth.Group("/visitor")
		{
			visitorAuth.POST("/sign-up", h.signUpVisitor)
			visitorAuth.POST("/sign-in", h.signInVisitor)
		}
		activityAuth := auth.Group("/activity")
		{
			activityAuth.POST("/sign-up", h.signUpActivity)
			activityAuth.POST("/sign-in", h.signInActivity)
		}
	}

	api := router.Group("/api")
	{
		stats := api.Group("/stats", h.visitorIdentity)
		{
			stats.GET("/activity_usages", h.getActivityUsageStats)
		}

		clubs := api.Group("/club", h.coachIdentity)
		{
			clubs.POST("/", h.createClub)
			clubs.GET("/", h.getAllClubs)
			clubs.GET("/:id", h.getClubById)
			//clubs.PUT("/:id", h.updateClub)
			clubs.DELETE("/:id", h.deleteClubById)
		}

		physicalInfos := api.Group("/phys_info", h.visitorIdentity)
		{
			physicalInfos.POST("/", h.createPhysicalInfo)
			//physicalInfos.GET("/", h.getAllPhysicalInfosByVisitorId)
			physicalInfos.GET("/", h.getPhysicalInfoByVisitorId)
			physicalInfos.PUT("/", h.updatePhysicalInfoByVisitorId)
			//physicalInfos.DELETE("/:id", h.deletePhysicalInfoByVisitorId)
		}

		trainings := api.Group("/training", h.coachIdentity)
		{
			trainings.POST("/", h.createTraining)
			//trainings.GET("/", h.getAllTrainings)
			trainings.GET("/:id", h.getTrainingById)
			//trainings.DELETE("/:id", h.deleteTrainingById)
		}

		trainings2 := api.Group("/trainings", h.visitorIdentity)
		{
			trainings2.GET("/", h.getAllTrainings)
			trainings2.GET("/club/:id", h.getClubById)
		}

		statestypes := api.Group("/states_types", h.coachIdentity)
		{
			statestypes.POST("/", h.createST)
			//statestypes.GET("/", h.getAllST)
			statestypes.GET("/:id", h.getSTById)
			statestypes.PUT("/:id", h.updateSTById)
			//statestypes.DELETE("/:id", h.deleteSTById)
		}

		activity := api.Group("/activity", h.coachIdentity)
		{
			activity.POST("/", h.createActivity)
			//activity.GET("/", h.getAllActivities)
			activity.GET("/:id", h.getActivityById)
			//activity.PUT("/:id", h.updateActivityById)
			activity.DELETE("/:id", h.deleteActivityById)
		}

		attendance := api.Group("/attendance", h.visitorIdentity)
		{
			attendance.POST("/", h.createAttendance)
			//attendance.GET("/", h.getAllAttendanceByVisitorId)
			attendance.GET("/:id", h.getAttendanceByIdAndVisitorId)
			//attendance.PUT("/:id", h.updateAttendanceByVisitorId)
			//attendance.DELETE("/:id", h.deleteAttendanceByVisitorId)
		}

		actUsage := api.Group("/activity_usage", h.visitorIdentity)
		{
			actUsage.POST("/", h.createActUsage)
			//actUsage.GET("/", h.getAllActUsageByVisitorId)
			actUsage.GET("/", h.getActUsageByIdAndVisitorId)
			//actUsage.PUT("/:id", h.updateActUsageByVisitorId)
			actUsage.DELETE("/:id", h.deleteActUsageByVisitorId)
		}

		pst := api.Group("/physical_state", h.visitorIdentity)
		{
			pst.POST("/", h.createPST)
			//pst.GET("/", h.getAllPSTByVisitorId)
			pst.GET("/:id", h.getPSTByVisitorId)
			//pst.DELETE("/:id", h.deletePSTByVisitorId)
		}

		ast := api.Group("/activity_state")
		{
			ast.POST("/", h.activityIdentity, h.createAST)
			//ast.GET("/", h.getAllASTByVisitorId)
			ast.GET("/:id", h.getASTById)
			//ast.DELETE("/:id", h.deleteASTByVisitorId)
		}

	}

	return router
}
