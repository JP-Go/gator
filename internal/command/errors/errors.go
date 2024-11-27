package errors

import "errors"

var ErrFeedNotFound = errors.New("Feed not found")
var ErrUserNotRegistered = errors.New("Current user not registered. Please register")
var ErrNotLoggedIn = errors.New("Not logged in")
