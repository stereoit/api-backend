package domain

// DuplicateError error
type DuplicateError struct{}

func (err *DuplicateError) Error() string {
	return "Duplicate Entry"
}
