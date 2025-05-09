package utils

import (
	"fmt"
	"math/rand"
)

func RandomOTP(length int) string {
	if length <= 0 {
		return ""
	}

	otp := make([]byte, 0, length)
	var lastDigit byte = 255
	repeatCount := 0

	for len(otp) < length {
		digit := byte(rand.Intn(10)) // عدد بین 0 تا 9

		if digit == lastDigit {
			repeatCount++
			if repeatCount >= 2 {
				continue // این رقم را رد کن چون باعث ۳ تکرار پشت سر هم می‌شود
			}
		} else {
			lastDigit = digit
			repeatCount = 0
		}

		otp = append(otp, digit)
	}

	// تبدیل اعداد به رشته
	result := ""
	for _, d := range otp {
		result += fmt.Sprintf("%d", d)
	}
	return result
}
