package customer

import "back-end1-mini-project/dto"

type CustomerParam struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

// SuccessCreateCustomer represents the success response for user creation
type SuccessCreateCustomer struct {
	dto.Response
	Data CustomerParam `json:"data"`
}

// SuccessUpdateCustomer represents the success response for user update
type SuccessUpdateCustomer struct {
	dto.Response
	Data CustomerParam `json:"data"`
}
