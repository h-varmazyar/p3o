package utils

import (
	"github.com/h-varmazyar/gopet/phone"
	"regexp"
)

func IsValidMobile(input string) bool {
	_, err := phone.GetPhoneNumberDetails(input)
	if err != nil {
		return false
	}
	return true
}

func IsValidEmail(input string) bool {
	r, err := regexp.Compile("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$")
	if err != nil {
		return false
	}
	return r.MatchString(input)
}
