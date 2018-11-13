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
	ExpiresIn   int    `json:"expires_in"`
}

// accessTokenService is service that when provided A client Id and client Response uses the rest client to
// fetch the access token
type accessTokenService struct {
	pathUrl                       string
	restClient                    interfaces.AmadeusRest
	unAuthenticatedRequestCreator *UnAuthenticatedRequestCreator
	tokenBufferTime               time.Duration
	accessToken                   *string
	expiresAt                     *time.Time
	timeGetter                    interfaces.TimeGetter
}

func (at *accessTokenService) getBearerToken(clientId, clientSecret string) (string, error) {
	if at.needsRefresh() {
		accessToken, expiresAt, err := at.getUpdatedAccessToken(clientId, clientSecret)
		if err != nil {
			return "", err
		}
		at.accessToken = accessToken
		at.expiresAt = expiresAt
	}
	return fmt.Sprintf("Bearer %s", *at.accessToken), nil
}

// Not threadsafe. TODO : Implement mutex
func (at *accessTokenService) needsRefresh() bool {
	if at.accessToken == nil {
		return true
	}
	return at.timeGetter.GetCurrentTime().Add(at.tokenBufferTime).After(*at.expiresAt)
}

func (at *accessTokenService) getUpdatedAccessToken(clientId, clientSecret string) (*string, *time.Time, error) {
	_, responseBody, err := at.fetchAccessToken(clientId, clientSecret)
	if err != nil {
		return nil, nil, err
	}
	accessToken, expiresAt := at.parseAccessToken(responseBody)
	return &accessToken, &expiresAt, nil
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
	expiresIn := responseObj.ExpiresIn
	expiresAt := at.timeGetter.GetCurrentTime().Add(time.Duration(expiresIn) * time.Second)
	return accessToken, expiresAt
}

func NewAccessTokenService(restClient interfaces.AmadeusRest,
	unAuthenticatedRequestCreator *UnAuthenticatedRequestCreator,
	tokenBufferTime time.Duration,
	timeGetter interfaces.TimeGetter) *accessTokenService {
	return &accessTokenService{
		pathUrl:                       "/v1/security/oauth2/token",
		restClient:                    restClient,
		unAuthenticatedRequestCreator: unAuthenticatedRequestCreator,
		tokenBufferTime:               tokenBufferTime,
		timeGetter:                    timeGetter,
	}
}
