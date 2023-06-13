package unit

import (
	"back-end1-mini-project/entities"
	"back-end1-mini-project/mocks"
	"back-end1-mini-project/modules/superadmin"
	"back-end1-mini-project/repositories"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

//func TestUsecaseSuperadmin_CreateCustomer(t *testing.T) {
//	type fields struct {
//		superadminRepo repositories.SuperAdminRepositoryInterface
//	}
//	type args struct {
//		customer superadmin.CustomerParam
//	}
//
//	ctrl := gomock.NewController(t)
//	mockRepo := mocks.NewMockSuperAdminRepositoryInterface(ctrl)
//
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    entities.Customer
//		wantErr bool
//	}{
//		{
//			name:   "Test Case 1",
//			fields: fields{superadminRepo: mockRepo},
//			args: args{customer: superadmin.CustomerParam{
//				FirstName: "dias",
//				LastName:  "pangestu",
//				Email:     "dias@gmail.com",
//				Avatar:    "dias.jpg",
//			}},
//			want: entities.Customer{
//				FirstName: "dias",
//				LastName:  "pangestu",
//				Email:     "dias@gmail.com",
//				Avatar:    "dias.jpg",
//			},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			uc := superadmin.UsecaseSuperadmin{
//				SuperadminRepo: tt.fields.superadminRepo,
//			}
//
//			mockRepo.EXPECT().CreateCustomer(gomock.Any()).Return(&tt.want, nil)
//
//			got, err := uc.CreateCustomer(tt.args.customer)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("CreateCustomer() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestUsecaseSuperadmin_CreateSuperadmin(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepositoryInterface
	}
	type args struct {
		superadmin superadmin.SuperAdminParam
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockSuperAdminRepositoryInterface(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Account
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Test Case 1",
			fields: fields{superadminRepo: mockRepo},
			args: args{superadmin: superadmin.SuperAdminParam{
				Username: "diaspangestu",
				Password: "123",
			}},
			want: &entities.Account{
				Username: "diaspangestu",
				Password: "123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := superadmin.UsecaseSuperadmin{
				SuperadminRepo: tt.fields.superadminRepo,
			}

			mockRepo.EXPECT().CreateSuperAdmin(gomock.Any()).Return(tt.want, nil)

			got, err := uc.CreateSuperadmin(tt.args.superadmin)
			got.Password = "123"
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSuperadmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSuperadmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseSuperadmin_DeleteCustomerByID(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepositoryInterface
	}
	type args struct {
		id uint
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockSuperAdminRepositoryInterface(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "asdasd",
			fields:  fields{superadminRepo: mockRepo},
			args:    args{id: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := superadmin.UsecaseSuperadmin{
				SuperadminRepo: tt.fields.superadminRepo,
			}
			mockRepo.EXPECT().GetCustomerById(gomock.Any()).Return(&entities.Customer{}, nil)
			mockRepo.EXPECT().DeleteCustomerById(gomock.Any(), gomock.Any()).Return(nil).Times(1)

			if err := uc.DeleteCustomerByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCustomerById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecaseSuperadmin_GetAllAdmins(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepositoryInterface
	}
	type args struct {
		username string
		page     int
		pageSize int
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockSuperAdminRepositoryInterface(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entities.Account
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Test Case 1",
			fields: fields{superadminRepo: mockRepo},
			args: args{
				username: "dias",
				page:     2,
				pageSize: 6,
			},
			want:    []*entities.Account{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := superadmin.UsecaseSuperadmin{
				SuperadminRepo: tt.fields.superadminRepo,
			}

			mockRepo.EXPECT().GetAllAdmins(tt.args.username, tt.args.page, tt.args.pageSize).Return(tt.want, nil)

			got, err := uc.GetAllAdmins(tt.args.username, tt.args.page, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllAdmins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllAdmins() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseSuperadmin_GetAllCustomers(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepositoryInterface
	}
	type args struct {
		first_name string
		last_name  string
		email      string
		page       int
		pageSize   int
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockSuperAdminRepositoryInterface(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entities.Customer
		wantErr bool
	}{
		{
			name:   "Test Case 1",
			fields: fields{superadminRepo: mockRepo},
			args: args{
				first_name: "dias",
				last_name:  "pangestu",
				email:      "dias@gmail.com",
				page:       2,
				pageSize:   6,
			},
			want:    []*entities.Customer{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := superadmin.UsecaseSuperadmin{
				SuperadminRepo: tt.fields.superadminRepo,
			}

			mockRepo.EXPECT().GetAllCustomers(tt.args.first_name, tt.args.last_name, tt.args.email, tt.args.page, tt.args.pageSize).Return(tt.want, nil)

			got, err := uc.GetAllCustomers(tt.args.first_name, tt.args.last_name, tt.args.email, tt.args.page, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllCustomers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllCustomers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseSuperadmin_GetApprovalRequest(t *testing.T) {
	type fields struct {
		SuperadminRepo repositories.SuperAdminRepositoryInterface
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*entities.Account
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := superadmin.UsecaseSuperadmin{
				SuperadminRepo: tt.fields.SuperadminRepo,
			}
			got, err := uc.GetApprovalRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApprovalRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetApprovalRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseSuperadmin_LoginSuperadmin(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepositoryInterface
	}
	type args struct {
		id       uint
		username string
		password string
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockSuperAdminRepositoryInterface(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := superadmin.UsecaseSuperadmin{
				SuperadminRepo: tt.fields.superadminRepo,
			}

			mockRepo.EXPECT().LoginSuperAdmin(tt.args.username).Return(tt.want, nil)

			got, _, err := uc.LoginSuperadmin(tt.args.id, tt.args.username, tt.args.password)

			if (err != nil) != tt.wantErr {
				t.Errorf("LoginSuperadmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginSuperadmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseSuperadmin_RejectedAdminRegister(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepositoryInterface
	}
	type args struct {
		id uint
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockSuperAdminRepositoryInterface(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test Case 1",
			fields:  fields{superadminRepo: mockRepo},
			args:    args{id: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := superadmin.UsecaseSuperadmin{
				SuperadminRepo: tt.fields.superadminRepo,
			}

			mockRepo.EXPECT().RejectAdminRegistration(tt.args.id).Return(nil)

			if err := uc.RejectedAdminRegister(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("RejectedAdminRegister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//func TestUsecaseSuperadmin_UpdateActivedAdmin(t *testing.T) {
//	type fields struct {
//		superadminRepo repositories.SuperAdminRepositoryInterface
//	}
//	type args struct {
//		id uint
//	}
//
//	ctrl := gomock.NewController(t)
//	mockRepo := mocks.NewMockSuperAdminRepositoryInterface(ctrl)
//
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		{
//			name:    "Test Case 1",
//			fields:  fields{superadminRepo: mockRepo},
//			args:    args{id: 1},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			uc := superadmin.UsecaseSuperadmin{
//				SuperadminRepo: tt.fields.superadminRepo,
//			}
//
//			mockRepo.EXPECT().UpdateAdminActiveStatus(tt.args.id, nil).Return(nil)
//
//			if err := uc.UpdateActivedAdmin(tt.args.id, true); (err != nil) != tt.wantErr {
//				t.Errorf("UpdateActivedAdmin() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

func TestUsecaseSuperadmin_UpdateDeadactivedAdmin(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepositoryInterface
	}
	type args struct {
		id uint
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockSuperAdminRepositoryInterface(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			fields:  fields{superadminRepo: mockRepo},
			args:    args{id: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := superadmin.UsecaseSuperadmin{
				SuperadminRepo: tt.fields.superadminRepo,
			}

			mockRepo.EXPECT().UpdateDeadactivedAdmin(tt.args.id).Return(nil)

			if err := uc.UpdateDeadactivedAdmin(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateDeadactivedAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
