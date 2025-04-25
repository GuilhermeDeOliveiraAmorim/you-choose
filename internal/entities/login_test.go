package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogin(t *testing.T) {
	tests := []struct {
		email    string
		password string
		expected bool
	}{
		{"valid@example.com", "ValidPassword1@", true},
		{"invalid-email", "ValidPassword1@", false},
		{"valid@example.com", "short", false},
		{"valid@example.com", "validpassword", false},
	}

	for _, tt := range tests {
		t.Run(tt.email, func(t *testing.T) {
			login, validationErrors := NewLogin(tt.email, tt.password)

			if tt.expected {
				assert.NotNil(t, login)
				assert.Empty(t, validationErrors)
			} else {
				assert.Nil(t, login)
				assert.NotEmpty(t, validationErrors)
			}
		})
	}
}

func TestValidateLogin(t *testing.T) {
	tests := []struct {
		email    string
		password string
		expected int
	}{
		{"valid@example.com", "ValidPassword1@", 0},
		{"invalid-email", "ValidPassword1@", 1},
		{"valid@example.com", "short", 1},
		{"valid@example.com", "validpassword", 1},
	}

	for _, tt := range tests {
		t.Run(tt.email, func(t *testing.T) {
			validationErrors := ValidateLogin(tt.email, tt.password)
			assert.Len(t, validationErrors, tt.expected)
		})
	}
}

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"valid@example.com", true},
		{"invalid-email", false},
		{"@example.com", false},
		{"valid@subdomain.example.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.email, func(t *testing.T) {
			result := isValidEmail(tt.email)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"asdASD123!@#", true},
		{"12345", false},
		{"qwert", false},
		{"ASDFG", false},
		{"NoDigitsAndSpecial", false},
		{"Valid1@", true},
		{"validpassword", false},
		{"VALIDPASSWORD", false},
	}

	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			result := isValidPassword(tt.password)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestHashEmail(t *testing.T) {
	lo := &Login{
		Email:    "valid@example.com",
		Password: "ValidPassword1@",
	}

	originalEmail := lo.Email
	lo.HashEmail()

	assert.NotEqual(t, originalEmail, lo.Email)
}

func TestHashPassword(t *testing.T) {
	lo := &Login{
		Email:    "valid@example.com",
		Password: "ValidPassword1@",
	}

	originalPassword := lo.Password
	err := lo.HashPassword()

	assert.Nil(t, err)

	assert.NotEqual(t, originalPassword, lo.Password)
}

func TestVerifyEmail(t *testing.T) {
	lo := &Login{
		Email:    "valid@example.com",
		Password: "ValidPassword1@",
	}
	lo.HashEmail()

	result := lo.VerifyEmail("valid@example.com")
	assert.True(t, result)

	result = lo.VerifyEmail("incorrect@example.com")
	assert.False(t, result)
}

func TestVerifyPassword(t *testing.T) {
	lo := &Login{
		Email:    "valid@example.com",
		Password: "ValidPassword1@",
	}
	err := lo.HashPassword()
	assert.Nil(t, err)

	result := lo.VerifyPassword("ValidPassword1@")
	assert.True(t, result)

	result = lo.VerifyPassword("InvalidPassword")
	assert.False(t, result)
}

func TestChangeEmail(t *testing.T) {
	lo := &Login{
		Email:    "valid@example.com",
		Password: "ValidPassword1@",
	}

	lo.ChangeEmail("new@example.com")

	assert.Equal(t, "new@example.com", lo.Email)
}

func TestChangePassword(t *testing.T) {
	lo := &Login{
		Email:    "valid@example.com",
		Password: "ValidPassword1@",
	}

	lo.ChangePassword("NewValidPassword1@")

	assert.Equal(t, "NewValidPassword1@", lo.Password)
}

func TestLoginEquals(t *testing.T) {
	lo1 := &Login{
		Email: "valid@example.com",
	}

	lo1.HashEmail()

	lo2 := &Login{
		Email: "valid@example.com",
	}

	lo2.HashEmail()

	assert.True(t, lo1.Equals(lo2))
}
