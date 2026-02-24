package entities

import (
	"fmt"

	"github.com/fabiomartins1988/mba-ddd/internal/common/domain"
	"github.com/fabiomartins1988/mba-ddd/internal/events/domain/value_objects"
)

type Customer struct {
	domain.AggregateRoot
	id   value_objects.CustomerID
	name value_objects.Name
	cpf  value_objects.Cpf
}

func NewCustomer(name, cpf string) (*Customer, error) {
	nameVO, err := value_objects.NewName(name)
	if err != nil {
		return nil, fmt.Errorf("invalid name: %w", err)
	}

	cpfVO, err := value_objects.NewCpf(cpf)
	if err != nil {
		return nil, fmt.Errorf("invalid cpf: %w", err)
	}

	id, _ := value_objects.NewCustomerID()

	return &Customer{
		id:   id,
		name: nameVO,
		cpf:  cpfVO,
	}, nil
}

// Reconstituicao a partir do banco (sem validacao, sem eventos)
func RestoreCustomer(id string, name string, cpf string) (*Customer, error) {
	customerID, err := value_objects.NewCustomerID(id)
	if err != nil {
		return nil, err
	}
	cpfVO, err := value_objects.NewCpf(cpf)
	if err != nil {
		return nil, err
	}
	nameVO, _ := value_objects.NewName(name)

	return &Customer{
		id:   customerID,
		cpf:  cpfVO,
		name: nameVO,
	}, nil
}

func (c *Customer) ChangeName(name string) error {
	nameVO, err := value_objects.NewName(name)
	if err != nil {
		return err
	}
	c.name = nameVO
	return nil
}

func (c *Customer) ID() value_objects.CustomerID { return c.id }
func (c *Customer) Cpf() value_objects.Cpf       { return c.cpf }
func (c *Customer) Name() value_objects.Name     { return c.name }

func (c *Customer) Equals(other *Customer) bool {
	if other == nil {
		return false
	}
	return c.id.Equals(other.id.Uuid)
}
