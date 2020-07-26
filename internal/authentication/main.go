package authentication

import (
	"fmt"
	"time"

	cognitosrp "github.com/alexrudd/cognito-srp/v2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cip "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/highwingio/incognito/internal/types"
)

// CognitoAuthenticator manages attempting to authenticate a user against a specified Cognito User Pool
type CognitoAuthenticator struct {
	Credentials *types.CognitoAuthenticationCredentials
}

// NewCognitoAuthenticator returns a new instance of a CognitoAuthenticator with the given parameters
func NewCognitoAuthenticator(
	credentials *types.CognitoAuthenticationCredentials,
) *CognitoAuthenticator {
	return &CognitoAuthenticator{
		Credentials: credentials,
	}
}

// SignIn attempts to authenticate the user against the specified User Pool
func (c *CognitoAuthenticator) SignIn() (*cip.AuthenticationResultType, error) {

	// Attempt to sign in
	sess := session.Must(session.NewSession())
	svc := cip.New(sess)
	csrp, _ := cognitosrp.NewCognitoSRP(
		c.Credentials.Username,
		c.Credentials.Password,
		c.Credentials.UserPoolID,
		c.Credentials.ClientID,
		nil,
	)

	req, resp := svc.InitiateAuthRequest(&cip.InitiateAuthInput{
		AuthFlow:       aws.String(cip.AuthFlowTypeUserSrpAuth),
		ClientId:       aws.String(csrp.GetClientId()),
		AuthParameters: aws.StringMap(csrp.GetAuthParams()),
	})
	err := req.Send()
	if err != nil {
		return nil, err
	}

	if *resp.ChallengeName != cip.ChallengeNameTypePasswordVerifier {
		return nil, fmt.Errorf("Unrecognized challenge name %q", aws.StringValue(resp.ChallengeName))
	}

	challengeInput, err := csrp.PasswordVerifierChallenge(aws.StringValueMap(resp.ChallengeParameters), time.Now())
	if err != nil {
		return nil, err
	}
	chal, chalResp := svc.RespondToAuthChallengeRequest(&cip.RespondToAuthChallengeInput{
		ChallengeName:      resp.ChallengeName,
		ChallengeResponses: aws.StringMap(challengeInput),
		ClientId:           aws.String(csrp.GetClientId()),
	})
	err = chal.Send()
	if err != nil {
		return nil, err
	}
	return chalResp.AuthenticationResult, nil
}
