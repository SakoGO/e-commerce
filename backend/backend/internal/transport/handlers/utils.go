package handlers

type Validator interface {
	ValidateStruct(interface{}) error
}
