package customer

import "back-end1-mini-project/dto"

type CustomerInterfaceController struct {
}
type CustomerController struct {
	uc CustomerUseCaseInterface
}

func NewCustomerController(uc CustomerUseCaseInterface) CustomerController {
	return CustomerController{
		uc: uc,
	}
}

func (ctrl CustomerController) CreateCustomer(req CustomerParam) (interface{}, error) {
	user, err := ctrl.uc.CreateCustomer(req)
	if err != nil {
		return dto.ErrorResponse{}, err
	}

	response := SuccessCreateCustomer{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Create Customer",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: CustomerParam{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Avatar:    user.Avatar,
		},
	}

	return response, nil
}

func (ctrl CustomerController) GetCustomerById(id uint) (interface{}, error) {
	user, err := ctrl.uc.GetCustomerById(id)
	if err != nil {
		return dto.ErrorResponse{}, err
	}

	response := SuccessCreateCustomer{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Get Customer",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: CustomerParam{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Avatar:    user.Avatar,
		},
	}

	return response, nil
}

func (ctrl CustomerController) UpdateCustomerById(id uint, user CustomerParam) (interface{}, error) {
	updatedCustomer, err := ctrl.uc.UpdateCustomerById(id, user)
	if err != nil {
		return dto.ErrorResponse{}, err
	}

	response := SuccessUpdateCustomer{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Update Customer",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: CustomerParam{
			FirstName: updatedCustomer.FirstName,
			LastName:  updatedCustomer.LastName,
			Email:     updatedCustomer.Email,
			Avatar:    updatedCustomer.Avatar,
		},
	}

	return response, nil
}

func (ctrl CustomerController) DeleteCustomerById(id uint) error {
	err := ctrl.uc.DeleteCustomerById(id)
	if err != nil {
		return err
	}

	return nil
}
