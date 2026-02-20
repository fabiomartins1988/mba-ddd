package value_objects

import (
	"testing"
)

func TestName_NewName_ValidName_ShouldCreateSuccessfully(t *testing.T) {
	// Arrange
	validName := "John Doe"

	// Act
	name, err := NewName(validName)

	// Assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if name.Value() != validName {
		t.Errorf("expected name value to be '%s', got '%s'", validName, name.Value())
	}
}

func TestName_NewName_ShortName_ShouldReturnError(t *testing.T) {
	// Arrange
	// Name needs at least 3 characters
	shortName := "Jo"

	// Act
	name, err := NewName(shortName)

	// Assert
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	expectedErrorMsg := "name must have at least 3 characters"
	if err.Error() != expectedErrorMsg {
		t.Errorf("expected error message '%s', got '%v'", expectedErrorMsg, err)
	}

	if name.Value() != "" {
		t.Errorf("expected empty name value, got '%s'", name.Value())
	}
}

func TestName_Equals_SameValues_ShouldReturnTrue(t *testing.T) {
	// Arrange
	name1, _ := NewName("John Doe")
	name2, _ := NewName("John Doe")

	// Act
	result := name1.Equals(name2)

	// Assert
	if result != true {
		t.Errorf("expected Equals to return true for matching names")
	}
}

func TestName_Equals_DifferentValues_ShouldReturnFalse(t *testing.T) {
	// Arrange
	name1, _ := NewName("John Doe")
	name2, _ := NewName("Jane Doe")

	// Act
	result := name1.Equals(name2)

	// Assert
	if result != false {
		t.Errorf("expected Equals to return false for different names")
	}
}

func TestName_Value_ShouldReturnStoredString(t *testing.T) {
	// Arrange
	expectedValue := "John Doe"
	name, _ := NewName(expectedValue)

	// Act
	result := name.Value()

	// Assert
	if result != expectedValue {
		t.Errorf("expected Value() to return '%s', got '%s'", expectedValue, result)
	}
}

func TestName_String_ShouldReturnStoredString(t *testing.T) {
	// Arrange
	expectedValue := "John Doe"
	name, _ := NewName(expectedValue)

	// Act
	result := name.String()

	// Assert
	if result != expectedValue {
		t.Errorf("expected String() to return '%s', got '%s'", expectedValue, result)
	}
}
