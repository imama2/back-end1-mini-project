package superadmin

import (
	"back-end1-mini-project/entities"
	"back-end1-mini-project/helper"
	"back-end1-mini-project/repositories"
	"time"
)

type UsecaseSuperadminInterface interface {
	CreateSuperadmin(superadmin SuperAdminParam) (*entities.Account, error)
	LoginSuperadmin(id uint, username, password string) (*entities.Account, string, error)
	CreateCustomer(user CustomerParam) (entities.Customer, error)
	DeleteCustomerByID(id uint) error
	GetAllCustomers(firstName, lastName, email string, page, pageSize int) ([]*entities.Customer, error)
	ApprovedAdminRegister(id uint) error
	RejectedAdminRegister(id uint) error
	UpdateActivedAdmin(id uint) error
	UpdateDeadactivedAdmin(id uint) error
	GetApprovalRequest() ([]*entities.Customer, error)
	GetAllAdmins(username string, page, pageSize int) ([]*entities.Customer, error)
}

type UsecaseSuperadmin struct {
	SuperadminRepo repositories.SuperAdminRepositoryInterface
}

func (uc UsecaseSuperadmin) CreateSuperadmin(superadmin SuperAdminParam) (*entities.Account, error) {
	hasPass := helpers.HashPass(superadmin.Password)

	newSuperadmin := &entities.Account{
		Username:  superadmin.Username,
		Password:  hasPass,
		RoleId:    1,
		Verified:  true,
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createSuperadmin, err := uc.SuperadminRepo.CreateSuperAdmin(newSuperadmin)
	if err != nil {
		return nil, err
	}

	return createSuperadmin, nil
}

func (uc UsecaseSuperadmin) LoginSuperadmin(id uint, username, password string) (*entities.Account, string, error) {
	superadmin, err := uc.SuperadminRepo.LoginSuperAdmin(username)
	if err != nil {
		return nil, "", err
	}

	// Verify hashed password
	comparePass := helpers.ComparePass([]byte(superadmin.Password), []byte(password))
	if !comparePass {
		return nil, "", err
	}

	// Generate JWT token
	tokenString, err := helpers.GenerateToken(id, username)
	if err != nil {
		return nil, "", err
	}

	return superadmin, tokenString, nil
}

func (uc UsecaseSuperadmin) CreateCustomer(user CustomerParam) (entities.Customer, error) {
	newCustomer := &entities.Customer{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Avatar:    user.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := uc.SuperadminRepo.CreateCustomer(newCustomer)
	if err != nil {
		return *newCustomer, err
	}

	return *newCustomer, nil
}

func (uc UsecaseSuperadmin) DeleteCustomerByID(id uint) error {
	// Get existing user data
	existingData, err := uc.SuperadminRepo.GetCustomerById(id)
	if err != nil {
		return err
	}

	return uc.SuperadminRepo.DeleteCustomerById(id, existingData)
}

func (uc UsecaseSuperadmin) GetAllCustomers(firstName, lastName, email string, page, pageSize int) ([]*entities.Customer, error) {
	users, err := uc.SuperadminRepo.GetAllCustomers(firstName, lastName, email, page, pageSize)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc UsecaseSuperadmin) ApprovedAdminRegister(id uint) error {
	err := uc.SuperadminRepo.ApproveAdminRegistration(id)
	if err != nil {
		return err
	}

	return nil
}

func (uc UsecaseSuperadmin) RejectedAdminRegister(id uint) error {
	err := uc.SuperadminRepo.RejectAdminRegistration(id)
	if err != nil {
		return err
	}

	return nil
}

func (uc UsecaseSuperadmin) UpdateActivedAdmin(id uint, isActive bool) error {
	err := uc.SuperadminRepo.UpdateAdminActiveStatus(id, isActive)
	if err != nil {
		return err
	}

	return nil
}

func (uc UsecaseSuperadmin) UpdateDeadactivedAdmin(id uint) error {
	err := uc.SuperadminRepo.UpdateDeadactivedAdmin(id)
	if err != nil {
		return err
	}

	return nil
}

func (uc UsecaseSuperadmin) GetApprovalRequest() ([]*entities.Account, error) {
	accounts, err := uc.SuperadminRepo.GetApprovalRequests()
	if err != nil {
		return nil, err
	}

	updatedAccounts := make([]*entities.Account, len(accounts))
	for i, account := range accounts {
		updatedAccounts[i] = &entities.Account{
			Id:        account.Id,
			Username:  account.Username,
			RoleId:    account.RoleId,
			Verified:  account.Verified,
			CreatedAt: account.CreatedAt,
			UpdatedAt: account.UpdatedAt,
		}
	}

	return updatedAccounts, nil
}

func (uc UsecaseSuperadmin) GetAllAdmins(username string, page, pageSize int) ([]*entities.Account, error) {
	admins, err := uc.SuperadminRepo.GetAllAdmins(username, page, pageSize)
	if err != nil {
		return nil, err
	}

	return admins, nil
}
