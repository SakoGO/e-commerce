package tests

import (
	"bytes"
	"e-commerce/backend/internal/model"
	"e-commerce/backend/internal/transport/handlers"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

type MockAuthService struct {
	mock.Mock
}

type MockValidator struct {
	mock.Mock
}

func (m *MockAuthService) SignUP(username, email, password, phone string) error {
	args := m.Called(username, email, password, phone)
	return args.Error(0)
}

func (m *MockAuthService) SignIN(email, password string) (string, error) {
	args := m.Called(email, password)
	return args.String(0), args.Error(1)
}

func (m *MockAuthService) FindByUsername(username string) (*model.User, error) {
	args := m.Called(username)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockAuthService) FindByEmail(email string) (*model.User, error) {
	args := m.Called(email)
	return args.Get(0).(*model.User), args.Error(1)
}

/*
// TODO: Tests for //UserFindByID

	func (m *MockAuthService) UserFindByID(userID int) (*model.User, error) {
		args := m.Called(userID)
		return args.Get(0).(*model.User), args.Error(1)
	}

// TODO: Tests for //UserDelete

	func (m *MockAuthService) UserDelete(userID int) error {
		args := m.Called(userID)
		return args.Error(0)
	}
*/
func (m *MockValidator) ValidateStruct(i interface{}) error {
	args := m.Called(i)
	return args.Error(0)
}

func TestAuthHandlerSignUP(t *testing.T) {
	tests := []struct {
		name             string
		body             model.User
		validatorError   error
		signUpError      error
		expectedCode     int
		expectedResponse string
	}{
		{
			name: "valid signup",
			body: model.User{
				Username: "SAKO",
				Email:    "Sarkis2292000@gmail.com",
				Password: "qwerty12345",
				Phone:    "89883884330",
			},
			validatorError:   nil,
			signUpError:      nil,
			expectedCode:     http.StatusCreated,
			expectedResponse: `{"-":"User successfully registered"}`,
		},
		{
			name: "validation failed testwwwww",
			body: model.User{
				Username: "",
				Email:    "Sarkis2292002@gmail.com",
				Password: "qwerty12345",
				Phone:    "89883884330",
			},
			validatorError:   assert.AnError,
			signUpError:      nil,
			expectedCode:     http.StatusBadRequest,
			expectedResponse: "incorrect data for registration",
		},
		{
			name: "signup failed test",
			body: model.User{
				Username: "SAK0",
				Email:    "Sarkis2292001@gmail.com",
				Password: "qwerty12345",
				Phone:    "89883884330",
			},
			validatorError:   nil,
			signUpError:      assert.AnError,
			expectedCode:     http.StatusConflict,
			expectedResponse: assert.AnError.Error() + "\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAuthService := new(MockAuthService)
			mockValidator := new(MockValidator)

			mockValidator.On("ValidateStruct", &tt.body).Return(tt.validatorError)
			mockAuthService.On("SignUP", tt.body.Username, tt.body.Email, tt.body.Password, tt.body.Phone).Return(tt.signUpError)

			handler := &handlers.Handler{
				AuthService: mockAuthService,
				Validator:   mockValidator,
			}

			body, _ := json.Marshal(tt.body)
			req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(body))
			w := httptest.NewRecorder()

			handler.SignUP(w, req)

			assert.Equal(t, strings.TrimSpace(strconv.Itoa(tt.expectedCode)), strings.TrimSpace(strconv.Itoa(w.Code)))
			assert.Equal(t, strings.TrimSpace(tt.expectedResponse), strings.TrimSpace(w.Body.String()))
		})
	}
}

func TestAuthHandlerSignIN(t *testing.T) {
	tests := []struct {
		name             string
		body             model.User
		signInError      error
		expectedCode     int
		expectedResponse string
		expectedToken    string
	}{
		{
			name: "valid test",
			body: model.User{
				Email:    "Sarkis229@gmail.com",
				Password: "qwerty12345",
			},
			signInError:      nil,
			expectedCode:     http.StatusAccepted,
			expectedResponse: `{"token":"validToken"}`,
			expectedToken:    "validToken",
		},
		{
			name: "signup failed test",
			body: model.User{
				Email:    "Sarkis229_invalid@gmail.com",
				Password: "wrong_qwerty12345",
			},
			signInError:      assert.AnError,
			expectedCode:     http.StatusBadRequest,
			expectedResponse: "incorrect data format",
			expectedToken:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserService := new(MockAuthService)

			mockUserService.On("SignIN", tt.body.Email, tt.body.Password).Return(tt.expectedToken, tt.signInError)

			handler := &handlers.Handler{
				AuthService: mockUserService,
			}

			body, _ := json.Marshal(tt.body)
			req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(body))
			w := httptest.NewRecorder()

			handler.SignIN(w, req)

			assert.Equal(t, strings.TrimSpace(strconv.Itoa(tt.expectedCode)), strings.TrimSpace(strconv.Itoa(w.Code)))
			assert.Equal(t, strings.TrimSpace(tt.expectedResponse), strings.TrimSpace(w.Body.String()))
		})
	}
}
