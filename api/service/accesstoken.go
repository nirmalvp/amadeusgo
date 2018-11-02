package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nirmalvp/amadeusgo/api/interfaces"
	"github.com/nirmalvp/amadeusgo/api/params"
	"github.com/nirmalvp/amadeusgo/api/request"
)

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
	expiresIn   int    `json:"expires_in"`
}

type accessTokenService struct {
	pathUrl                       string
	restClient                    interfaces.AmadeusRest
	unAuthenticatedRequestCreator *UnAuthenticatedRequestCreator
	TokenBufferTime               time.Duration
	AccessToken                   *string
	expiresAt                     *time.Time
}

func (at *accessTokenService) getBearerToken(clientId, clientSecret string) (string, error) {
	if at.needsRefresh() {
		accessToken, expiresIn, err := at.getUpdatedAccessToken(clientId, clientSecret)
		if err != nil {
			return "", err
		}
		at.AccessToken = accessToken
		at.expiresAt = expiresIn
	}
	return fmt.Sprintf("Bearer %s", *at.AccessToken), nil
}

// Not threadsafe. TODO : Implement mutex
func (at *accessTokenService) needsRefresh() bool {
	if at.AccessToken == nil {
		return true
	}
	return time.Now().Add(at.TokenBufferTime).After(*at.expiresAt)
}

func (at *accessTokenService) getUpdatedAccessToken(clientId, clientSecret string) (*string, *time.Time, error) {
	_, responseBody, err := at.fetchAccessToken(clientId, clientSecret)
	if err != nil {
		return nil, nil, err
	}
	accessToken, expiresIn := at.parseAccessToken(responseBody)
	return &accessToken, &expiresIn, nil
}
func (at *accessTokenService) fetchAccessToken(clientId, clientSecret string) (int, []byte, error) {
	params := params.
		With("grant_type", "client_credentials").
		And("client_id", clientId).
		And("client_secret", clientSecret)
	accessTokenRequest := at.unAuthenticatedRequestCreator.Create(request.POST, at.pathUrl, params)
	return at.restClient.Send(accessTokenRequest)
}

func (at *accessTokenService) parseAccessToken(responseStr []byte) (string, time.Time) {
	responseObj := new(accessTokenResponse)
	json.Unmarshal(responseStr, responseObj)
	accessToken := responseObj.AccessToken
	expiresIn := responseObj.expiresIn
	expiresAt := time.Now().Add(time.Duration(expiresIn) * time.Second)
	return accessToken, expiresAt
}

func NewAccessTokenService(restClient interfaces.AmadeusRest, unAuthenticatedRequestCreator *UnAuthenticatedRequestCreator, tokenBufferTime time.Duration) *accessTokenService {
	return &accessTokenService{
		pathUrl:                       "/v1/security/oauth2/token",
		restClient:                    restClient,
		unAuthenticatedRequestCreator: unAuthenticatedRequestCreator,
		TokenBufferTime:               tokenBufferTime,
	}
}
