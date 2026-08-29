package services

import (
	"ntl-test/backend/internal/models"
	"ntl-test/backend/internal/repositories"

	"github.com/google/uuid"
)

type AdminService struct {
	applicationRepository *repositories.ApplicationRepository
}

func NewAdminService(
	applicationRepository *repositories.ApplicationRepository,
) *AdminService {
	return &AdminService{
		applicationRepository: applicationRepository,
	}
}

func (s *AdminService) GetApplications() ([]models.Application, error) {
	return s.applicationRepository.FindAll()
}

func (s *AdminService) GetApplicationByID(
	id uuid.UUID,
) (*models.Application, error) {
	return s.applicationRepository.FindByID(id)
}