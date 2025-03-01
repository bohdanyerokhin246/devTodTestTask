package routes

import (
	"database/sql"
	_ "devTodTestTask/docs"
	"devTodTestTask/internal/handlers"
	"devTodTestTask/internal/repo"
	"devTodTestTask/internal/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {

	catRepo := &repo.CatRepository{DB: db}
	catService := &services.CatService{Repo: catRepo}
	catHandler := &handlers.CatHandler{Service: catService}
	missionRepo := &repo.MissionRepository{DB: db}
	missionService := &services.MissionService{Repo: missionRepo}
	missionHandler := &handlers.MissionHandler{Service: missionService}

	r.GET("/swagger", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//
	r.POST("/cats", catHandler.CreateCatHandler)
	r.GET("/cats", catHandler.ListCatsHandler)
	r.GET("/cats/:id", catHandler.CatByIDHandler)
	r.PUT("/cats", catHandler.UpdateCatHandler)
	r.DELETE("/cats/:id", catHandler.DeleteCatHandler)

	//
	r.POST("/missions", missionHandler.CreateMissionHandler)
	r.GET("/missions", missionHandler.ListMissionsHandler)
	r.GET("/missions/:id", missionHandler.GetMissionByIDHandler)
	r.PUT("/missions/:id", missionHandler.UpdateMissionStatusHandler)
	r.DELETE("/missions/:id", missionHandler.DeleteMissionHandler)

	//
	r.POST("/missions/:mission_id/targets", missionHandler.AddTargetToMissionHandler)
	r.PUT("/missions/assign/cat/:cat_id", missionHandler.AssignCatToMissionHandler)
	r.PUT("/targets/:target_id/status", missionHandler.UpdateTargetStatusHandler)
	r.PUT("/targets/:target_id/notes", missionHandler.UpdateTargetNotesHandler)
	r.DELETE("/targets/:target_id", missionHandler.DeleteTargetHandler)
}
