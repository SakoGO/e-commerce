package validator

import (
	"e-commerce/backend/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

// User validation tests using switch-case structure
func TestUserValidation(t *testing.T) {
	validator := NewGoValidator()

	// Test cases with different types of validation issues
	tests := []struct {
		name     string
		user     model.User
		expected string
	}{
		{
			name: "Valid user",
			user: model.User{
				Username: "Ruslan",
				Email:    "Ruslan@gmail.com",
				Password: "Ruslanches",
				Phone:    "89183219212",
			},
			expected: "", // No error expected
		},
		{
			name: "Invalid email",
			user: model.User{
				Username: "Ruslan",
				Email:    "Ruslangmail.com", // Invalid email format
				Password: "Ruslanches",
				Phone:    "89183219212",
			},
			expected: "email: Ruslangmail.com does not validate as email",
		},
		{
			name: "Invalid password (too short)",
			user: model.User{
				Username: "Ruslan",
				Email:    "Ruslang@mail.com",
				Password: "Rusik", // Invalid password: less than 6 characters
				Phone:    "89183219212",
			},
			expected: "password: Rusik does not validate as length(6|20)",
		},
		{
			name: "Invalid phone (contains symbols)",
			user: model.User{
				Username: "Ruslan",
				Email:    "Ruslan@gmail.com",
				Password: "Ruslanches",
				Phone:    "89183s219212", // Invalid phone number with symbols
			},
			expected: "phone: 89183s219212 does not validate as numeric",
		},
		{
			name: "Invalid phone (too short)",
			user: model.User{
				Username: "Ruslan",
				Email:    "Ruslan@gmail.com",
				Password: "Ruslanches",
				Phone:    "8918321921", // Invalid phone number: too short
			},
			expected: "phone: 8918321921 does not validate as matches(^[0-9]{11}$)",
		},
		{
			name: "Empty email",
			user: model.User{
				Username: "Ruslan",
				Email:    "", // Empty email
				Password: "Ruslanches",
				Phone:    "89183219212",
			},
			expected: "email: non zero value required",
		},
		{
			name: "Empty password",
			user: model.User{
				Username: "Ruslan",
				Email:    "Ruslan@gmail.com",
				Password: "", // Empty password
				Phone:    "89183219212",
			},
			expected: "password: non zero value required",
		},
		{
			name: "Empty phone",
			user: model.User{
				Username: "Ruslan",
				Email:    "Ruslan@gmail.com",
				Password: "Ruslanches",
				Phone:    "", // Empty phone
			},
			expected: "phone: non zero value required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateStruct(tt.user)

			// Use switch-case for handling different validation cases
			switch {
			case tt.expected == "":
				// If no error is expected, assert nil error
				assert.Nil(t, err, "Expected no error for valid user")
			default:
				// If an error is expected, assert the error message contains the expected string
				assert.NotNil(t, err, "Expected error but got nil")
				assert.Contains(t, err.Error(), tt.expected, "Unexpected error message")
			}
		})
	}
}

/*
func TestAddressValidation(t *testing.T) {
	validator := NewGoValidator()

	// test: valid data
	validAddress := model.Address{
		AddressName: "Kolotushkina",
		HouseNum:    229,
		Street:      "Pushkina",
		City:        "Krasnodar",
		Region:      "Krasnodarskiy Kray",
		Phone:       "89183219212",
	}

	err := validator.ValidateStruct(validAddress)
	assert.Nil(t, err, "Expected no error for valid address")

	// test: empty AddressName
	emptyAddressName := model.Address{
		AddressName: "", // Empty AddressName
		HouseNum:    229,
		Street:      "Pushkina",
		City:        "Krasnodar",
		Region:      "Krasnodarskiy Kray",
		Phone:       "89183219212",
	}

	err = validator.ValidateStruct(emptyAddressName)
	assert.NotNil(t, err, "Expected error for empty AddressName")

	// test: empty HouseNum
	emptyHouseNum := model.Address{
		AddressName: "Kolotushkina",
		HouseNum:    0, // Invalid HouseNum (0)
		Street:      "Pushkina",
		City:        "Krasnodar",
		Region:      "Krasnodarskiy Kray",
		Phone:       "89183219212",
	}

	err = validator.ValidateStruct(emptyHouseNum)
	assert.NotNil(t, err, "Expected error for empty HouseNum")

	// test: invalid Phone (non-numeric)
	invalidPhoneSymbols := model.Address{
		AddressName: "Kolotushkina",
		HouseNum:    229,
		Street:      "Pushkina",
		City:        "Krasnodar",
		Region:      "Krasnodarskiy Kray",
		Phone:       "89183s219212", // Invalid phone number with symbols
	}

	err = validator.ValidateStruct(invalidPhoneSymbols)
	assert.NotNil(t, err, "Expected error for invalid phone (non-numeric)")

	// test: invalid Phone (incorrect length)
	invalidPhoneLength := model.Address{
		AddressName: "Kolotushkina",
		HouseNum:    229,
		Street:      "Pushkina",
		City:        "Krasnodar",
		Region:      "Krasnodarskiy Kray",
		Phone:       "8918321929", // Invalid phone number (length)
	}

	err = validator.ValidateStruct(invalidPhoneLength)
	assert.NotNil(t, err, "Expected error for invalid phone (incorrect length)")

	// test: empty Phone
	emptyPhone := model.Address{
		AddressName: "Kolotushkina",
		HouseNum:    229,
		Street:      "Pushkina",
		City:        "Krasnodar",
		Region:      "Krasnodarskiy Kray",
		Phone:       "", // Empty phone number
	}

	err = validator.ValidateStruct(emptyPhone)
	assert.NotNil(t, err, "Expected error for empty phone")

	// test: empty Street
	emptyStreet := model.Address{
		AddressName: "Kolotushkina",
		HouseNum:    229,
		Street:      "", // Empty street name
		City:        "Krasnodar",
		Region:      "Krasnodarskiy Kray",
		Phone:       "89183219212",
	}

	err = validator.ValidateStruct(emptyStreet)
	assert.NotNil(t, err, "Expected error for empty street")

	// test: empty City
	emptyCity := model.Address{
		AddressName: "Kolotushkina",
		HouseNum:    229,
		Street:      "Pushkina",
		City:        "", // Empty city name
		Region:      "Krasnodarskiy Kray",
		Phone:       "89183219212",
	}

	err = validator.ValidateStruct(emptyCity)
	assert.NotNil(t, err, "Expected error for empty city")

	// test: empty Region
	emptyRegion := model.Address{
		AddressName: "Kolotushkina",
		HouseNum:    229,
		Street:      "Pushkina",
		City:        "Krasnodar",
		Region:      "", // Empty region name
		Phone:       "89183219212",
	}

	err = validator.ValidateStruct(emptyRegion)
	assert.NotNil(t, err, "Expected error for empty region")
}
*/
