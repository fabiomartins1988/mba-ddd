package domain

type AggregateRoot struct {
	events []DomainEvent
}

func (a *AggregateRoot) AddEvent(event DomainEvent) {
	a.events = append(a.events, event)
}

func (a *AggregateRoot) ClearEvents() {
	a.events = nil
}

func (a *AggregateRoot) Events() []DomainEvent {
	return a.events
}
