package clrs

import (
	"errors"
)

var (
	errOverflow   = errors.New("overflow")
	errUnderflow  = errors.New("underflow")
	errOutOfSpace = errors.New("out of space")
	errExist      = errors.New("already exists")
	errNotExist   = errors.New("does not exist")
)
