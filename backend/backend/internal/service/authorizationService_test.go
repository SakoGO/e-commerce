package service

import (
	"e-commerce/backend/internal/model"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthRepository struct {
	mock.Mock
}

func (m *MockAuthRepository) SignUP(user *model.User) error {
	return m.Called(user).Error(0)
}
func (m *MockAuthRepository) SignIN(user *model.User) error {
	return m.Called(user).Error(0)
}
func (m *MockAuthRepository) FindByUsername(username string) (*model.User, error) {
	if user := m.Called(username).Get(0); user != nil {
		return user.(*model.User), m.Called(username).Error(1)
	}
	return nil, m.Called(username).Error(1)
}

func (m *MockAuthRepository) FindByEmail(email string) (*model.User, error) {
	if user := m.Called(email).Get(0); user != nil {
		return user.(*model.User), m.Called(email).Error(1)
	}
	return nil, m.Called(email).Error(1)
}

func (m *MockAuthRepository) FindByPhone(phone string) (*model.User, error) {
	if user := m.Called(phone).Get(0); user != nil {
		return user.(*model.User), m.Called(phone).Error(1)
	}
	return nil, m.Called(phone).Error(1)
}

func TestAuthService_SignUP(t *testing.T) {
	tests := []struct {
		name          string
		username      string
		email         string
		password      string
		phone         string
		expectedError string
		usernameTaken bool
		emailTaken    bool
		phoneTaken    bool
	}{
		{
			name:          "validTest",
			username:      "SAKO",
			email:         "LowScalp@gmail.com",
			password:      "qwerty_12345",
			phone:         "89991112233",
			expectedError: "",
			usernameTaken: false,
			emailTaken:    false,
			phoneTaken:    false,
		},
		{
			name:          "expectedUsername",
			username:      "SAKO_expected",
			email:         "LowScalp@gmail.com",
			password:      "qwerty_12345",
			phone:         "89991112233",
			expectedError: "username SAKO_expected is already taken",
			usernameTaken: true,
			emailTaken:    false,
			phoneTaken:    false,
		},
		{
			name:          "expectedEmail",
			username:      "SAKO",
			email:         "LowScalp_expected@gmail.com",
			password:      "qwerty_12345",
			phone:         "89991112233",
			expectedError: "email LowScalp_expected@gmail.com is already taken",
			usernameTaken: false,
			emailTaken:    true,
			phoneTaken:    false,
		},
		{
			name:          "expectedPhone",
			username:      "SAKO",
			email:         "LowScalp@gmail.com",
			password:      "qwerty_12345",
			phone:         "89997777777",
			expectedError: "phone 89997777777 is already taken",
			usernameTaken: false,
			emailTaken:    false,
			phoneTaken:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockRepo := new(MockAuthRepository)

			if tt.usernameTaken {
				mockRepo.On("FindByUsername", tt.username).Return(&model.User{}, nil)
			} else {
				mockRepo.On("FindByUsername", tt.username).Return(nil, nil)
			}

			if tt.emailTaken {
				mockRepo.On("FindByEmail", tt.email).Return(&model.User{}, nil)
			} else {
				mockRepo.On("FindByEmail", tt.email).Return(nil, nil)
			}

			if tt.phoneTaken {
				mockRepo.On("FindByPhone", tt.phone).Return(&model.User{}, nil)
			} else {
				mockRepo.On("FindByPhone", tt.phone).Return(nil, nil)
			}

			mockRepo.On("SignUP", mock.Anything).Return(nil)

			authService := NewAuthService(mockRepo)

			err := authService.SignUP(tt.username, tt.email, tt.password, tt.phone)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAuthService_SignIN(t *testing.T) {

	os.Setenv("JWT_SECRET_KEY", "test_secret_key")
	defer os.Unsetenv("JWT_SECRET_KEY")

	tests := []struct {
		name              string
		email             string
		password          string
		expectedError     string
		emailNotFound     bool
		passwordIncorrect bool
	}{
		{
			name:              "validTest",
			email:             "SakoScalp@gmail.com",
			password:          "12345qwerty",
			expectedError:     "",
			emailNotFound:     false,
			passwordIncorrect: false,
		},
		{
			name:              "email not found",
			email:             "SakoScalp_not@gmail.com",
			password:          "12345qwerty",
			expectedError:     "account SakoScalp_not@gmail.com is not found",
			emailNotFound:     true,
			passwordIncorrect: false,
		},
		{
			name:              "incorrect password",
			email:             "SakoScalp@gmail.com",
			password:          "12345qwerty_inc",
			expectedError:     "incorrect password",
			emailNotFound:     false,
			passwordIncorrect: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockAuthRepository)

			user := &model.User{
				Email:    tt.email,
				Password: "$2a$10$rf4dahBlDclsKjrBEHUm1eFtnxcoOI1v.ITtL0kCiE9K.ApDg8iEq", // hash for "12345qwerty"
			}

			if tt.emailNotFound {
				mockRepo.On("FindByEmail", tt.email).Return(nil, nil)
			} else {
				mockRepo.On("FindByEmail", tt.email).Return(user, nil)
			}

			authService := NewAuthService(mockRepo)
			token, err := authService.SignIN(tt.email, tt.password)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}
		})
	}
}
