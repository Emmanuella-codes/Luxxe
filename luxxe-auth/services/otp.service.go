package services

import (
	"context"
	"strconv"

	TempStore "github.com/Emmanuella-codes/Luxxe/luxxe-storage"
)

type IssueOTPStruct struct {
	Email string 
	ExpirationTime int
	ManualOTP int
}

type IssueOTPRes struct {
	EmailOTP string
	TimeLeft int
}

func IssueOtp(ctx context.Context, ios *IssueOTPStruct) *IssueOTPRes {
	email, expirationTime, manualOtp := ios.Email, ios.ExpirationTime, ios.ManualOTP

	var otp int
	if manualOtp != 0 {
		otp = manualOtp
	} else {
		otp = GetFourRandomNumbers()
	}

	if expirationTime == 0 {
		expirationTime = 600
	}

	otpString := strconv.Itoa(otp)
	otpStringHash := GenerateStringHash(otpString)
	TempStore.Set(
		ctx,
		&TempStore.SetStruct{
			Key: 						email,
			Value: 					otpStringHash,
			ExpirationTime: expirationTime,
		},
	)

	return &IssueOTPRes{
		EmailOTP: otpString,
		TimeLeft: expirationTime / 60, // convert seconds to minutes
	}
}

type VerifyOtpStruct struct {
	Email     string
	Otp       string
	KeepAlive bool
}

// TODO: @security Hash the OTPs
func VerifyOtp(ctx context.Context, vos *VerifyOtpStruct) bool {
	email, otp, keepAlive := vos.Email, vos.Otp, vos.KeepAlive
	emailOtpHash, err := TempStore.Get(ctx, email)
	if err != nil {
		return false
	}

	if !CompareStringHash(emailOtpHash, otp) {
		return false
	}
	if !keepAlive {
		IssueOtp(ctx, &IssueOTPStruct{Email: email})
	}
	return true
}
