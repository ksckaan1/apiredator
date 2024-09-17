package work

import "errors"

var (
	ErrRequestTimeout = errors.New("work timeout")
	ErrStopWork    = errors.New("stop work")
)
