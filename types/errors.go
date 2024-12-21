package types

// This gets returned when any errors related to token config
type TokenNotFoundError struct{}

func (t *TokenNotFoundError) Error() string {
	return "token.json file not found"
}

// This gets returned when any type of http exception occurs
type HttpError struct {
	StatusCode int
}

func (t *HttpError) Error() string {
	return "a http exception occured"
}
