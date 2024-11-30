package validator

import (
	"e-commerce/backend/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

// User validation tests
func TestUserValidation(t *testing.T) {
	validator := NewGoValidator()

	// test: valid data
	validUser := model.User{
		Username: "Ruslan",
		Email:    "Ruslan@gmail.com",
		Password: "Ruslanches",
		Phone:    "89183219212",
	}

	err := validator.ValidateStruct(validUser)
	assert.Nil(t, err, "Expected no error for valid user")

	// test: invalid email
	invalidEmail := model.User{
		Username: "Ruslan",
		Email:    "Ruslangmail.com", // Invalid email format
		Password: "Ruslanches",
		Phone:    "89183219212",
	}

	err = validator.ValidateStruct(invalidEmail)
	assert.NotNil(t, err, "Expected error for invalid email")

	// test: invalid password
	invalidPassword := model.User{
		Username: "Ruslan",
		Email:    "Ruslang@mail.com",
		Password: "Rusik", // Invalid password: less than 6 characters
		Phone:    "89183219212",
	}

	err = validator.ValidateStruct(invalidPassword)
	assert.NotNil(t, err, "Expected error for invalid password")

	// test: invalid phone symbols
	invalidPhoneSymbol := model.User{
		Username: "Ruslan",
		Email:    "Ruslan@gmail.com",
		Password: "Ruslanches",
		Phone:    "89183s219212", // Invalid phone number with symbols
	}

	err = validator.ValidateStruct(invalidPhoneSymbol)
	assert.NotNil(t, err, "Expected error for invalid phone (contains symbols)")

	// test: invalid phone
	invalidPhone := model.User{
		Username: "Ruslan",
		Email:    "Ruslan@gmail.com",
		Password: "Ruslanches",
		Phone:    "8918321921", // Invalid phone number: too short
	}

	err = validator.ValidateStruct(invalidPhone)
	assert.NotNil(t, err, "Expected error for invalid phone (incorrect number length)")

	// test: empty email
	emptyEmail := model.User{
		Username: "Ruslan",
		Email:    "", // Empty email
		Password: "Ruslanches",
		Phone:    "89183219212",
	}

	err = validator.ValidateStruct(emptyEmail)
	assert.NotNil(t, err, "Expected error for empty email")

	// test: empty password
	emptyPass := model.User{
		Username: "Ruslan",
		Email:    "Ruslan@gmail.com",
		Password: "", // Empty password
		Phone:    "89183219212",
	}

	err = validator.ValidateStruct(emptyPass)
	assert.NotNil(t, err, "Expected error for empty password")

	// test: empty phone
	emptyPhone := model.User{
		Username: "Ruslan",
		Email:    "Ruslan@gmail.com",
		Password: "Ruslanches",
		Phone:    "", // Empty phone
	}

	err = validator.ValidateStruct(emptyPhone)
	assert.NotNil(t, err, "Expected error for empty phone")
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
