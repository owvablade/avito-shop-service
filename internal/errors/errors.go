package errors

import (
	"net/http"
)

type Error struct {
	StatusCode int
	Message    string
}

func (e *Error) Error() string {
	return e.Message
}

var (
	ErrUserNotFound         = &Error{StatusCode: http.StatusBadRequest, Message: "user not found"}
	ErrMerchNotFound        = &Error{StatusCode: http.StatusBadRequest, Message: "merch not found"}
	ErrInsufficientCoins    = &Error{StatusCode: http.StatusBadRequest, Message: "not enough coins"}
	ErrMerchNameNotFound    = &Error{StatusCode: http.StatusBadRequest, Message: "merch name not found"}
	ErrTransferToYourself   = &Error{StatusCode: http.StatusBadRequest, Message: "cannot transfer to yourself"}
	ErrRequiredFieldNotSet  = &Error{StatusCode: http.StatusBadRequest, Message: "not all required fields are set"}
	ErrNegativeOrZeroAmount = &Error{StatusCode: http.StatusBadRequest, Message: "amount cant be negative or zero"}

	ErrWrongPassword = &Error{StatusCode: http.StatusUnauthorized, Message: "wrong password"}
	ErrTokenNotValid = &Error{StatusCode: http.StatusUnauthorized, Message: "token not valid"}
	ErrAuthorization = &Error{StatusCode: http.StatusUnauthorized, Message: "authorization failed"}

	ErrInternalServer          = &Error{StatusCode: http.StatusInternalServerError, Message: "internal server error"}
	ErrUnexpectedSigningMethod = &Error{StatusCode: http.StatusInternalServerError, Message: "unexpected signing method"}
)
