package repositories

import (
	"back-end1-mini-project/entities"
	"gorm.io/gorm"
)

type SuperAdminRepository interface {
	// Admin
	CreateSuperAdmin(superAdmin *entities.Account) (*entities.Account, error)
	ApproveAdminRegistration(id uint) error
	GetAllAdmins(username string, page, pageSize int) ([]*entities.Account, error)
	LoginSuperAdmin(username string) (*entities.Account, error)
	RejectAdminRegistration(id uint) error
	UpdateAdminActiveStatus(id uint, isActive bool) error
	UpdateDeadactivedAdmin(id uint) error

	// Customer
	CreateCustomer(user *entities.Customer) (*entities.Customer, error)
	DeleteCustomerById(id uint, user *entities.Customer) error
	GetAllCustomers(firstName, lastName, email string, page, pageSize int) ([]*entities.Customer, error)
	GetCustomerById(id uint) (*entities.Customer, error)

	// Approval
	GetApprovalRequests() ([]*entities.Account, error)
}

type SuperAdmin struct {
	db *gorm.DB
}

func NewSuperAdmin(db *gorm.DB) SuperAdmin {
	return SuperAdmin{
		db: db,
	}
}

func (repo SuperAdmin) CreateSuperAdmin(superAdmin *entities.Account) (*entities.Account, error) {
	err := repo.db.Create(superAdmin).Error
	if err != nil {
		return nil, err
	}

	return superAdmin, nil
}

func (repo SuperAdmin) LoginSuperAdmin(username string) (*entities.Account, error) {
	superAdmin := &entities.Account{}

	err := repo.db.Where("username = ?", username).First(superAdmin).Error
	if err != nil {
		return nil, err
	}

	return superAdmin, nil
}

func (repo SuperAdmin) CreateCustomer(user *entities.Customer) (*entities.Customer, error) {
	err := repo.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo SuperAdmin) GetCustomerById(id uint) (*entities.Customer, error) {
	user := &entities.Customer{}

	err := repo.db.Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo SuperAdmin) DeleteCustomerById(id uint) error {
	return repo.db.Delete(&entities.Customer{}, id).Error
}

func (repo SuperAdmin) GetAllCustomers(firstName, lastName, email string, page, pageSize int) ([]*entities.Customer, error) {
	var users []*entities.Customer

	query := repo.db.Model(&entities.Customer{})

	if firstName != "" {
		query = query.Where("first_name LIKE ?", "%"+firstName+"%")
	} else if lastName != "" {
		query = query.Where("last_name LIKE ?", "%"+lastName+"%")
	} else if email != "" {
		query = query.Where("email LIKE ?", "%"+email+"%")
	}

	// Pagination
	offset := (page - 1) * pageSize

	err := query.Offset(offset).Limit(pageSize).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repo SuperAdmin) ApproveAdminRegistration(id uint) error {
	admin := &entities.Account{}

	err := repo.db.Where("id = ?", id).First(admin).Error
	if err != nil {
		return err
	}

	admin.Verified = true
	return repo.db.Save(admin).Error
}

func (repo SuperAdmin) RejectAdminRegistration(id uint) error {
	admin := &entities.Account{}

	err := repo.db.Where("id = ?", id).First(admin).Error
	if err != nil {
		return err
	}

	return repo.db.Delete(admin).Error
}

func (repo SuperAdmin) UpdateAdminActiveStatus(id uint, isActive bool) error {
	admin := &entities.Account{}

	err := repo.db.Where("id = ?", id).First(admin).Error
	if err != nil {
		return err
	}

	admin.Active = isActive
	return repo.db.Save(admin).Error
}

func (repo SuperAdmin) GetApprovalRequests() ([]*entities.Account, error) {
	var admins []*entities.Account

	err := repo.db.Model(&entities.Account{}).
		Select("id, username, role_id, is_verified, created_at, updated_at").
		Where("role_id = ? AND is_verified = ?", 2, false).
		Find(&admins).Error
	if err != nil {
		return nil, err
	}

	return admins, nil
}

func (repo SuperAdmin) GetAllAdmins(username string, page, pageSize int) ([]*entities.Account, error) {
	var admins []*entities.Account

	query := repo.db.Model(&entities.Account{})

	if username != "" {
		query = query.Where("username LIKE ?", "%"+username+"%")
	}

	query.Where("role_id = ? AND is_verified = ?", 2, true)

	// Pagination
	offset := (page - 1) * pageSize

	err := query.Offset(offset).Limit(pageSize).Find(&admins).Error
	if err != nil {
		return nil, err
	}

	return admins, nil

}

func (repo SuperAdmin) UpdateDeadactivedAdmin(id uint) error {
	// Check admin data by id
	err := repo.db.Where("id = ?", id).First(&entities.Account{}).Error
	if err != nil {
		return err
	}

	admin := &entities.Account{}
	err = repo.db.First(admin, id).Error
	if err != nil {
		return err
	}

	// Update Active
	if admin.Active == true {
		err = repo.db.Model(&entities.Account{}).Where("id = ?", id).Update("active", false).Error
		if err != nil {
			return err
		}
	}

	return nil
}
