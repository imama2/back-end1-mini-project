package customer

import (
	"back-end1-mini-project/entities"
	"back-end1-mini-project/repositories"
	"time"
)

type CustomerUseCaseInterface interface {
	CreateCustomer(customer CustomerParam) (entities.Customer, error)
	GetCustomerById(id uint) (entities.Customer, error)
	UpdateCustomerById(id uint, user CustomerParam) (entities.Customer, error)
	DeleteCustomerById(id uint) error
}

type CustomerUseCase struct {
	userRepo repositories.CustomerRepositoryInterface
}

func NewCustomerUseCase(userRepo repositories.CustomerRepositoryInterface) CustomerUseCase {
	return CustomerUseCase{
		userRepo: userRepo,
	}
}

func (uc CustomerUseCase) CreateCustomer(user CustomerParam) (entities.Customer, error) {
	newCustomer := entities.Customer{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Avatar:    user.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdCustomer, err := uc.userRepo.CreateCustomer(&newCustomer)
	if err != nil {
		return newCustomer, err
	}

	return *createdCustomer, nil
}

func (uc CustomerUseCase) GetCustomerById(id uint) (entities.Customer, error) {
	user, err := uc.userRepo.GetCustomerById(id)
	if err != nil {
		return entities.Customer{}, err
	}

	return *user, nil
}

func (uc CustomerUseCase) UpdateCustomerById(id uint, user CustomerParam) (entities.Customer, error) {
	existingCustomer, err := uc.userRepo.GetCustomerById(id)
	if err != nil {
		return entities.Customer{}, err
	}

	existingCustomer.FirstName = user.FirstName
	existingCustomer.LastName = user.LastName
	existingCustomer.Email = user.Email
	existingCustomer.Avatar = user.Avatar
	existingCustomer.UpdatedAt = time.Now()

	updatedCustomer, err := uc.userRepo.UpdateCustomerById(id, existingCustomer)
	if err != nil {
		return entities.Customer{}, err
	}

	return *updatedCustomer, nil
}

func (uc CustomerUseCase) DeleteCustomerById(id uint) error {
	err := uc.userRepo.DeleteCustomerById(id)
	if err != nil {
		return err
	}

	return nil
}
