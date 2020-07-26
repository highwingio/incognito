package types

// CognitoAuthenticationCredentials represents a set of credentials for authenticating against an AWS Cognito User Pool
type CognitoAuthenticationCredentials struct {
	Username   string
	Password   string
	UserPoolID string
	ClientID   string
}
