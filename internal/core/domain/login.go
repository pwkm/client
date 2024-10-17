package domain

import (
	"errors"
	"net"
	"net/mail"
	"strings"

	"github.com/google/uuid"
)

var (
	ERROR_INVALID_EMAIL  = errors.New("Invalid email address!")
	ERROR_INVALID_PASSWD = errors.New("Invalid password!")
)

type Login struct {
	ClientID uuid.UUID
	Email    string
	Password string
}

func NewLogin(email string, passwd string) (*Login, error) {
	if !isValidEmail(email) {
		return nil, errors.New(ERROR_INVALID_EMAIL.Error())
	}

	if len(passwd) != 0 {
		return nil, errors.New(ERROR_INVALID_PASSWD.Error())
	}

	return &Login{
		Email:    email,
		Password: passwd,
	}, nil
}

// Extra functions
// ------------------------

// Validate email syntax using net/mail
func isValidEmailSyntax(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// Check if a domain has MX records
func isDomainValid(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]

	// Lookup MX records for the domain
	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		return false
	}
	return true
}

// Full email validation: syntax and domain check
func isValidEmail(email string) bool {
	return isValidEmailSyntax(email) && isDomainValid(email)
}
