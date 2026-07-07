package usecase

import (
	"github.com/Vincent088/btg-backend-test/go-backend/internal/entity"
	"github.com/Vincent088/btg-backend-test/go-backend/internal/repository"
)

type CustomerUsecase struct {
	customerRepo    *repository.CustomerRepository
	familyRepo      *repository.FamilyListRepository
	nationalityRepo *repository.NationalityRepository
}

func NewCustomerUsecase(
	customerRepo *repository.CustomerRepository,
	familyRepo *repository.FamilyListRepository,
	nationalityRepo *repository.NationalityRepository,
) *CustomerUsecase {
	return &CustomerUsecase{
		customerRepo:    customerRepo,
		familyRepo:      familyRepo,
		nationalityRepo: nationalityRepo,
	}
}

func (u *CustomerUsecase) CreateCustomerWithFamily(c entity.Customer) (int, error) {
	tx, err := u.customerRepo.BeginTx()
	if err != nil {
		return 0, err
	}

	cstID, err := u.customerRepo.Create(tx, c)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, f := range c.Family {
		f.CstID = cstID
		if err := u.familyRepo.Create(tx, f); err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return cstID, nil
}

func (u *CustomerUsecase) GetCustomerByID(id int) (*entity.Customer, error) {
	c, err := u.customerRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	nat, err := u.nationalityRepo.GetByID(c.NationalityID)
	if err == nil {
		c.Nationality = nat
	}

	family, err := u.familyRepo.GetByCustomerID(c.CstID)
	if err == nil {
		c.Family = family
	}

	return c, nil
}

func (u *CustomerUsecase) GetAllCustomers() ([]entity.Customer, error) {
	customers, err := u.customerRepo.GetAll()
	if err != nil {
		return nil, err
	}

	for i := range customers {
		nat, err := u.nationalityRepo.GetByID(customers[i].NationalityID)
		if err == nil {
			customers[i].Nationality = nat
		}
		family, err := u.familyRepo.GetByCustomerID(customers[i].CstID)
		if err == nil {
			customers[i].Family = family
		}
	}

	return customers, nil
}

func (u *CustomerUsecase) UpdateCustomerWithFamily(c entity.Customer) error {
	if err := u.customerRepo.Update(c); err != nil {
		return err
	}

	tx, err := u.customerRepo.BeginTx()
	if err != nil {
		return err
	}

	if err := u.familyRepo.DeleteByCustomerID(tx, c.CstID); err != nil {
		tx.Rollback()
		return err
	}

	for _, f := range c.Family {
		f.CstID = c.CstID
		if err := u.familyRepo.Create(tx, f); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (u *CustomerUsecase) DeleteCustomer(id int) error {
	return u.customerRepo.Delete(id)
}
