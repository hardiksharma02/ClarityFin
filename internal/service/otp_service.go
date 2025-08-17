package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hardiksharma/clarityfin-api/internal/config"
	"github.com/hardiksharma/clarityfin-api/internal/domain"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// otpService implements the OTPService interface
type otpService struct {
	otpRepo      domain.OTPRepository
	smsConfig    config.SMSConfig
	twilioClient *twilio.RestClient
}

// NewOTPService creates a new instance of OTPService
func NewOTPService(otpRepo domain.OTPRepository, smsConfig config.SMSConfig) domain.OTPService {
	var twilioClient *twilio.RestClient
	if smsConfig.Provider == "twilio" {
		twilioClient = twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: smsConfig.Twilio.AccountSID,
			Password: smsConfig.Twilio.AuthToken,
		})
	}

	return &otpService{
		otpRepo:      otpRepo,
		smsConfig:    smsConfig,
		twilioClient: twilioClient,
	}
}

// GenerateOTP generates a new OTP for the given phone number
func (s *otpService) GenerateOTP(phoneNumber string) error {
	// Generate a 6-digit OTP
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	// Set expiration time (5 minutes from now)
	expiresAt := time.Now().Add(5 * time.Minute)

	// Create OTP record
	otp := &domain.OTP{
		PhoneNumber: phoneNumber,
		Code:        code,
		ExpiresAt:   expiresAt,
		IsUsed:      false,
	}

	// Save to database
	if err := s.otpRepo.Create(otp); err != nil {
		return err
	}

	// Send OTP via SMS
	return s.SendOTP(phoneNumber, code)
}

// VerifyOTP verifies the OTP for the given phone number
func (s *otpService) VerifyOTP(phoneNumber, code string) (bool, error) {
	// Find OTP in database
	otp, err := s.otpRepo.FindByPhoneNumberAndCode(phoneNumber, code)
	if err != nil {
		return false, err
	}

	// Mark OTP as used
	if err := s.otpRepo.MarkAsUsed(otp.ID); err != nil {
		return false, err
	}

	return true, nil
}

// SendOTP sends OTP via SMS using Twilio or MSG91
func (s *otpService) SendOTP(phoneNumber, code string) error {
	switch s.smsConfig.Provider {
	case "twilio":
		return s.sendViaTwilio(phoneNumber, code)
	case "msg91":
		return s.sendViaMSG91(phoneNumber, code)
	default:
		// For development/testing, just log the OTP
		fmt.Printf("OTP for %s: %s\n", phoneNumber, code)
		return nil
	}
}

// sendViaTwilio sends OTP via Twilio SMS
func (s *otpService) sendViaTwilio(phoneNumber, code string) error {
	if s.twilioClient == nil {
		return fmt.Errorf("Twilio client not initialized")
	}

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phoneNumber)
	params.SetFrom(s.smsConfig.Twilio.FromNumber)
	params.SetBody(fmt.Sprintf("Your ClarityFin verification code is: %s. Valid for 5 minutes.", code))

	_, err := s.twilioClient.Api.CreateMessage(params)
	return err
}

// sendViaMSG91 sends OTP via MSG91 SMS
func (s *otpService) sendViaMSG91(phoneNumber, code string) error {
	// Implementation for MSG91 would go here
	// For now, just log the OTP
	fmt.Printf("MSG91 OTP for %s: %s\n", phoneNumber, code)
	return nil
}
