package repositories

import (
	"ntl-test/backend/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApplicationRepository struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) *ApplicationRepository {
	return &ApplicationRepository{
		db: db,
	}
}

func (r *ApplicationRepository) Create(
	application *models.Application,
) error {
	return r.db.Create(application).Error
}

func (r *ApplicationRepository) FindByID(
	id uuid.UUID,
) (*models.Application, error) {
	var application models.Application

	if err := r.db.
		Where("id = ?", id).
		First(&application).Error; err != nil {
		return nil, err
	}

	return &application, nil
}

func (r *ApplicationRepository) FindByUserID(
	userID uuid.UUID,
) ([]models.Application, error) {
	var applications []models.Application

	if err := r.db.
		Where("user_id = ?", userID).
		Find(&applications).Error; err != nil {
		return nil, err
	}

	return applications, nil
}

func (r *ApplicationRepository) FindAll() ([]models.Application, error) {
	var applications []models.Application

	if err := r.db.
		Preload("User").
		Find(&applications).Error; err != nil {
		return nil, err
	}

	return applications, nil
}

func (r *ApplicationRepository) FindByTransactionNo(
	transactionNo string,
) (*models.Application, error) {
	var application models.Application

	if err := r.db.
		Preload("User").
		Where("transaction_no = ?", transactionNo).
		First(&application).Error; err != nil {
		return nil, err
	}

	return &application, nil
}
