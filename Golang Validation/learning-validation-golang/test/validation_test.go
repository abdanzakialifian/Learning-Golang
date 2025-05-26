package test

import (
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func TestValidate(t *testing.T) {
	if validate == nil {
		t.Error("Validate is nil")
	}
}

func TestVaribleValidation(t *testing.T) {
	user := "zaki"
	err := validate.Var(user, "required")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestEqualVariableValidation(t *testing.T) {
	password := "secret"
	confirmPassword := "secret"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")

	if err != nil {
		t.Error(err.Error())
	}
}

func TestMultipleTag(t *testing.T) {
	customerNumber := 12345

	err := validate.Var(customerNumber, "required,numeric")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestTagParameter(t *testing.T) {
	customerNumber := "99999"

	err := validate.Var(customerNumber, "required,numeric,min=5,max=10")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestStructValidation(t *testing.T) {
	type LoginRequest struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	loginRequest := LoginRequest{
		Email:    "zaki@gmail.com",
		Password: "zaki23",
	}

	if err := validate.Struct(loginRequest); err != nil {
		t.Error(err.Error())
	}
}

func TestValidationErrors(t *testing.T) {
	type LoginRequest struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	loginRequest := LoginRequest{
		Email:    "zaki",
		Password: "zaki",
	}

	if err := validate.Struct(loginRequest); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			t.Error("Error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestStructCrossFieldValidation(t *testing.T) {
	type RegisterUser struct {
		Email           string `validate:"required,email"`
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}

	request := RegisterUser{
		Email:           "zaki@gmail.com",
		Password:        "secret",
		ConfirmPassword: "secret",
	}

	err := validate.Struct(request)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestNestedStructValidation(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      string  `validate:"required"`
		Name    string  `validate:"required"`
		Address Address `validate:"required"`
	}

	user := User{
		Id:   "1",
		Name: "Zaki",
		Address: Address{
			City:    "Banjarnegara",
			Country: "Indonesia",
		},
	}

	err := validate.Struct(user)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestCollectionValidation(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
	}

	user := User{
		Id:   "1",
		Name: "Zaki",
		Addresses: []Address{
			{
				City:    "Banjarnegara",
				Country: "Indonesia",
			},
			{
				City:    "Tangerang",
				Country: "Indonesia",
			},
			{
				City:    "Jakarta",
				Country: "Indonesia",
			},
		},
	}

	if err := validate.Struct(user); err != nil {
		t.Error(err.Error())
	}
}

func TestBasicCollectionValidation(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
		Hobbies   []string  `validate:"dive,required,min=3"`
	}

	user := User{
		Id:   "1",
		Name: "Zaki",
		Addresses: []Address{
			{
				City:    "Banjarnegara",
				Country: "Indonesia",
			},
			{
				City:    "Tangerang",
				Country: "Indonesia",
			},
			{
				City:    "Jakarta",
				Country: "Indonesia",
			},
		},
		Hobbies: []string{
			"Coding",
			"Gaming",
			"Learning",
		},
	}

	if err := validate.Struct(user); err != nil {
		t.Error(err.Error())
	}
}

func TestMapValidations(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"dive,required,min=3"`
		Schools   map[string]School `validate:"dive,keys,required,min=2,endkeys"`
	}

	user := User{
		Id:   "1",
		Name: "Zaki",
		Addresses: []Address{
			{
				City:    "Banjarnegara",
				Country: "Indonesia",
			},
			{
				City:    "Tangerang",
				Country: "Indonesia",
			},
			{
				City:    "Jakarta",
				Country: "Indonesia",
			},
		},
		Hobbies: []string{
			"Coding",
			"Gaming",
			"Learning",
		},
		Schools: map[string]School{
			"SD": {
				Name: "Sekolah Dasar",
			},
			"SMP": {
				Name: "Sekolah Menengah Pertama",
			},
			"SMA": {
				Name: "Sekolah Menengah Akhir",
			},
		},
	}

	if err := validate.Struct(user); err != nil {
		t.Error(err.Error())
	}
}

func TestBasicMapValidations(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"dive,required,min=3"`
		Schools   map[string]School `validate:"dive,keys,required,min=2,endkeys"`
		Wallet    map[string]int    `validate:"dive,keys,required,endkeys,required,gt=1000"`
	}

	user := User{
		Id:   "1",
		Name: "Zaki",
		Addresses: []Address{
			{
				City:    "Banjarnegara",
				Country: "Indonesia",
			},
			{
				City:    "Tangerang",
				Country: "Indonesia",
			},
			{
				City:    "Jakarta",
				Country: "Indonesia",
			},
		},
		Hobbies: []string{
			"Coding",
			"Gaming",
			"Learning",
		},
		Schools: map[string]School{
			"SD": {
				Name: "Sekolah Dasar",
			},
			"SMP": {
				Name: "Sekolah Menengah Pertama",
			},
			"SMA": {
				Name: "Sekolah Menengah Akhir",
			},
		},
		Wallet: map[string]int{
			"BCA":     1000000,
			"Mandiri": 10000,
			"BNI":     100000,
		},
	}

	if err := validate.Struct(user); err != nil {
		t.Error(err.Error())
	}
}

func TestAliasValidations(t *testing.T) {
	validate.RegisterAlias("varchar", "required,max=255")

	type Seller struct {
		Id     string `validate:"varchar,min=5"`
		Name   string `validate:"varchar"`
		Owner  string `validate:"varchar"`
		Slogan string `validate:"varchar"`
	}

	seller := Seller{
		Id:     "12345",
		Name:   "Zaki",
		Owner:  "Abdan",
		Slogan: "Uhuy",
	}

	if err := validate.Struct(seller); err != nil {
		t.Error(err.Error())
	}
}

func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if !ok {
		return false
	}

	if value != strings.ToUpper(value) {
		return false
	}

	if len(value) < 5 {
		return false
	}

	return true
}

func TestCustomValidationFunction(t *testing.T) {
	validate.RegisterValidation("username", MustValidUsername)

	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	request := LoginRequest{
		Username: "ABDANZAKI",
		Password: "secret",
	}

	if err := validate.Struct(request); err != nil {
		t.Error(err.Error())
	}
}

func MustValidPin(field validator.FieldLevel) bool {
	regex := regexp.MustCompile("^[0-9]+$")
	param, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}

	value, ok := field.Field().Interface().(string)

	if !ok {
		return false
	}

	if !regex.MatchString(value) {
		return false
	}

	return len(value) == param
}

func TestCustomParameterValidations(t *testing.T) {
	validate.RegisterValidation("pin", MustValidPin)

	type LoginRequest struct {
		Phone string `validate:"required,number"`
		PIN   string `validate:"required,pin=6"`
	}

	request := LoginRequest{
		Phone: "085123456789",
		PIN:   "112233",
	}

	if err := validate.Struct(request); err != nil {
		t.Error(err.Error())
	}
}

func TestOrRuleValidations(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email|numeric"`
		Password string `validate:"required"`
	}

	request := LoginRequest{
		Username: "123456",
		Password: "secret",
	}

	if err := validate.Struct(request); err != nil {
		t.Error(err.Error())
	}
}

func MustEqualsIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()

	if !ok {
		return false
	}

	first, ok := field.Field().Interface().(string)

	if !ok {
		return false
	}

	firstValue := strings.ToUpper(first)
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestCustomCrossFieldValidations(t *testing.T) {
	validate.RegisterValidation("eqic", MustEqualsIgnoreCase)

	type User struct {
		Username string `validate:"required,eqic=Email|eqic=Phone"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required,numeric"`
		Name     string `validate:"required"`
	}

	user := User{
		Username: "zAki@example.com",
		Email:    "zaki@example.com",
		Phone:    "08123456789",
		Name:     "zaki",
	}

	if err := validate.Struct(user); err != nil {
		t.Error(err.Error())
	}
}

type RegisterRequest struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,numeric"`
	Password string `validate:"required"`
}

func MustValidRegisterSuccess(level validator.StructLevel) {
	registerRequest, ok := level.Current().Interface().(RegisterRequest)
	if !ok {
		panic("Failed Convert Struct")
	}

	if registerRequest.Username == registerRequest.Email || registerRequest.Username == registerRequest.Phone {
		// -> Success
	} else {
		level.ReportError(registerRequest.Username, "Username", "username", "username", "")
	}
}

func TestStructLevelValidations(t *testing.T) {
	validate.RegisterStructValidation(MustValidRegisterSuccess, RegisterRequest{})

	request := RegisterRequest{
		Username: "08123456789",
		Email:    "zaki@example.com",
		Phone:    "08123456789",
		Password: "secret",
	}

	if err := validate.Struct(request); err != nil {
		panic(err.Error())
	}
}
