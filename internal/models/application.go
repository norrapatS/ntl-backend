package models

import (
	"time"

	"github.com/google/uuid"
)

type Application struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID        uuid.UUID `gorm:"type:uuid;not null;index" json:"userId"`
	User          User      `gorm:"foreignKey:UserID" json:"user"`
	NationalID    string    `gorm:"size:13;not null" json:"nationalId"`
	Prefix        string    `gorm:"size:20" json:"prefix"`
	FirstName     string    `gorm:"size:100;not null" json:"firstName"`
	LastName      string    `gorm:"size:100;not null" json:"lastName"`
	BirthDate     time.Time `json:"birthDate"`
	Gender        string    `gorm:"size:20" json:"gender"`
	MaritalStatus string    `gorm:"size:20" json:"maritalStatus"`

	HouseNo     string `gorm:"size:50" json:"houseNo"`
	Moo         string `gorm:"size:20" json:"moo"`
	Road        string `gorm:"size:100" json:"road"`
	SubDistrict string `gorm:"size:100" json:"subDistrict"`
	District    string `gorm:"size:100" json:"district"`
	Province    string `gorm:"size:100" json:"province"`
	Zipcode     string `gorm:"size:5" json:"zipcode"`

	Occupation    string  `gorm:"size:100" json:"occupation"`
	MonthlyIncome float64 `json:"monthlyIncome"`
	YearlyIncome  float64 `json:"yearlyIncome"`
	HouseholdSize int     `json:"householdSize"`
	Debt          float64 `json:"debt"`
	LandOwned     string  `gorm:"size:50" json:"landOwned"`

	Status string `gorm:"size:30;not null;default:'draft'" json:"status"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
