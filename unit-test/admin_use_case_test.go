package unit

import (
	"back-end1-mini-project/entities"
	"back-end1-mini-project/mocks"
	"back-end1-mini-project/modules/admin"
	"back-end1-mini-project/modules/customer"
	"back-end1-mini-project/repositories"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestAdminUseCase_CreateCustomer(t *testing.T) {
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

func TestAdminUseCase_DeleteCustomerByID(t *testing.T) {
	type fields struct {
		adminRepo repositories.AdminRepositoryInterface
	}
	type args struct {
		id uint
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockAdminRepositoryInterface(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test Case 1",
			fields:  fields{adminRepo: mockRepo},
			args:    args{id: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := admin.AdminUseCase{
				AdminRepos: tt.fields.adminRepo,
			}

			mockRepo.EXPECT().GetCustomerById(gomock.Any()).Return(&entities.Customer{}, nil)

			mockRepo.EXPECT().DeleteCustomerById(gomock.Any(), gomock.Any()).Return(nil).Times(1)

			if err := uc.DeleteCustomerByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCustomerById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdminUseCase_GetAllCustomers(t *testing.T) {
	type fields struct {
		adminRepo repositories.AdminRepositoryInterface
	}
	type args struct {
		first_name string
		last_name  string
		email      string
		page       int
		pageSize   int
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockAdminRepositoryInterface(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entities.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Test Case 1",
			fields: fields{adminRepo: mockRepo},
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
			uc := admin.AdminUseCase{
				AdminRepos: tt.fields.adminRepo,
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

func TestAdminUseCase_LoginAdmin(t *testing.T) {
	type fields struct {
		adminRepo repositories.AdminRepositoryInterface
	}
	type args struct {
		id       uint
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Account
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := admin.AdminUseCase{
				AdminRepos: tt.fields.adminRepo,
			}
			got, got1, err := uc.LoginAdmin(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginAdmin() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LoginAdmin() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAdminUseCase_RegisterAdmin(t *testing.T) {
	type fields struct {
		adminRepo repositories.AdminRepositoryInterface
	}
	type args struct {
		admin admin.AdminParam
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockAdminRepositoryInterface(ctrl)

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
			fields: fields{adminRepo: mockRepo},
			args: args{admin: admin.AdminParam{
				ID:       1,
				Username: "dias",
				Password: "dias123",
			}},
			want: &entities.Account{
				Username: "dias",
				Password: "dias123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := admin.AdminUseCase{
				AdminRepos: tt.fields.adminRepo,
			}

			mockRepo.EXPECT().RegisterAdmin(gomock.Any()).Return(tt.want, nil)

			got, err := uc.RegisterAdmin(tt.args.admin)
			got.Password = "dias123"
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterAdmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdminUseCase_SaveCustomersFromAPI(t *testing.T) {
	type fields struct {
		adminRepo repositories.AdminRepositoryInterface
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockAdminRepositoryInterface(ctrl)

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			fields:  fields{adminRepo: mockRepo},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := admin.AdminUseCase{
				AdminRepos: tt.fields.adminRepo,
			}

			mockRepo.EXPECT().SaveCustomersFromAPI(gomock.Any()).Return(nil)

			if err := uc.SaveCustomersFromAPI(); (err != nil) != tt.wantErr {
				t.Errorf("SaveCustomersFromAPI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
