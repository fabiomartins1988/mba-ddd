package value_objects

import (
	"fmt"

	"github.com/google/uuid"
)

type Uuid struct {
	value string
}

func NewUuid(id ...string) (Uuid, error) {
	if len(id) > 0 && id[0] != "" {
		if _, err := uuid.Parse(id[0]); err != nil {
			return Uuid{}, fmt.Errorf("value %s must be a valid UUID", id[0])
		}
		return Uuid{value: id[0]}, nil
	}
	return Uuid{value: uuid.NewString()}, nil
}

func MustNewUuid(id ...string) Uuid {
	u, err := NewUuid(id...)
	if err != nil {
		panic(err)
	}
	return u
}

func (u Uuid) Value() string      { return u.value }
func (u Uuid) Equals(o Uuid) bool { return u.value == o.value }
func (u Uuid) String() string     { return u.value }
func (u Uuid) IsZero() bool       { return u.value == "" }
