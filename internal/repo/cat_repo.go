package repo

import (
	"database/sql"
	"devTodTestTask/internal/models"
	"fmt"
	"time"
)

type CatRepository struct {
	DB *sql.DB
}

func (repo *CatRepository) CreateCat(cat *models.Cat) error {
	query := `INSERT INTO 
				    cats (name, experience, breed, salary,created_at)
              VALUES 
                  ($1, $2, $3, $4, $5) 
              RETURNING id,created_at`

	return repo.DB.QueryRow(query, cat.Name, cat.Experience, cat.Breed, cat.Salary, time.Now()).Scan(&cat.ID, &cat.CreatedAt)
}

func (repo *CatRepository) ListCats() ([]models.Cat, error) {
	var cats []models.Cat
	rows, err := repo.DB.Query(`SELECT 
    									id, name, experience, breed, salary, created_at, updated_at 
									FROM 
									    public.cats 
									WHERE 
									    deleted_at IS NULL`)
	if err != nil {
		return nil, fmt.Errorf("could not list cats: %v", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var cat models.Cat
		err = rows.Scan(&cat.ID, &cat.Name, &cat.Experience, &cat.Breed, &cat.Salary, &cat.CreatedAt, &cat.UpdatedAt)
		if err != nil {
			return nil, err
		}
		cats = append(cats, cat)
	}
	return cats, nil
}

func (repo *CatRepository) GetCatByID(id uint) (*models.Cat, error) {
	var cat models.Cat
	query := `	SELECT 
				    id, name, experience, breed, salary, created_at, updated_at
				FROM 
				    cats 
				WHERE 
				    id = $1 AND deleted_at IS NULL`
	err := repo.DB.QueryRow(query, id).Scan(&cat.ID, &cat.Name, &cat.Experience, &cat.Breed, &cat.Salary, &cat.CreatedAt, &cat.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("could not get cat: %v", err)
	}
	return &cat, nil
}

func (repo *CatRepository) UpdateCat(cat *models.Cat) error {
	query := `	UPDATE 
				    cats 
				SET 
				    salary = $1, updated_at = $2 
				WHERE 
				    id = $3`
	_, err := repo.DB.Exec(query, cat.Salary, time.Now(), cat.ID)
	return err
}

func (repo *CatRepository) DeleteCat(id uint) error {
	query := `	UPDATE 
				    cats 
				SET 
				    deleted_at = $1 
				WHERE 
				    id = $2`
	_, err := repo.DB.Exec(query, time.Now(), id)
	return err
}
