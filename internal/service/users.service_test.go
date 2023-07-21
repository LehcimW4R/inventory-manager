package service

import (
	"context"
	"os"
	"testing"

	"github.com/LehcimW4R/inventory-manager/internal/models"
	"github.com/LehcimW4R/inventory-manager/internal/repository"
	"github.com/LehcimW4R/inventory-manager/utils"
	"github.com/stretchr/testify/mock"
)

var repo *repository.MockRepository
var s Service

func TestMain(m *testing.M) {
	validPassword, _ := utils.Encrypt([]byte("validPassword"))
	encryptedPassword := utils.ToBase64(validPassword)
	u := &models.User{Email: "test@exists.com", Password: encryptedPassword}

	repo = &repository.MockRepository{}
	repo.On("GetUserByEmail", mock.Anything, "test@test.com").Return(nil, nil)
	repo.On("GetUserByEmail", mock.Anything, "test@exists.com").Return(u, nil)
	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

	repo.On("GetUserRoles", mock.Anything, int64(1)).Return([]models.UserRole{{UserID: 1, RoleID: 1}}, nil)
	repo.On("SaveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("RemoveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	s = New(repo)

	code := m.Run()
	os.Exit(code)
}

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		CaseName      string
		Email         string
		UserName      string
		Password      string
		ExpectedError error
	}{
		{
			CaseName:      "RegisterUser_Success",
			Email:         "test@test.com",
			UserName:      "test",
			Password:      "validPassword",
			ExpectedError: nil,
		},
		{
			CaseName:      "RegisterUser_UserAlreadyExists",
			Email:         "test@exists.com",
			UserName:      "test",
			Password:      "validPassword",
			ExpectedError: ErrUserAlreadyExists,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.CaseName, func(t *testing.T) {
			t.Parallel()

			repo.Mock.Test(t)

			err := s.RegisterUser(ctx, tc.Email, tc.UserName, tc.Password)

			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	testCases := []struct {
		CaseName      string
		Email         string
		Password      string
		ExpectedError error
	}{
		{
			CaseName:      "LoginUser_Success",
			Email:         "test@exists.com",
			Password:      "validPassword",
			ExpectedError: nil,
		},
		{
			CaseName:      "LoginUser_InvalidPassword",
			Email:         "test@exists.com",
			Password:      "invalidPassword",
			ExpectedError: ErrInvalidCredentials,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.CaseName, func(t *testing.T) {
			t.Parallel()

			repo.Mock.Test(t)

			_, err := s.LoginUser(ctx, tc.Email, tc.Password)

			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestAddUserRole(t *testing.T) {
	testCases := []struct {
		CaseName      string
		UserID        int64
		RoleID        int64
		ExpectedError error
	}{
		{
			CaseName:      "AddUserRole_Success",
			UserID:        1,
			RoleID:        2,
			ExpectedError: nil,
		},
		{
			CaseName:      "AddUserRole_UserAlreadyHasRole",
			UserID:        1,
			RoleID:        1,
			ExpectedError: ErrRoleAlreadyAdded,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.CaseName, func(t *testing.T) {
			t.Parallel()

			repo.Mock.Test(t)

			err := s.AddUserRole(ctx, tc.UserID, tc.RoleID)

			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestRemoveUserRole(t *testing.T) {
	testCases := []struct {
		CaseName      string
		UserID        int64
		RoleID        int64
		ExpectedError error
	}{
		{
			CaseName:      "RemoveUserRole_Success",
			UserID:        1,
			RoleID:        1,
			ExpectedError: nil,
		},
		{
			CaseName:      "RemoveUserRole_UserDoesNotHaveRole",
			UserID:        1,
			RoleID:        3,
			ExpectedError: ErrRoleNotFound,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.CaseName, func(t *testing.T) {
			t.Parallel()

			repo.Mock.Test(t)

			err := s.RemoveUserRole(ctx, tc.UserID, tc.RoleID)

			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
