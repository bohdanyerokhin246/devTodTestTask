package services

import (
	"devTodTestTask/internal/models"
	"devTodTestTask/internal/repo"
	"errors"
)

type MissionService struct {
	Repo *repo.MissionRepository
}

func (s *MissionService) CreateMission(mission *models.Mission) error {
	if mission.IsComplete {
		return errors.New("mission cannot be created as completed")
	}

	return s.Repo.CreateMission(mission)
}

func (s *MissionService) ListMissions() ([]models.Mission, error) {
	return s.Repo.ListMissions()
}

func (s *MissionService) GetMissionByID(id uint) (*models.Mission, error) {
	return s.Repo.GetMissionByID(id)
}

func (s *MissionService) UpdateMissionStatus(mission *models.Mission) error {
	if mission.IsComplete {
		return errors.New("cannot update mission after completion")
	}
	return s.Repo.UpdateMissionStatus(mission)
}

func (s *MissionService) DeleteMission(id uint) error {
	return s.Repo.DeleteMission(id)
}

func (s *MissionService) AddTargetToMission(missionID uint, target *models.Target) error {
	return s.Repo.AddTargetToMission(missionID, target)
}

func (s *MissionService) AssignCatToMission(missionID, catID uint) error {
	return s.Repo.AssignCatToMission(missionID, catID)
}

func (s *MissionService) UpdateTargetStatus(target *models.Target) error {
	return s.Repo.UpdateTargetStatus(target)
}

func (s *MissionService) UpdateTargetNotes(target *models.Target) error {
	return s.Repo.UpdateTargetNotes(target)
}

func (s *MissionService) DeleteTarget(id uint) error {
	return s.Repo.DeleteTarget(id)
}
