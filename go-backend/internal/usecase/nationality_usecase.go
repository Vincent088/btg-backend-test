package usecase

import (
	"github.com/Vincent088/btg-backend-test/go-backend/internal/entity"
	"github.com/Vincent088/btg-backend-test/go-backend/internal/repository"
)

type NationalityUsecase struct {
	repo *repository.NationalityRepository
}

func NewNationalityUsecase(repo *repository.NationalityRepository) *NationalityUsecase {
	return &NationalityUsecase{repo: repo}
}

func (u *NationalityUsecase) GetAllNationalities() ([]entity.Nationality, error) {
	list, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []entity.Nationality{}
	}
	return list, nil
}
