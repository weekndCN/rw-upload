package render

var (
	// ErrNotFound error not found
	ErrNotFound = New("No File Found")
	// ErrOverLimit .
	ErrOverLimit = New("File Size Over Than Limit")
	// ErrEmptyName .
	ErrEmptyName = New("File Name Empty")
	// ErrNotImplemented .
	ErrNotImplemented = New("Not Implemented")
	// ErrFormatSupport .
	ErrFormatSupport = New("Format not support")
)

// Error represents a json-encoded API error
type Error struct {
	Message string `json:"msg"`
}

//  Error message
func (e *Error) Error() string {
	return e.Message
}

// New returns a new error message.
func New(text string) error {
	return &Error{Message: text}
}
