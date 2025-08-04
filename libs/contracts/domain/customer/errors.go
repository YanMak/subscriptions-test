package customerDomainContract

import "errors"

var ErrSubscriptionNotFound = errors.New("subscription not found")
var ErrInvalidID = errors.New("invalid id")
var ErrInvalidUserID = errors.New("invalid userID")
