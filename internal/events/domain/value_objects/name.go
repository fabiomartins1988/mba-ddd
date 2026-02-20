package value_objects

import "fmt"

type Name struct {
	value string
}

func NewName(name string) (Name, error) {
	if len(name) < 3 {
		return Name{}, fmt.Errorf("name must have at least 3 characters")
	}
	return Name{value: name}, nil
}

func (n Name) Value() string {
	return n.value
}

func (n Name) Equals(other Name) bool {
	return n.value == other.value
}

func (n Name) String() string {
	return n.value
}
