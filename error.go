package clrs

import (
	"errors"
)

var (
	errOverflow   = errors.New("overflow")
	errUnderflow  = errors.New("underflow")
	errOutOfSpace = errors.New("out of space")
)
