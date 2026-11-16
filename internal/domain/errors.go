package domain

import (
	"errors"
	"fmt"
)

type Fault struct {
	Status  int
	Code    string
	Message string
}

func (e *Fault) Error() string { return e.Message }
func Invalid(format string, args ...any) error {
	return &Fault{400, "invalid", fmt.Sprintf(format, args...)}
}
func Missing(kind, id string) error {
	return &Fault{404, "not_found", fmt.Sprintf("%s %s does not exist", kind, id)}
}
func Conflict(message string) error { return &Fault{409, "conflict", message} }
func Capacity(message string) error { return &Fault{429, "capacity", message} }
func Classify(err error) (int, string, string) {
	var f *Fault
	if errors.As(err, &f) {
		return f.Status, f.Code, f.Message
	}
	return 500, "internal", "operation failed; consult the server event log"
}
