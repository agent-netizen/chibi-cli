package types

type TokenNotFoundError struct {}

func (t *TokenNotFoundError) Error() string {
	return "token.json file not found"
}

