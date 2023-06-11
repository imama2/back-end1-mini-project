package repositories

import (
	"back-end1-mini-project/entities"
	"gorm.io/gorm"
)

type CustomerRepositoryInterface interface {
	// Create
	CreateCustomer(customer *entities.Customer) (*entities.Customer, error)

	// Retrieve
	GetCustomerById(id uint) (*entities.Customer, error)

	// Update
	UpdateCustomerById(id uint, customer *entities.Customer) (*entities.Customer, error)

	// Delete
	DeleteCustomerById(id uint) error
}

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return CustomerRepository{
		db: db,
	}
}

func (repo CustomerRepository) CreateCustomer(customer *entities.Customer) (*entities.Customer, error) {
	err := repo.db.Model(&entities.Customer{}).Create(customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (repo CustomerRepository) GetCustomerById(id uint) (*entities.Customer, error) {
	customer := &entities.Customer{}

	err := repo.db.Model(&entities.Customer{}).Where("id = ?", id).First(customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (repo CustomerRepository) UpdateCustomerById(id uint, customer *entities.Customer) (*entities.Customer, error) {
	err := repo.db.Model(&entities.Customer{}).Where("id = ?", id).Save(customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (repo CustomerRepository) DeleteCustomerById(id uint) error {
	err := repo.db.Model(&entities.Customer{}).Where("id = ?", id).Delete(&entities.Customer{}).Error
	if err != nil {
		return err
	}

	return nil
}
