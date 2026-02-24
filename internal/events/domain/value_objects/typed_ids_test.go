package value_objects

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewCustomerID_NoIdProvided_ShouldCreateSuccessfully(t *testing.T) {
	id, err := NewCustomerID()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if id.Value() == "" {
		t.Errorf("expected non-empty customer id")
	}
}

func TestNewCustomerID_ValidIdProvided_ShouldCreateSuccessfully(t *testing.T) {
	validID := uuid.NewString()

	id, err := NewCustomerID(validID)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if id.Value() != validID {
		t.Errorf("expected value %s, got %s", validID, id.Value())
	}
}

func TestNewCustomerID_InvalidIdProvided_ShouldReturnError(t *testing.T) {
	invalidID := "invalid-uuid"

	id, err := NewCustomerID(invalidID)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !id.IsZero() {
		t.Errorf("expected IsZero to be true")
	}
}

func TestNewPartnerID_NoIdProvided_ShouldCreateSuccessfully(t *testing.T) {
	id, err := NewPartnerID()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if id.Value() == "" {
		t.Errorf("expected non-empty partner id")
	}
}

func TestNewEventID_NoIdProvided_ShouldCreateSuccessfully(t *testing.T) {
	id, err := NewEventID()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if id.Value() == "" {
		t.Errorf("expected non-empty event id")
	}
}

func TestNewEventSectionID_NoIdProvided_ShouldCreateSuccessfully(t *testing.T) {
	id, err := NewEventSectionID()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if id.Value() == "" {
		t.Errorf("expected non-empty event section id")
	}
}

func TestNewEventSpotID_NoIdProvided_ShouldCreateSuccessfully(t *testing.T) {
	id, err := NewEventSpotID()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if id.Value() == "" {
		t.Errorf("expected non-empty event spot id")
	}
}

func TestNewOrderID_NoIdProvided_ShouldCreateSuccessfully(t *testing.T) {
	id, err := NewOrderID()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if id.Value() == "" {
		t.Errorf("expected non-empty order id")
	}
}
