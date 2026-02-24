package value_objects

type CustomerID struct {
	Uuid
}

type PartnerID struct {
	Uuid
}

type EventID struct {
	Uuid
}

type EventSectionID struct {
	Uuid
}

type EventSpotID struct {
	Uuid
}

type OrderID struct {
	Uuid
}

func NewCustomerID(id ...string) (CustomerID, error) {
	u, err := NewUuid(id...)
	return CustomerID{Uuid: u}, err
}

func NewPartnerID(id ...string) (PartnerID, error) {
	u, err := NewUuid(id...)
	return PartnerID{Uuid: u}, err
}

func NewEventID(id ...string) (EventID, error) {
	u, err := NewUuid(id...)
	return EventID{Uuid: u}, err
}

func NewEventSectionID(id ...string) (EventSectionID, error) {
	u, err := NewUuid(id...)
	return EventSectionID{Uuid: u}, err
}

func NewEventSpotID(id ...string) (EventSpotID, error) {
	u, err := NewUuid(id...)
	return EventSpotID{Uuid: u}, err
}

func NewOrderID(id ...string) (OrderID, error) {
	u, err := NewUuid(id...)
	return OrderID{Uuid: u}, err
}
