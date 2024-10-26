package memberships

import (
	"testing"

	"github.com/AthThobari/simple_music_catalog_go/internal/configs"
	"github.com/AthThobari/simple_music_catalog_go/internal/models/memberships"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_service_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)

	type args struct {
		request memberships.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "failed when get user",
			args: args{
				request: memberships.LoginRequest{
					Email: "test@gmail.com",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed when password not match",
			args: args{
				request: memberships.LoginRequest{
					Email: "test@gmail.com",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email: "test@gmail.com",
					Password: "wrong password",
					Username: "testusername",
				},nil)
			},
		},
		{
			name: "success",
			args: args{
				request: memberships.LoginRequest{
					Email: "test@gmail.com",
					Password: "password",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email: "test@gmail.com",
					Password: "$2a$10$5McRU/oDfoQdFXXeT8Zly.6MbPg/nyro4Fep/YpxW3xHKLFYQ5D02",
					Username: "agaton",
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				cfg: &configs.Config{
					Service: configs.Service{
						SecretKey: "abc",
					},
				},
				repository: mockRepo,
			}
			got, err := s.Login(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				assert.NotEmpty(t, got)
			} else {
				assert.Empty(t, got)
			}
		})
	}
}
