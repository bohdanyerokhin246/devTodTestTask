package repo

import (
	"database/sql"
	"devTodTestTask/internal/models"
	"fmt"
	"time"
)

type MissionRepository struct {
	DB *sql.DB
}

func (repo *MissionRepository) CreateMission(mission *models.Mission) error {
	// Check if the cat already has an active mission
	var activeMissionCount int
	err := repo.DB.QueryRow(`SELECT COUNT(*) FROM missions WHERE cat_id = $1 AND is_complete = FALSE`, mission.CatID).Scan(&activeMissionCount)
	if err != nil {
		return fmt.Errorf("error checking active mission: %v", err)
	}

	// If the cat already has an active mission, prevent creating a new one
	if activeMissionCount > 0 {
		return fmt.Errorf("this cat already has an active mission")
	}

	// Insert a new mission
	query := `INSERT INTO missions (cat_id, is_complete, created_at)
			  VALUES ($1, $2, $3) RETURNING id`
	err = repo.DB.QueryRow(query, mission.CatID, mission.IsComplete, time.Now()).Scan(&mission.ID)
	if err != nil {
		return fmt.Errorf("could not create mission: %v", err)
	}

	// Add targets to the mission
	for _, target := range mission.Targets {
		targetQuery := `INSERT INTO targets (mission_id, name, country, notes, is_complete, created_at)
						VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
		err := repo.DB.QueryRow(targetQuery, mission.ID, target.Name, target.Country, target.Notes, target.IsComplete, time.Now()).Scan(&target.ID)
		if err != nil {
			return fmt.Errorf("could not create target: %v", err)
		}
	}
	return nil
}

func (repo *MissionRepository) ListMissions() ([]models.Mission, error) {
	var missions []models.Mission

	// Query to get all missions
	query := `SELECT id, cat_id, is_complete, created_at, updated_at, deleted_at FROM missions`
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not get missions: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var mission models.Mission
		if err := rows.Scan(
			&mission.ID,
			&mission.CatID,
			&mission.IsComplete,
			&mission.CreatedAt,
			&mission.UpdatedAt,
			&mission.DeletedAt); err != nil {
			return nil, fmt.Errorf("could not scan mission: %v", err)
		}

		// Get all targets for this mission
		targetQuery := `SELECT id, mission_id, name, country, notes, is_complete, created_at, updated_at, deleted_at FROM targets WHERE mission_id = $1`
		targetRows, err := repo.DB.Query(targetQuery, mission.ID)
		if err != nil {
			return nil, fmt.Errorf("could not get targets: %v", err)
		}
		defer targetRows.Close()

		var targets []models.Target
		for targetRows.Next() {
			var target models.Target
			if err := targetRows.Scan(
				&target.ID,
				&target.MissionID,
				&target.Name,
				&target.Country,
				&target.Notes,
				&target.IsComplete,
				&target.CreatedAt,
				&target.UpdatedAt,
				&target.DeletedAt); err != nil {
				return nil, err
			}
			targets = append(targets, target)
		}

		mission.Targets = targets
		missions = append(missions, mission)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return missions, nil
}

func (repo *MissionRepository) GetMissionByID(id uint) (*models.Mission, error) {
	var mission models.Mission
	query := `SELECT id, cat_id, is_complete, created_at, updated_at FROM missions WHERE id = $1`
	err := repo.DB.QueryRow(query, id).Scan(&mission.ID, &mission.CatID, &mission.IsComplete, &mission.CreatedAt, &mission.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("could not get mission: %v", err)
	}

	// Get all targets for this mission
	targetQuery := `SELECT id, mission_id, name, country, notes, is_complete, created_at, updated_at FROM targets WHERE mission_id = $1`
	rows, err := repo.DB.Query(targetQuery, mission.ID)
	if err != nil {
		return nil, fmt.Errorf("could not get targets: %v", err)
	}
	defer rows.Close()

	var targets []models.Target
	for rows.Next() {
		var target models.Target
		if err := rows.Scan(&target.ID, &target.MissionID, &target.Name, &target.Country, &target.Notes, &target.IsComplete, &target.CreatedAt, &target.UpdatedAt); err != nil {
			return nil, err
		}
		targets = append(targets, target)
	}

	mission.Targets = targets
	return &mission, nil
}

func (repo *MissionRepository) UpdateMissionStatus(mission *models.Mission) error {
	// Update mission completion status
	query := `UPDATE missions SET is_complete = $2, updated_at = $3 WHERE id = $4`
	_, err := repo.DB.Exec(query, mission.IsComplete, time.Now(), mission.ID)
	return err
}

func (repo *MissionRepository) DeleteMission(id uint) error {
	var catID uint
	// Get the cat ID associated with the mission
	err := repo.DB.QueryRow(`SELECT cat_id FROM missions WHERE id = $1`, id).Scan(&catID)
	if err != nil {
		return fmt.Errorf("mission not found")
	}

	// Prevent deleting a mission that has a cat assigned
	if catID != 0 {
		return fmt.Errorf("cannot delete mission assigned to a cat")
	}

	// Set deleted_at timestamp to logically delete the mission
	query := `	UPDATE 
				    missions 
				SET 
				    deleted_at = $1 
				WHERE 
				    id = $2`
	_, err = repo.DB.Exec(query, time.Now(), id)
	return err
}

func (repo *MissionRepository) AddTargetToMission(missionID uint, target *models.Target) error {
	// Check if the mission exists and if it is completed
	var isComplete bool
	err := repo.DB.QueryRow(`SELECT is_complete FROM missions WHERE id = $1`, missionID).Scan(&isComplete)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("mission not found")
		}
		return fmt.Errorf("error checking mission: %v", err)
	}

	// Prevent adding targets to completed missions
	if isComplete {
		return fmt.Errorf("cannot add target to a completed mission")
	}

	// Check if there are already 3 targets in the mission
	var targetCount int
	err = repo.DB.QueryRow(`SELECT COUNT(*) FROM targets WHERE mission_id = $1`, missionID).Scan(&targetCount)
	if err != nil {
		return fmt.Errorf("error counting targets: %v", err)
	}

	if targetCount >= 3 {
		return fmt.Errorf("cannot add more than 3 targets to a mission")
	}

	// Insert a new target for the mission
	query := `INSERT INTO targets (mission_id, name, country, notes, is_complete, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err = repo.DB.QueryRow(query, missionID, target.Name, target.Country, target.Notes, target.IsComplete, time.Now(), time.Now()).Scan(&target.ID)
	if err != nil {
		return fmt.Errorf("error inserting target: %v", err)
	}

	return nil
}

func (repo *MissionRepository) AssignCatToMission(missionID, catID uint) error {
	// Check if the mission exists
	var existingMissionID uint
	err := repo.DB.QueryRow(`SELECT id FROM missions WHERE id = $1`, missionID).Scan(&existingMissionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("mission not found")
		}
		return fmt.Errorf("error checking mission: %v", err)
	}

	// Check if the cat already has an active mission
	var activeMissionCount int
	err = repo.DB.QueryRow(`SELECT COUNT(*) FROM missions WHERE cat_id = $1 AND is_complete = FALSE`, catID).Scan(&activeMissionCount)
	if err != nil {
		return fmt.Errorf("error checking active mission: %v", err)
	}

	// Prevent assigning a cat that already has an active mission
	if activeMissionCount > 0 {
		return fmt.Errorf("this cat already has an active mission")
	}

	// Assign the cat to the mission
	query := `UPDATE missions SET cat_id = $1 WHERE id = $2`
	_, err = repo.DB.Exec(query, catID, missionID)
	if err != nil {
		return fmt.Errorf("could not assign cat to mission: %v", err)
	}

	return nil
}

func (repo *MissionRepository) UpdateTargetStatus(target *models.Target) error {
	// Check if the target is complete
	var isTargetComplete bool
	query := `SELECT is_complete FROM targets WHERE id = $1`
	err := repo.DB.QueryRow(query, target.ID).Scan(&isTargetComplete)
	if err != nil {
		return fmt.Errorf("could not find target: %v", err)
	}

	// Prevent updating a completed target
	if isTargetComplete {
		return fmt.Errorf("cannot update: target is complete")
	}

	// Check if the mission of the target is complete
	var missionID uint
	query = `SELECT mission_id FROM targets WHERE id = $1`
	err = repo.DB.QueryRow(query, target.ID).Scan(&missionID)
	if err != nil {
		return fmt.Errorf("could not find mission for target: %v", err)
	}

	var isMissionComplete bool
	query = `SELECT is_complete FROM missions WHERE id = $1`
	err = repo.DB.QueryRow(query, missionID).Scan(&isMissionComplete)
	if err != nil {
		return fmt.Errorf("could not find mission: %v", err)
	}

	// Prevent updating if the mission is complete
	if isMissionComplete {
		return fmt.Errorf("cannot update: mission is complete")
	}

	// Update target status to complete
	query = `UPDATE targets SET is_complete = $1 WHERE id = $2`
	_, err = repo.DB.Exec(query, target.IsComplete, target.ID)
	return err
}

func (repo *MissionRepository) UpdateTargetNotes(target *models.Target) error {
	// Check if the target is complete
	var isTargetComplete bool
	query := `SELECT is_complete FROM targets WHERE id = $1`
	err := repo.DB.QueryRow(query, target.ID).Scan(&isTargetComplete)
	if err != nil {
		return fmt.Errorf("could not find target: %v", err)
	}

	// Prevent updating notes of a completed target
	if isTargetComplete {
		return fmt.Errorf("cannot update notes: target is complete")
	}

	// Check if the mission of the target is complete
	var missionID uint
	query = `SELECT mission_id FROM targets WHERE id = $1`
	err = repo.DB.QueryRow(query, target.ID).Scan(&missionID)
	if err != nil {
		return fmt.Errorf("could not find mission for target: %v", err)
	}

	var isMissionComplete bool
	query = `SELECT is_complete FROM missions WHERE id = $1`
	err = repo.DB.QueryRow(query, missionID).Scan(&isMissionComplete)
	if err != nil {
		return fmt.Errorf("could not find mission: %v", err)
	}

	// Prevent updating notes if the mission is complete
	if isMissionComplete {
		return fmt.Errorf("cannot update notes: mission is complete")
	}

	// Update target notes
	query = `UPDATE targets SET notes = $1 WHERE id = $2`
	_, err = repo.DB.Exec(query, target.Notes, target.ID)
	return err
}

func (repo *MissionRepository) DeleteTarget(id uint) error {
	// Check if the target is complete
	var isComplete bool
	query := `SELECT is_complete FROM targets WHERE id = $1`
	err := repo.DB.QueryRow(query, id).Scan(&isComplete)
	if err != nil {
		return fmt.Errorf("target not found")
	}

	// Prevent deleting a completed target
	if isComplete {
		return fmt.Errorf("cannot delete a completed target")
	}

	// Set deleted_at timestamp to logically delete the target
	query = `	UPDATE 
				    targets 
				SET 
				    deleted_at = $1 
				WHERE 
				    id = $2`
	_, err = repo.DB.Exec(query, time.Now(), id)
	return err
}
