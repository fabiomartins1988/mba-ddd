package entities

import (
	"github.com/fabiomartins1988/mba-ddd/internal/common/domain"
	"github.com/fabiomartins1988/mba-ddd/internal/events/domain/value_objects"
)

type Customer struct {
	domain.AggregateRoot
	id   string
	name value_objects.Name
	cpf  value_objects.Cpf
}

func NewCustomer(id string, name value_objects.Name, cpf value_objects.Cpf) *Customer {
	return &Customer{
		id:   id,
		name: name,
		cpf:  cpf,
	}
}

func (c *Customer) ID() string {
	return c.id
}

func (c *Customer) Name() value_objects.Name {
	return c.name
}

func (c *Customer) Cpf() value_objects.Cpf {
	return c.cpf
}
