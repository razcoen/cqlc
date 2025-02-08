package gocqlc

import "errors"

var ErrNilSession = errors.New("session cannot be nil")
var ErrClosedSession = errors.New("session cannot be closed")
