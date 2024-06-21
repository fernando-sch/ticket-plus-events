package usecase

import "github.com/fernando-sch/ticket-plus-events/internal/events/domain"

type ListEventsOutputDTO struct {
	Events []EventDTO
}

type ListEventsUseCase struct {
	repo domain.EventRepository
}

func NewListEventUseCase(repo domain.EventRepository) *ListEventsUseCase {
	return &ListEventsUseCase{repo: repo}
}

func (uc *ListEventsUseCase) Execute() (*ListEventsOutputDTO, error) {
	events, err := uc.repo.ListEvents()

	if err != nil {
		return nil, err
	}

	eventDTOs := make([]EventDTO, len(events))

	for i, event := range events {
		eventDTOs[i] = EventDTO{
			ID:           event.ID,
			Name:         event.Name,
			Location:     event.Location,
			Organization: event.Organization,
			Rating:       string(event.Rating),
			Date:         event.Date.Format("2006-01-02 15:04:05"),
			Capacity:     event.Capacity,
			Price:        event.Price,
			PartnerID:    event.PartnerID,
			ImageURL: 		event.ImageURL,
		}
	}

	return &ListEventsOutputDTO{Events: eventDTOs}, nil
}
