package repositories

import (
	"back-end1-mini-project/entities"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"io"
	"net/http"
	"time"
)

type AdminRepositoryInterface interface {
	// Admin
	CreateApproval(adminId uint) (*entities.Approval, error)
	GetApprovalById(id uint) (*entities.Approval, error)
	LoginAdmin(username string) (*entities.Account, error)
	RegisterAdmin(admin *entities.Account) (*entities.Account, error)

	// API Integration
	FetchCustomersFromAPI() ([]*entities.Customer, error)
	SaveCustomersFromAPI(url string) error

	// Customer
	CreateCustomer(user *entities.Customer) (*entities.Customer, error)
	DeleteCustomerById(id uint, user *entities.Customer) error
	GetAllCustomers(firstName, lastName, email string, page, pageSize int) ([]*entities.Customer, error)
	GetCustomerByEmail(email string) (*entities.Customer, error)
	GetCustomerById(id uint) (*entities.Customer, error)
}

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return AdminRepository{
		db: db,
	}
}

func (repo AdminRepository) LoginAdmin(username string) (*entities.Account, error) {
	admin := &entities.Account{}

	err := repo.db.Model(&entities.Account{}).Where("username = ? AND verified = ? AND active = ?", username, true, true).First(admin).Error
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (repo AdminRepository) RegisterAdmin(admin *entities.Account) (*entities.Account, error) {
	err := repo.db.Model(&entities.Account{}).Create(admin).Error
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (repo AdminRepository) CreateCustomer(user *entities.Customer) (*entities.Customer, error) {
	err := repo.db.Model(&entities.Customer{}).Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo AdminRepository) GetCustomerById(id uint) (*entities.Customer, error) {
	user := &entities.Customer{}

	err := repo.db.Model(&entities.Customer{}).Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo AdminRepository) DeleteCustomerById(id uint, user *entities.Customer) error {
	err := repo.db.Model(&entities.Customer{}).Where("id = ?", id).Delete(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo AdminRepository) GetAllCustomers(firstName, lastName, email string, page, pageSize int) ([]*entities.Customer, error) {
	var users []*entities.Customer

	query := repo.db.Model(&entities.Customer{})

	if firstName != "" {
		query = query.Where("first_name LIKE ?", "%"+firstName+"%")
	} else if lastName != "" {
		query = query.Where("last_name LIKE ?", "%"+lastName+"%")
	} else if email != "" {
		query = query.Where("email LIKE ?", "%"+email+"%")
	}

	offset := (page - 1) * pageSize

	err := query.Offset(offset).Limit(pageSize).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repo AdminRepository) GetCustomerByEmail(email string) (*entities.Customer, error) {
	user := &entities.Customer{}

	err := repo.db.Model(&entities.Customer{}).Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

type GetCustomersResponse struct {
	Customers []*entities.Customer `json:"data"`
}

func (repo AdminRepository) SaveCustomersFromAPI(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			// Tangani error penutupan body di sini
		}
	}()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	userAPIResponse := new(GetCustomersResponse)

	err = json.Unmarshal(body, userAPIResponse)
	if err != nil {
		return err
	}

	for _, user := range userAPIResponse.Customers {
		_, err := repo.GetCustomerByEmail(user.Email)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				newCustomer := &entities.Customer{
					FirstName: user.FirstName,
					LastName:  user.LastName,
					Email:     user.Email,
					Avatar:    user.Avatar,
				}
				_, err = repo.CreateCustomer(newCustomer)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			// Customer already exists, skip saving
		}
	}

	return nil
}

func (repo AdminRepository) CreateApproval(adminId uint) (*entities.Approval, error) {
	approval := &entities.Approval{
		AdminId:      adminId,
		SuperAdminId: 1,
		Status:       false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := repo.db.Model(&entities.Approval{}).Create(approval).Error
	if err != nil {
		return nil, err
	}

	return approval, nil
}

func (repo AdminRepository) GetApprovalById(id uint) (*entities.Approval, error) {
	approval := &entities.Approval{}

	err := repo.db.Model(&entities.Approval{}).Where("id = ?", id).First(approval).Error
	if err != nil {
		return nil, err
	}

	return approval, nil
}
