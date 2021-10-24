package life

import "errors"

const (
	notFoundError          = "not found"
	lastPageError          = "last page"
	incorrectPositionError = "incorrect position"
)

func NotFoundError() error {
	return errors.New(notFoundError)
}

func LastPageError() error {
	return errors.New(lastPageError)
}

func IncorrectPositionError() error {
	return errors.New(incorrectPositionError)
}
