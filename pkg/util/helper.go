package util

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/labstack/gommon/log"
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

func ParseStringToDate(formatUpload, goalTime string) (*time.Time, error) {
	if goalTime == "" || goalTime == "0" || goalTime == "0.00" {
		goalTime = "01/01/2000"
	}

	dateParse, err := time.Parse(formatUpload, goalTime)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	oriDate := dateParse.Format("2006-01-02")
	periode, err := time.Parse("2006-01-02", oriDate)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &periode, nil
}

func ParseStringToFloat(source string) float64 {
	if source == "" {
		return 0
	}

	number, err := strconv.ParseFloat(source, 64)
	if err != nil {
		return 0
	}

	return number
}

func LastDay(source *time.Time) int {
	y, m, _ := source.Date()
	lastDay := time.Date(y, m+1, 0, 0, 0, 0, 0, time.UTC)
	return lastDay.Day()
}
