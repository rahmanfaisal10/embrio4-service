package util

import (
	"math/rand"
	"time"
)

const (
	otpChars = "1234567890"
)

func TimeNow() *time.Time {
	timer := time.Now()
	return &timer
}

func GenerateOTP(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

func ParseEmailString(req string) []string {
	return []string{
		req,
	}
}
