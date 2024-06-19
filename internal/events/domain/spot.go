package domain

import "errors"

type SpotStatus string

var (
	ErrSpotNumberInvalid = errors.New("invalid spot number")
	ErrSpotNotFound = errors.New("spot not found")
	ErrSpotAlreadyReserved = errors.New("spot already reserved")
	ErrSpotNameLength = errors.New("spot name must be at least 2 characters long")
	ErrSpotNameRequired = errors.New("spot name is requires")
	ErrSpotNameStartWithLetter = errors.New("spot name must start with a letter")
	ErrSpotNameEndWithNumber = errors.New("spot name must end with a number")
)

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold SpotStatus = "sold"
)

type Spot struct {
	ID string
	EventID string
	Name string
	Status SpotStatus
	TicketID string
}

func (s Spot) Validate() error {
	if s.Name == "" {
		return ErrSpotNameRequired
	}

	if len(s.Name) < 2 {
		return ErrSpotNameLength
	}

	// Validates spot name format
	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotNameStartWithLetter
	}

	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrSpotNameEndWithNumber
	}

	return nil
}
