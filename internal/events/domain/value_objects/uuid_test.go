package value_objects

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewUuid_NoIdProvided_ShouldCreateSuccessfully(t *testing.T) {
	u, err := NewUuid()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if u.Value() == "" {
		t.Errorf("expected non-empty uuid string")
	}
	_, errParse := uuid.Parse(u.Value())
	if errParse != nil {
		t.Errorf("expected valid uuid, got parse error: %v", errParse)
	}
	if u.IsZero() {
		t.Errorf("expected IsZero to be false")
	}
}

func TestNewUuid_ValidIdProvided_ShouldCreateSuccessfully(t *testing.T) {
	validID := uuid.NewString()

	u, err := NewUuid(validID)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if u.Value() != validID {
		t.Errorf("expected value %s, got %s", validID, u.Value())
	}
	if u.String() != validID {
		t.Errorf("expected string %s, got %s", validID, u.String())
	}
}

func TestNewUuid_InvalidIdProvided_ShouldReturnError(t *testing.T) {
	invalidID := "invalid-uuid"

	u, err := NewUuid(invalidID)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	expectedErrorMsg := "value invalid-uuid must be a valid UUID"
	if err.Error() != expectedErrorMsg {
		t.Errorf("expected error message '%s', got '%v'", expectedErrorMsg, err)
	}
	if !u.IsZero() {
		t.Errorf("expected IsZero to be true, got false")
	}
}

func TestMustNewUuid_NoIdProvided_ShouldNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("expected no panic, got %v", r)
		}
	}()

	u := MustNewUuid()

	if u.Value() == "" {
		t.Errorf("expected non-empty uuid string")
	}
}

func TestMustNewUuid_ValidIdProvided_ShouldNotPanic(t *testing.T) {
	validID := uuid.NewString()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("expected no panic, got %v", r)
		}
	}()

	u := MustNewUuid(validID)

	if u.Value() != validID {
		t.Errorf("expected value %s, got %s", validID, u.Value())
	}
}

func TestMustNewUuid_InvalidIdProvided_ShouldPanic(t *testing.T) {
	invalidID := "invalid-uuid"

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic, got none")
		}
	}()

	_ = MustNewUuid(invalidID)
}

func TestUuidEquals_MatchingUuids_ShouldReturnTrue(t *testing.T) {
	validID := uuid.NewString()
	u1 := MustNewUuid(validID)
	u2 := MustNewUuid(validID)

	isEqual := u1.Equals(u2)

	if !isEqual {
		t.Errorf("expected Equals to return true for matching uuids")
	}
}

func TestUuidEquals_DifferentUuids_ShouldReturnFalse(t *testing.T) {
	u1 := MustNewUuid()
	u2 := MustNewUuid()

	isEqual := u1.Equals(u2)

	if isEqual {
		t.Errorf("expected Equals to return false for different uuids")
	}
}
