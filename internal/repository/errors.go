package repository

import "errors"

var ErrNotFound = errors.New("quote not found")
var ErrDuplicate = errors.New("duplicate quote")
