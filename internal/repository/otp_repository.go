package repository

import (
	"time"

	"github.com/hardiksharma/clarityfin-api/internal/domain"
	"gorm.io/gorm"
)

// otpRepository implements the OTPRepository interface
type otpRepository struct {
	db *gorm.DB
}

// NewOTPRepository creates a new instance of OTPRepository
func NewOTPRepository(db *gorm.DB) domain.OTPRepository {
	return &otpRepository{db: db}
}

// Create creates a new OTP in the database
func (r *otpRepository) Create(otp *domain.OTP) error {
	return r.db.Create(otp).Error
}

// FindByPhoneNumberAndCode finds an OTP by phone number and code
func (r *otpRepository) FindByPhoneNumberAndCode(phoneNumber, code string) (*domain.OTP, error) {
	var otp domain.OTP
	err := r.db.Where("phone_number = ? AND code = ? AND is_used = ? AND expires_at > ?",
		phoneNumber, code, false, time.Now()).First(&otp).Error
	if err != nil {
		return nil, err
	}
	return &otp, nil
}

// MarkAsUsed marks an OTP as used
func (r *otpRepository) MarkAsUsed(id uint) error {
	return r.db.Model(&domain.OTP{}).Where("id = ?", id).Update("is_used", true).Error
}

// DeleteExpired deletes expired OTPs
func (r *otpRepository) DeleteExpired() error {
	return r.db.Where("expires_at < ?", time.Now()).Delete(&domain.OTP{}).Error
}
