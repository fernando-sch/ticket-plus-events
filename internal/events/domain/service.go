package domain

import (
	"errors"
	"fmt"
)

type spotService struct{}

var (
	ErrInvalidQuantity = errors.New("quantity must be greater than zero")
)

func NewSpotService() *spotService {
	return &spotService{}
}

func (s *spotService) GenerateSpots(event *Event, quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	for i := range quantity {
		spotName := fmt.Sprintf("%c%d", 'A'+i/10, i%10+1)
		_ ,err := event.AddSpot(spotName)

		if err != nil {
			return err
		}
	}

	return nil
}