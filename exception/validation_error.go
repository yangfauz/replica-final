package exception

type ValidationError struct {
	Message string
}

func (validation ValidationError) Error() string {
	return validation.Message
}
