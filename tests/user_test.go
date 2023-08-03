package tests

import (
	"testing"
	"websiteapi/service"
)

func TestIsValidEmail(t *testing.T) {
	if !service.IsValidEmail("test@example.com") {
		t.Error("Expected true for valid email")
	}

	if service.IsValidEmail("test@.com") {
		t.Error("Expected false for invalid email")
	}
}

func TestIsValidDate(t *testing.T) {
	if !service.IsValidDate("2023-04-06") {
		t.Error("Expected true for valid date")
	}

	if service.IsValidDate("2023-13-06") {
		t.Error("Expected false for invalid date")
	}
}

func TestIsValidUserID(t *testing.T) {
	if !service.IsValidUserID("12345") {
		t.Error("Expected true for valid user_id")
	}

	if service.IsValidUserID("1234a") {
		t.Error("Expected false for invalid user_id")
	}
}
