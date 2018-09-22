package model

const (
	//OfflineScope scope value to request refresh token
	OfflineScope = "offline"
	//RefrestTokenType refresh token type value
	RefrestTokenType = "refresh"
)

//TokenService manage tokens abstraction layer
type TokenService interface {
	//NewToken creates new access token for the user
	NewToken(u User, scopes []string, app AppData) (Token, error)
	//NewRefreshToken creates new refresh token for the user
	NewRefreshToken(u User, scopes []string, app AppData) (Token, error)
	//RefreshToken issues the new access token with access token
	RefreshToken(token Token) (Token, error)
	Parse(string) (Token, error)
	String(Token) (string, error)
}

//Token is app token to give user chan
type Token interface {
	Validate() error
}

//Validator calidate token with external requester
type Validator interface {
	Validate(Token) error
}

//TokenMapping is service to match tokens to services. etc
type TokenMapping interface {
}