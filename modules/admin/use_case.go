package admin

import (
	"back-end1-mini-project/entities"
	helpers "back-end1-mini-project/helper"
	"back-end1-mini-project/repositories"
	"time"
)

type AdminUseCaseInterface interface {
	LoginAdmin(username, password string) (entities.Account, error)
	RegisterAdmin(admin AdminParam) (entities.Account, error)
	DeleteCustomerByID(id uint) error
	CreateCustomer(user *entities.Customer) (*entities.Customer, error)
	GetAllCustomers(firstName, lastName, email string, page, pageSize int) ([]*entities.Customer, error)
	SaveCustomersFromAPI() error
}

type AdminUseCase struct {
	adminRepo repositories.AdminRepository
	//AdminRepo *mocks.MockAdminRepositoryInterface
}

func (uc AdminUseCase) LoginAdmin(username, password string) (*entities.Account, string, error) {
	admin, err := uc.adminRepo.LoginAdmin(username)
	if err != nil {
		return nil, "", err
	}

	// Verify hashed password
	comparePass := helpers.ComparePass([]byte(admin.Password), []byte(password))
	if !comparePass {
		return nil, "", err
	}

	// Generate JWT token
	tokenString, err := helpers.GenerateToken(admin.Id, username)
	if err != nil {
		return nil, "", err
	}

	return admin, tokenString, nil
}

func (uc AdminUseCase) RegisterAdmin(admin AdminParam) (*entities.Account, error) {
	hashedPassword := helpers.HashPass(admin.Password)

	newAdmin := &entities.Account{
		Username:  admin.Username,
		Password:  hashedPassword,
		RoleId:    2,
		Verified:  false,
		Active:    false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdAdmin, err := uc.adminRepo.RegisterAdmin(newAdmin)
	if err != nil {
		return nil, err
	}

	return createdAdmin, nil
}

func (uc AdminUseCase) CreateCustomer(user *entities.Customer) (entities.Customer, error) {
	newCustomer := &entities.Customer{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Avatar:    user.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdCustomer, err := uc.adminRepo.CreateCustomer(newCustomer)
	if err != nil {
		return *newCustomer, err
	}

	return *createdCustomer, nil
}

func (uc AdminUseCase) DeleteCustomerByID(id uint) error {
	// Get Existing Customer Data
	existingData, err := uc.adminRepo.GetCustomerById(id)
	if err != nil {
		return err
	}

	return uc.adminRepo.DeleteCustomerById(id, existingData)
}

func (uc AdminUseCase) GetAllCustomers(firstName, lastName, email string, page, pageSize int) ([]*entities.Customer, error) {
	users, err := uc.adminRepo.GetAllCustomers(firstName, lastName, email, page, pageSize)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc AdminUseCase) SaveCustomersFromAPI() error {
	url := "https://reqres.in/api/users?page=2"

	err := uc.adminRepo.SaveCustomersFromAPI(url)
	if err != nil {
		return err
	}

	return nil
}
