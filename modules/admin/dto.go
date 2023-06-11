package admin

import "back-end1-mini-project/dto"

type AdminParam struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomerParam struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type SuccessCreateAdmin struct {
	dto.Response
}

type SuccessCreateUser struct {
	dto.Response
	Data CustomerParam `json:"data"`
}

type SuccessLoginAdmin struct {
	dto.Response
	Username string `json:"username"`
	Token    string `json:"token"`
}

type SuccessGetAllUsers struct {
	dto.Response
	Data interface{} `json:"data"`
}

type SuccessFetchUsersFromAPI struct {
	dto.Response
}
