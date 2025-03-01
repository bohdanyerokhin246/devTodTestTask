package handlers

import (
	"devTodTestTask/internal/models"
	"devTodTestTask/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MissionHandler struct {
	Service *services.MissionService
}

// CreateMissionHandler godoc
// @Summary Create a new mission
// @Description Create a new mission and store it in the database
// @Param mission body models.Mission true "Mission data"
// @Success 201 {object} models.Mission "Successfully created mission"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /missions [post]
func (h *MissionHandler) CreateMissionHandler(c *gin.Context) {
	var mission models.Mission
	if err := c.ShouldBindJSON(&mission); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if err := h.Service.CreateMission(&mission); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, mission)
}

// ListMissionsHandler godoc
// @Summary Get list of all missions
// @Description Get a list of all missions in the database
// @Success 200 {array} models.Mission "List of missions"
// @Failure 404 {object} ErrorResponse "Missions not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /missions [get]
func (h *MissionHandler) ListMissionsHandler(c *gin.Context) {
	missions, err := h.Service.ListMissions()
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Missions not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": missions})
}

// GetMissionByIDHandler godoc
// @Summary Get mission by ID
// @Description Get a specific mission by its ID
// @Param id path int true "Mission ID"
// @Success 200 {object} models.Mission "Mission data"
// @Failure 404 {object} ErrorResponse "Mission not found"
// @Failure 400 {object} ErrorResponse "Invalid ID format"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /missions/{id} [get]
func (h *MissionHandler) GetMissionByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	mission, err := h.Service.GetMissionByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Mission not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": mission})
}

// UpdateMissionStatusHandler godoc
// @Summary Update mission status
// @Description Update the status of an existing mission.
// @Param mission body models.Mission true "Updated mission data"
// @Success 200 {object} ErrorResponse "Successfully updated mission status"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /missions [put]
func (h *MissionHandler) UpdateMissionStatusHandler(c *gin.Context) {
	var mission models.Mission
	if err := c.ShouldBindJSON(&mission); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	err := h.Service.UpdateMissionStatus(&mission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Mission status updated successfully"})
}

// DeleteMissionHandler godoc
// @Summary Delete a mission
// @Description Delete a mission by its ID
// @Param id path int true "Mission ID"
// @Success 200 {object} ErrorResponse "Successfully deleted mission"
// @Failure 400 {object} ErrorResponse "Invalid ID format"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /missions/{id} [delete]
func (h *MissionHandler) DeleteMissionHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.Service.DeleteMission(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted mission"})
}

// AddTargetToMissionHandler godoc
// @Summary Add a target to a mission
// @Description Add a target to a specific mission
// @Param mission_id path int true "Mission ID"
// @Param target body models.Target true "Target data"
// @Success 201 {object} models.Target "Successfully added target to mission"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /missions/{mission_id}/targets [post]
func (h *MissionHandler) AddTargetToMissionHandler(c *gin.Context) {
	var target models.Target
	if err := c.ShouldBindJSON(&target); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	missionID, _ := strconv.Atoi(c.Param("mission_id"))
	if err := h.Service.AddTargetToMission(uint(missionID), &target); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, target)
}

// AssignCatToMissionHandler godoc
// @Summary Assign a cat to a mission
// @Description Assign a cat to a specific mission
// @Param mission_id path int true "Mission ID"
// @Param cat_id path int true "Cat ID"
// @Success 200 {object} ErrorResponse "Successfully assigned cat to mission"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /missions/{mission_id}/cats/{cat_id} [put]
func (h *MissionHandler) AssignCatToMissionHandler(c *gin.Context) {
	missionID, err := strconv.Atoi(c.Param("mission_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid mission_id"})
		return
	}

	catID, err := strconv.Atoi(c.Param("cat_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid cat_id"})
		return
	}

	if err := h.Service.AssignCatToMission(uint(missionID), uint(catID)); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cat assigned to mission successfully"})
}

// UpdateTargetStatusHandler godoc
// @Summary Update target status
// @Description Update the status of a specific target
// @Param target body models.Target true "Target status data"
// @Success 200 {object} ErrorResponse "Target status updated successfully"
// @Failure 400 {object} ErrorResponse "Invalid ID format"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /targets/status [put]
func (h *MissionHandler) UpdateTargetStatusHandler(c *gin.Context) {
	var target models.Target
	if err := c.ShouldBindJSON(&target); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if err := h.Service.UpdateTargetStatus(&target); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Target status updated successfully"})
}

// UpdateTargetNotesHandler godoc
// @Summary Update target notes
// @Description Update notes for a specific target
// @Param target body models.Target true "Target notes data"
// @Success 200 {object} ErrorResponse "Target notes updated successfully"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /targets/notes [put]
func (h *MissionHandler) UpdateTargetNotesHandler(c *gin.Context) {
	var target models.Target

	if err := c.ShouldBindJSON(&target); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if err := h.Service.UpdateTargetNotes(&target); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Target notes updated successfully"})
}

// DeleteTargetHandler godoc
// @Summary Delete a target
// @Description Delete a target by its ID
// @Param id path int true "Target ID"
// @Success 204 {object} ErrorResponse "Successfully deleted target"
// @Failure 400 {object} ErrorResponse "Invalid ID format"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /targets/{id} [delete]
func (h *MissionHandler) DeleteTargetHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteTarget(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "Successfully deleted target"})
}
