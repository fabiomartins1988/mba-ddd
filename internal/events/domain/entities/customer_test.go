package entities

import (
	"testing"

	"github.com/fabiomartins1988/mba-ddd/internal/events/domain/value_objects"
	"github.com/google/uuid"
)

func TestCustomer_NewCustomer_ValidInput_ShouldCreateSuccessfully(t *testing.T) {
	validName := "John Doe"
	validCpf := "12345678909"

	customer, err := NewCustomer(validName, validCpf)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if customer.Name().Value() != validName {
		t.Errorf("expected name %s, got %s", validName, customer.Name().Value())
	}

	expectedCpf, _ := value_objects.NewCpf(validCpf)
	if customer.Cpf() != expectedCpf {
		t.Errorf("expected cpf %v, got %v", expectedCpf, customer.Cpf())
	}
	if customer.ID().IsZero() {
		t.Errorf("expected customer ID to not be zero")
	}
}

func TestCustomer_NewCustomer_InvalidName_ShouldReturnError(t *testing.T) {
	invalidName := "Jo"
	validCpf := "12345678909"

	customer, err := NewCustomer(invalidName, validCpf)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if customer != nil {
		t.Errorf("expected customer to be nil")
	}
}

func TestCustomer_NewCustomer_InvalidCpf_ShouldReturnError(t *testing.T) {
	validName := "John Doe"
	invalidCpf := "123"

	customer, err := NewCustomer(validName, invalidCpf)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if customer != nil {
		t.Errorf("expected customer to be nil")
	}
}

func TestCustomer_RestoreCustomer_ValidInput_ShouldRestoreSuccessfully(t *testing.T) {
	id := uuid.NewString()
	name := "John Doe"
	cpf := "12345678909"

	customer, err := RestoreCustomer(id, name, cpf)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if customer.ID().Value() != id {
		t.Errorf("expected id %s, got %s", id, customer.ID().Value())
	}
	if customer.Name().Value() != name {
		t.Errorf("expected name %s, got %s", name, customer.Name().Value())
	}

	expectedCpf, _ := value_objects.NewCpf(cpf)
	if customer.Cpf() != expectedCpf {
		t.Errorf("expected cpf %v, got %v", expectedCpf, customer.Cpf())
	}
}

func TestCustomer_RestoreCustomer_InvalidID_ShouldReturnError(t *testing.T) {
	invalidID := "invalid"
	name := "John Doe"
	cpf := "12345678909"

	customer, err := RestoreCustomer(invalidID, name, cpf)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if customer != nil {
		t.Errorf("expected customer to be nil")
	}
}

func TestCustomer_ChangeName_ValidName_ShouldChangeSuccessfully(t *testing.T) {
	customer, _ := NewCustomer("John Doe", "12345678909")
	newName := "Jane Doe"

	err := customer.ChangeName(newName)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if customer.Name().Value() != newName {
		t.Errorf("expected name %s, got %s", newName, customer.Name().Value())
	}
}

func TestCustomer_ChangeName_InvalidName_ShouldReturnError(t *testing.T) {
	customer, _ := NewCustomer("John Doe", "12345678909")
	invalidName := "Jo"

	err := customer.ChangeName(invalidName)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if customer.Name().Value() == invalidName {
		t.Errorf("expected name not to be changed")
	}
}

func TestCustomer_Equals_SameID_ShouldReturnTrue(t *testing.T) {
	id := uuid.NewString()
	customer1, err1 := RestoreCustomer(id, "John Doe", "12345678909")
	if err1 != nil {
		t.Fatalf("unexpected error: %v", err1)
	}
	customer2, err2 := RestoreCustomer(id, "Jane Doe", "12345678909")
	if err2 != nil {
		t.Fatalf("unexpected error: %v", err2)
	}

	isEqual := customer1.Equals(customer2)

	if !isEqual {
		t.Errorf("expected Equals to return true for matching IDs")
	}
}

func TestCustomer_Equals_DifferentID_ShouldReturnFalse(t *testing.T) {
	customer1, err1 := NewCustomer("John Doe", "12345678909")
	if err1 != nil {
		t.Fatalf("unexpected error: %v", err1)
	}
	customer2, err2 := NewCustomer("Jane Doe", "12345678909")
	if err2 != nil {
		t.Fatalf("unexpected error: %v", err2)
	}

	isEqual := customer1.Equals(customer2)

	if isEqual {
		t.Errorf("expected Equals to return false for different IDs")
	}
}

func TestCustomer_Equals_NilOther_ShouldReturnFalse(t *testing.T) {
	customer, _ := NewCustomer("John Doe", "12345678909")

	isEqual := customer.Equals(nil)

	if isEqual {
		t.Errorf("expected Equals to return false for nil")
	}
}
