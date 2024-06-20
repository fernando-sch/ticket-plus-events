package service

type ReservationRequest struct {
	Spots []string `json:"spots"`
	EventID string `json:"event_id"`
	TicketType string `json:"ticket_type"`
	Email string `json:"email"`
}

type ReservationResponse struct {
	ID string `json:"id"`
	Email string `json:"email"`
	Spot string `json:"spot"`
	TicketType string `json:"ticket_type"`
	Status string `json:"status"`
	EventID string `json:"event_id"`
}

type Partner interface {
	MakeReservation(req *ReservationRequest) ([]ReservationResponse, error)
}