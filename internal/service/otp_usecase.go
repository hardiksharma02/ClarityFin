package service

import (
	"github.com/hardiksharma/clarityfin-api/internal/domain"
)

// otpUseCase implements the OTPUseCase interface
type otpUseCase struct {
	otpService domain.OTPService
}

// NewOTPUseCase creates a new instance of OTPUseCase
func NewOTPUseCase(otpService domain.OTPService) domain.OTPUseCase {
	return &otpUseCase{
		otpService: otpService,
	}
}

// SendOTP handles OTP generation and sending
func (uc *otpUseCase) SendOTP(phoneNumber string) error {
	return uc.otpService.GenerateOTP(phoneNumber)
}

// VerifyOTP handles OTP verification
func (uc *otpUseCase) VerifyOTP(phoneNumber, code string) (bool, error) {
	return uc.otpService.VerifyOTP(phoneNumber, code)
}
