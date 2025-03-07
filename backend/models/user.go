package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email         string `gorm:"unique;not null"`
	Password      string `json:"-"`
	FullName      string
	Phone         string
	BloodGroup    string
	Allergies     string
	EmergencyContacts []EmergencyContact
	MedicalRecords   []MedicalRecord
	Location      Location
}

type EmergencyContact struct {
	gorm.Model
	UserID      uint
	Name        string
	Phone       string
	Relationship string
}

type Location struct {
	gorm.Model
	UserID      uint
	Latitude    float64
	Longitude   float64
	UpdatedAt   time.Time
}

type MedicalRecord struct {
	gorm.Model
	UserID      uint
	RecordType  string
	Details     string
	Timestamp   time.Time
} 