package admin

import (
	"back-end1-mini-project/dto"
	"back-end1-mini-project/entities"
)

type AdminControllerInterface interface {
	LoginAdmin(username, password string) (interface{}, error)
	RegisterAdmin(req AdminParam) (interface{}, error)
	CreateUser(user CustomerParam) (interface{}, error)
	DeleteUserById(id uint) error
	GetAllUsers(first_name, last_name, email string, page, pageSize int) (interface{}, error)
	SaveUsersFromAPI() (interface{}, error)
}

type AdminController struct {
	uc AdminUseCase
}

func (ctrl AdminController) LoginAdmin(id uint, username, password string) (interface{}, error) {
	admin, tokenString, err := ctrl.uc.LoginAdmin(username, password)
	if err != nil {
		return nil, err
	}

	response := SuccessLoginAdmin{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Login Successful",
			Message:      "Success",
			ResponseTime: "",
		},
		Username: admin.Username,
		Token:    tokenString,
	}

	return response, nil
}

func (ctrl AdminController) RegisterAdmin(req AdminParam) (interface{}, error) {
	_, err := ctrl.uc.RegisterAdmin(req)
	if err != nil {
		return SuccessCreateAdmin{}, err
	}

	response := SuccessCreateAdmin{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Register Admin",
			Message:      "Success",
			ResponseTime: "",
		},
	}

	return response, nil
}

// CreateUser Admin
func (ctrl AdminController) CreateUser(req CustomerParam) (interface{}, error) {
	newUser := &entities.Customer{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Avatar:    req.Avatar,
	}

	createdUser, err := ctrl.uc.CreateCustomer(newUser)
	if err != nil {
		return SuccessCreateUser{}, err
	}

	response := SuccessCreateUser{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Create User",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: CustomerParam{
			FirstName: createdUser.FirstName,
			LastName:  createdUser.LastName,
			Email:     createdUser.Email,
			Avatar:    createdUser.Avatar,
		},
	}

	return response, nil
}

// DeleteUserById Admin
func (ctrl AdminController) DeleteUserById(id uint) error {
	err := ctrl.uc.DeleteCustomerByID(id)
	if err != nil {
		return err
	}

	return nil
}

func (ctrl AdminController) GetAllUsers(first_name, last_name, email string, page, pageSize int) (interface{}, error) {
	request, err := ctrl.uc.GetAllCustomers(first_name, last_name, email, page, pageSize)
	if err != nil {
		return SuccessGetAllUsers{}, err
	}

	response := SuccessGetAllUsers{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Get All User Data",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: request,
	}

	return response, nil
}

func (ctrl AdminController) SaveUsersFromAPI() (interface{}, error) {
	err := ctrl.uc.SaveCustomersFromAPI()
	if err != nil {
		return SuccessFetchUsersFromAPI{}, err
	}

	response := SuccessFetchUsersFromAPI{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Fetch User Data from API",
			Message:      "Success",
			ResponseTime: "",
		},
	}

	return response, nil
}
