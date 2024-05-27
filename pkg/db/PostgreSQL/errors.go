package db

import (
	"github.com/h-varmazyar/p3o/pkg/errors"
)

// err format 910ab
var (
	ErrInvalidDB = errors.NewWithCode("invalid_db", 91000)
)
