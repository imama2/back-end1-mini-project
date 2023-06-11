package unit

import (
	"back-end1-mini-project/entities"
	"back-end1-mini-project/mocks"
	"back-end1-mini-project/modules/customer"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateCustomer(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockCustomerRepositoryInterface(mockCtrl)

	usecase := customer.NewCustomerUseCase(mockRepo)

	customerParam := customer.CustomerParam{
		FirstName: "John",
		LastName:  "Bro",
		Email:     "johnbro@jomblo.com",
		Avatar:    "avatar.jpg",
	}
	mockRepo.EXPECT().CreateCustomer(gomock.Any()).Return(&entities.Customer{
		Id:        1,
		FirstName: customerParam.FirstName,
		LastName:  customerParam.LastName,
		Email:     customerParam.Email,
		Avatar:    customerParam.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil)

	createdCustomer, err := usecase.CreateCustomer(customerParam)

	assert.NoError(t, err)
	assert.NotNil(t, createdCustomer)
	assert.Equal(t, uint(1), createdCustomer.Id)
	assert.Equal(t, "John", createdCustomer.FirstName)
	assert.Equal(t, "Bro", createdCustomer.LastName)
	assert.Equal(t, "johnbro@jomblo.com", createdCustomer.Email)
	assert.Equal(t, "avatar.jpg", createdCustomer.Avatar)
}

func TestGetCustomerById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockCustomerRepositoryInterface(mockCtrl)
	usecase := customer.NewCustomerUseCase(mockRepo)

	expectedCustomer := &entities.Customer{
		Id:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Avatar:    "avatar.jpg",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.EXPECT().GetCustomerById(expectedCustomer.Id).Return(expectedCustomer, nil)

	user, err := usecase.GetCustomerById(expectedCustomer.Id)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedCustomer.Id, user.Id)
	assert.Equal(t, expectedCustomer.FirstName, user.FirstName)
	assert.Equal(t, expectedCustomer.LastName, user.LastName)
	assert.Equal(t, expectedCustomer.Email, user.Email)
	assert.Equal(t, expectedCustomer.Avatar, user.Avatar)
}

func TestUpdateCustomerById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockCustomerRepositoryInterface(mockCtrl)
	usecase := customer.NewCustomerUseCase(mockRepo)

	customerParam := customer.CustomerParam{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Avatar:    "avatar.jpg",
	}

	existingCustomer := &entities.Customer{
		Id:        1,
		FirstName: "Existing",
		LastName:  "Customer",
		Email:     "existing@example.com",
		Avatar:    "existing-avatar.jpg",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.EXPECT().GetCustomerById(existingCustomer.Id).Return(existingCustomer, nil)
	mockRepo.EXPECT().UpdateCustomerById(existingCustomer.Id, gomock.Any()).Return(existingCustomer, nil)

	updatedCustomer, err := usecase.UpdateCustomerById(existingCustomer.Id, customerParam)

	assert.NoError(t, err)
	assert.NotNil(t, updatedCustomer)
	assert.Equal(t, existingCustomer.Id, updatedCustomer.Id)
	assert.Equal(t, customerParam.FirstName, updatedCustomer.FirstName)
	assert.Equal(t, customerParam.LastName, updatedCustomer.LastName)
	assert.Equal(t, customerParam.Email, updatedCustomer.Email)
	assert.Equal(t, customerParam.Avatar, updatedCustomer.Avatar)
}

func TestDeleteCustomerById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockCustomerRepositoryInterface(mockCtrl)
	usecase := customer.NewCustomerUseCase(mockRepo)

	userID := uint(1)

	mockRepo.EXPECT().DeleteCustomerById(userID).Return(nil)

	err := usecase.DeleteCustomerById(userID)
	assert.NoError(t, err)
}
