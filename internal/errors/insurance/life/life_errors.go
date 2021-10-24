package life

import "errors"

var (
	NotFoundError          = errors.New("not found")
	LastPageError          = errors.New("last page")
	IncorrectPositionError = errors.New("incorrect position")
)
