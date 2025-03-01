package services

import (
	"devTodTestTask/internal/models"
	"devTodTestTask/internal/repo"
)

type CatService struct {
	Repo *repo.CatRepository
}

func (s *CatService) CreateCat(cat *models.Cat) error {
	return s.Repo.CreateCat(cat)
}

func (s *CatService) ListCats() ([]models.Cat, error) {
	return s.Repo.ListCats()
}

func (s *CatService) CatByID(id uint) (*models.Cat, error) {
	return s.Repo.GetCatByID(id)
}

func (s *CatService) UpdateCat(cat *models.Cat) error {
	return s.Repo.UpdateCat(cat)
}

func (s *CatService) DeleteCat(id uint) error {
	return s.Repo.DeleteCat(id)
}
