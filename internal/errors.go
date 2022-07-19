package internal

import "errors"

var (
	NotFoundErr = errors.New("ResourceNotFound")
	ConflictErr = errors.New("ResourceAlreadyExists")
)
