package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hardiksharma/clarityfin-api/internal/domain"
	"github.com/hardiksharma/clarityfin-api/internal/dto"
	"github.com/hardiksharma/clarityfin-api/pkg/response"
)

// OTPHandler handles OTP-related HTTP requests
type OTPHandler struct {
	otpUseCase domain.OTPUseCase
}

// NewOTPHandler creates a new instance of OTPHandler
func NewOTPHandler(otpUseCase domain.OTPUseCase) *OTPHandler {
	return &OTPHandler{
		otpUseCase: otpUseCase,
	}
}

// SendOTP handles OTP sending
func (h *OTPHandler) SendOTP(c *gin.Context) {
	var req dto.SendOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.otpUseCase.SendOTP(req.PhoneNumber)
	if err != nil {
		response.InternalServerError(c, "Failed to send OTP")
		return
	}

	response.Success(c, nil, "OTP sent successfully")
}

// VerifyOTP handles OTP verification
func (h *OTPHandler) VerifyOTP(c *gin.Context) {
	var req dto.VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	valid, err := h.otpUseCase.VerifyOTP(req.PhoneNumber, req.Code)
	if err != nil {
		response.BadRequest(c, "Invalid OTP")
		return
	}

	if !valid {
		response.BadRequest(c, "Invalid OTP")
		return
	}

	response.Success(c, nil, "OTP verified successfully")
}
