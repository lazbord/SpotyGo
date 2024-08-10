package api

import "errors"

func CheckCreditential(email string) (string, error) {
	if email != "lazareb" {
		return "", errors.New("invalid email")
	}
	return "monuserid", nil
}
