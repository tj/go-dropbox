package dropbox

// Error response.
type Error struct {
	Status     string
	StatusCode int
	Summary    string `json:"error_summary"`
}

// Error string.
func (e *Error) Error() string {
	return e.Summary
}
