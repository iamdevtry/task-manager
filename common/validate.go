package common

import (
	"errors"
	"net/mail"
	"regexp"
)

func ValidMailAddress(address string) error {
	_, err := mail.ParseAddress(address)
	if err != nil {
		return err
	}

	return nil
}

func ValidPhoneNumber(phoneNumber string) error {
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

	if !re.MatchString(phoneNumber) {
		return errors.New("invalid phone number")
	}

	return nil
}
