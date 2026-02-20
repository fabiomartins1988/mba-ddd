package value_objects

import (
	"testing"
)

func TestCpf_NewCpf_ValidCharacters_ShouldCreateSuccessfully(t *testing.T) {
	validCpfRaw := "12345678909"

	cpf, err := NewCpf(validCpfRaw)
	// Assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if cpf.value != validCpfRaw {
		t.Errorf("expected value %s, got %s", validCpfRaw, cpf.value)
	}
}

func TestCpf_NewCpf_ValidWithFormatting_ShouldStripNonDigitsAndCreateSuccessfully(t *testing.T) {
	// Arrange
	formattedCpf := "123.456.789-09"
	expectedCpfValue := "12345678909"

	// Act
	cpf, err := NewCpf(formattedCpf)

	// Assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if cpf.value != expectedCpfValue {
		t.Errorf("expected value %s, got %s", expectedCpfValue, cpf.value)
	}
}

func TestCpf_NewCpf_InvalidLength_ShouldReturnError(t *testing.T) {
	// Arrange
	invalidCpf := "123"

	// Act
	_, err := NewCpf(invalidCpf)

	// Assert
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	expectedErrorMsg := "CPF must have 11 digits, got 3"
	if err.Error() != expectedErrorMsg {
		t.Errorf("expected error message '%s', got '%v'", expectedErrorMsg, err)
	}
}

func TestCpf_NewCpf_AllSameDigits_ShouldReturnError(t *testing.T) {
	// Arrange
	invalidCpf := "11111111111"

	// Act
	_, err := NewCpf(invalidCpf)

	// Assert
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	expectedErrorMsg := "CPF must not be all same digits"
	if err.Error() != expectedErrorMsg {
		t.Errorf("expected error message '%s', got '%v'", expectedErrorMsg, err)
	}
}

func TestCpf_NewCpf_InvalidCheckDigits_ShouldReturnError(t *testing.T) {
	// Arrange
	invalidCpf := "12345678900" // Invalid check digits

	// Act
	_, err := NewCpf(invalidCpf)

	// Assert
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	expectedErrorMsg := "CPF is invalid"
	if err.Error() != expectedErrorMsg {
		t.Errorf("expected error message '%s', got '%v'", expectedErrorMsg, err)
	}
}
