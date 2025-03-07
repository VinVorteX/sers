package services

import (
	"fmt"
	"os"
	"sers/models"
	"time"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"gorm.io/gorm"
)

type SOSService struct {
	db           *gorm.DB
	twilioClient *twilio.RestClient
}

func NewSOSService(db *gorm.DB) *SOSService {
	// Initialize Twilio client
	twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})

	return &SOSService{
		db:           db,
		twilioClient: twilioClient,
	}
}

func (s *SOSService) TriggerEmergency(userID uint, latitude, longitude float64, message string) error {
	// Start transaction
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Update user location
	if err := tx.Model(&models.Location{}).
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{
			"latitude":   latitude,
			"longitude":  longitude,
			"updated_at": time.Now(),
		}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update location: %v", err)
	}

	// Get user details and emergency contacts
	var user models.User
	if err := tx.Preload("EmergencyContacts").First(&user, userID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to fetch user details: %v", err)
	}

	// Find nearby responders (within 5km radius)
	var responders []models.User
	if err := s.findNearbyResponders(latitude, longitude, 5.0, &responders); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to find responders: %v", err)
	}

	// Send SMS notifications to emergency contacts
	for _, contact := range user.EmergencyContacts {
		if err := s.sendEmergencySMS(contact.Phone, user.FullName, latitude, longitude); err != nil {
			// Log error but continue with other notifications
			fmt.Printf("Failed to send SMS to %s: %v\n", contact.Phone, err)
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (s *SOSService) findNearbyResponders(lat, lon, radiusKm float64, responders *[]models.User) error {
	// Approximate distance calculation using Haversine formula
	// Note: This is a simplified version. For production, consider using PostGIS
	query := `
		SELECT * FROM users u
		JOIN locations l ON u.id = l.user_id
		WHERE (6371 * acos(cos(radians(?)) * cos(radians(l.latitude)) * 
		cos(radians(l.longitude) - radians(?)) + 
		sin(radians(?)) * sin(radians(l.latitude)))) < ?
	`
	return s.db.Raw(query, lat, lon, lat, radiusKm).Scan(responders).Error
}

func (s *SOSService) sendEmergencySMS(to, userName string, lat, lon float64) error {
	fromNumber := os.Getenv("TWILIO_PHONE_NUMBER")
	messageBody := fmt.Sprintf("EMERGENCY: %s needs immediate assistance! Location: https://maps.google.com/?q=%f,%f",
		userName, lat, lon)
	params := &twilioApi.CreateMessageParams{
		To:   &to,
		From: &fromNumber,
		Body: &messageBody,
	}

	_, err := s.twilioClient.Api.CreateMessage(params)
	return err
}
