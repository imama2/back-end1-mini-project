package superadmin

import (
	"back-end1-mini-project/dto"
)

type ControllerSuperadminInterface interface {
	CreateSuperadmin(req SuperAdminParam) (interface{}, error)
	LoginSuperadmin(id uint, username, password string) (interface{}, error)
	CreateUser(req CustomerParam) (interface{}, error)
	DeleteUserById(id uint) error
	GetAllUsers(firstName, lastName, email string, page, pageSize int) (interface{}, error)
	ApproveAdminRegistration(id uint) (interface{}, error)
	RejectAdminRegistration(id uint) (interface{}, error)
	UpdateAdminActiveStatus(id uint) (interface{}, error)
	UpdateAdminInactiveStatus(id uint) (interface{}, error)
	GetApprovalRequests() (interface{}, error)
	GetAllAdmins(username string, page, pageSize int) (interface{}, error)
}

type ControllerSuperadmin struct {
	uc UsecaseSuperadminInterface
}

func (ctrl ControllerSuperadmin) CreateSuperadmin(req SuperAdminParam) (interface{}, error) {
	_, err := ctrl.uc.CreateSuperadmin(req)
	if err != nil {
		return SuccessCreateSuperadmin{}, err
	}

	response := SuccessCreateSuperadmin{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Create Superadmin",
			Message:      "Success",
			ResponseTime: "",
		},
	}

	return response, nil
}

func (ctrl ControllerSuperadmin) LoginSuperadmin(id uint, username, password string) (interface{}, error) {
	superAdmin, tokenString, err := ctrl.uc.LoginSuperadmin(id, username, password)
	if err != nil {
		return nil, err
	}

	response := SuccessLoginSuperadmin{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Login Successfull",
			Message:      "Success",
			ResponseTime: "",
		},
		Username: superAdmin.Username,
		Token:    tokenString,
	}

	return response, nil
}

// CreateUser Superadmin
func (ctrl ControllerSuperadmin) CreateUser(req CustomerParam) (interface{}, error) {
	customer, err := ctrl.uc.CreateCustomer(req)
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
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
		},
	}

	return response, nil
}

// DeleteUserById Superadmin
func (ctrl ControllerSuperadmin) DeleteUserById(id uint) error {
	err := ctrl.uc.DeleteCustomerByID(id)
	if err != nil {
		return err
	}

	return nil
}

// GetAllUsers Superadmin
func (ctrl ControllerSuperadmin) GetAllUsers(firstName, lastName, email string, page, pageSize int) (interface{}, error) {
	users, err := ctrl.uc.GetAllCustomers(firstName, lastName, email, page, pageSize)
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
		Data: users,
	}

	return response, nil
}

func (ctrl ControllerSuperadmin) ApproveAdminRegistration(id uint) (interface{}, error) {
	err := ctrl.uc.ApprovedAdminRegister(id)
	if err != nil {
		return SuccessUpdateRegisterAdmin{}, err
	}

	response := SuccessUpdateRegisterAdmin{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Approved Registration Admin",
			Message:      "Approved",
			ResponseTime: "",
		},
	}

	return response, nil
}

func (ctrl ControllerSuperadmin) RejectAdminRegistration(id uint) (interface{}, error) {
	err := ctrl.uc.RejectedAdminRegister(id)
	if err != nil {
		return SuccessUpdateRegisterAdmin{}, err
	}

	response := SuccessUpdateRegisterAdmin{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Rejected Registration Admin",
			Message:      "Rejected",
			ResponseTime: "",
		},
	}

	return response, nil
}

func (ctrl ControllerSuperadmin) UpdateAdminActiveStatus(id uint) (interface{}, error) {
	err := ctrl.uc.UpdateActivedAdmin(id)
	if err != nil {
		return SuccessUpdateRegisterAdmin{}, err
	}

	response := SuccessUpdateRegisterAdmin{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Admin is Activated Now",
			Message:      "Success",
			ResponseTime: "",
		},
	}

	return response, nil
}

func (ctrl ControllerSuperadmin) UpdateAdminInactiveStatus(id uint) (interface{}, error) {
	err := ctrl.uc.UpdateDeadactivedAdmin(id)
	if err != nil {
		return SuccessUpdateRegisterAdmin{}, err
	}

	response := SuccessUpdateRegisterAdmin{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Admin is Deactivated Now",
			Message:      "Success",
			ResponseTime: "",
		},
	}

	return response, nil
}

func (ctrl ControllerSuperadmin) GetApprovalRequests() (interface{}, error) {
	requests, err := ctrl.uc.GetApprovalRequest()
	if err != nil {
		return SuccessGetApprovalRequest{}, err
	}

	response := SuccessGetApprovalRequest{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Get All Approval Requests Data",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: requests,
	}

	return response, nil
}

func (ctrl ControllerSuperadmin) GetAllAdmins(username string, page, pageSize int) (interface{}, error) {
	admins, err := ctrl.uc.GetAllAdmins(username, page, pageSize)
	if err != nil {
		return SuccessGetAllAdmins{}, err
	}

	response := SuccessGetAllAdmins{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Get All Admin Data",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: admins,
	}

	return response, nil
}
