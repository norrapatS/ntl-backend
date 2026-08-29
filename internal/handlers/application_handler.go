package handlers

import (
	"net/http"

	"ntl-test/backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ApplicationHandler struct {
	applicationService *services.ApplicationService
}

func NewApplicationHandler(
	applicationService *services.ApplicationService,
) *ApplicationHandler {
	return &ApplicationHandler{
		applicationService: applicationService,
	}
}

type CreateApplicationRequest struct {
	NationalID    string  `json:"nationalId" binding:"required,len=13"`
	Prefix        string  `json:"prefix"`
	FirstName     string  `json:"firstName" binding:"required"`
	LastName      string  `json:"lastName" binding:"required"`
	BirthDate     string  `json:"birthDate"`
	Gender        string  `json:"gender"`
	MaritalStatus string  `json:"maritalStatus"`

	HouseNo     string `json:"houseNo"`
	Moo         string `json:"moo"`
	Road        string `json:"road"`
	SubDistrict string `json:"subDistrict"`
	District    string `json:"district"`
	Province    string `json:"province"`
	Zipcode     string `json:"zipcode"`

	Occupation    string  `json:"occupation"`
	MonthlyIncome float64 `json:"monthlyIncome"`
	YearlyIncome  float64 `json:"yearlyIncome"`
	HouseholdSize int     `json:"householdSize"`
	Debt          float64 `json:"debt"`
	LandOwned     string  `json:"landOwned"`
}

func (h *ApplicationHandler) Create(c *gin.Context) {
	var request CreateApplicationRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
		})
		return
	}

	userIDValue, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	userID, ok := userIDValue.(uuid.UUID)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid user id",
		})
		return
	}

	application, err := h.applicationService.Create(
		services.CreateApplicationInput{
			UserID:        userID,
			NationalID:    request.NationalID,
			Prefix:        request.Prefix,
			FirstName:     request.FirstName,
			LastName:      request.LastName,
			BirthDate:     request.BirthDate,
			Gender:        request.Gender,
			MaritalStatus: request.MaritalStatus,
			HouseNo:       request.HouseNo,
			Moo:           request.Moo,
			Road:          request.Road,
			SubDistrict:   request.SubDistrict,
			District:      request.District,
			Province:      request.Province,
			Zipcode:       request.Zipcode,
			Occupation:    request.Occupation,
			MonthlyIncome: request.MonthlyIncome,
			YearlyIncome:  request.YearlyIncome,
			HouseholdSize: request.HouseholdSize,
			Debt:          request.Debt,
			LandOwned:     request.LandOwned,
		},
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "application created successfully",
		"data":    application,
	})
}