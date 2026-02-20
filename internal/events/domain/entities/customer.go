package entities

import "github.com/fabiomartins1988/mba-ddd/internal/common/domain"

type Customer struct {
	domain.AggregateRoot
	id   string
	name string
	cpf  string
}

func NewCustomer(id string, name string, cpf string) *Customer {
	return &Customer{
		id:   id,
		name: name,
		cpf:  cpf,
	}
}

func (c *Customer) ID() string {
	return c.id
}

func (c *Customer) Name() string {
	return c.name
}

func (c *Customer) Cpf() string {
	return c.cpf
}
