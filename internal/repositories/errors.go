package repositories

import (
	"fmt"
)

type ErrDuplicateKey struct {
	Object string
	Key    string
}

func (e ErrDuplicateKey) Error() string {
	return fmt.Sprintf("%s already exists with key: %s", e.Object, e.Key)
}

type ErrNotFound struct {
	Object string
	Key    string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("failed to find %s identified with key: %s", e.Object, e.Key)
}
