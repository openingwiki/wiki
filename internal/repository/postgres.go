package repository

import "errors"

func IsNotFound(err error) bool { return errors.Is(err, errNotFound) }


