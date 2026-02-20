package domain

import "time"

type DomainEvent interface {
	AggregateID() string
	Name() string
	OccurredAt() time.Time
}
