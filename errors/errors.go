package errors

type ErrNoRows struct{ message string }
type ErrDbConn struct{ message string }

// NewErrNoRows ...
func NewErrNoRows(msg string) *ErrNoRows { return &ErrNoRows{msg} }
func (e *ErrNoRows) Error() string       { return e.message }

// NewErrNoRows ...
func NewErrDbConn(msg string) *ErrDbConn { return &ErrDbConn{msg} }
func (e *ErrDbConn) Error() string       { return e.message }
