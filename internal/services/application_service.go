package services

import (
	"errors"

	"ntl-test/backend/internal/models"
	"ntl-test/backend/internal/repositories"

	"github.com/google/uuid"
)

type ApplicationService struct {
	applicationRepository *repositories.ApplicationRepository
}

func NewApplicationService(
	applicationRepository *repositories.ApplicationRepository,
) *ApplicationService {
	return &ApplicationService{
		applicationRepository: applicationRepository,
	}
}

type CreateApplicationInput struct {
	UserID uuid.UUID

	NationalID    string
	Prefix        string
	FirstName     string
	LastName      string
	BirthDate     string
	Gender        string
	MaritalStatus string

	HouseNo     string
	Moo         string
	Road        string
	SubDistrict string
	District    string
	Province    string
	Zipcode     string

	Occupation    string
	MonthlyIncome float64
	YearlyIncome  float64
	HouseholdSize int
	Debt          float64
	LandOwned     string
}

func (s *ApplicationService) Create(
	input CreateApplicationInput,
) (*models.Application, error) {
	if input.UserID == uuid.Nil {
		return nil, errors.New("user not authenticated")
	}

	if input.NationalID == "" {
		return nil, errors.New("national ID is required")
	}

	if input.FirstName == "" {
		return nil, errors.New("first name is required")
	}

	if input.LastName == "" {
		return nil, errors.New("last name is required")
	}

	application := &models.Application{
		ID:            uuid.New(),
		UserID:        input.UserID,
		NationalID:    input.NationalID,
		Prefix:        input.Prefix,
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		Gender:        input.Gender,
		MaritalStatus: input.MaritalStatus,
		HouseNo:       input.HouseNo,
		Moo:           input.Moo,
		Road:          input.Road,
		SubDistrict:   input.SubDistrict,
		District:      input.District,
		Province:      input.Province,
		Zipcode:       input.Zipcode,
		Occupation:    input.Occupation,
		MonthlyIncome: input.MonthlyIncome,
		YearlyIncome:  input.YearlyIncome,
		HouseholdSize: input.HouseholdSize,
		Debt:          input.Debt,
		LandOwned:     input.LandOwned,
		Status:        "draft",
	}

	if err := s.applicationRepository.Create(application); err != nil {
		return nil, err
	}

	return application, nil
}