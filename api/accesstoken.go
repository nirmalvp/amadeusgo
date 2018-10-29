package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nirmalvp/amadeusgo/api/interfaces"
	"github.com/nirmalvp/amadeusgo/api/params"
)

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
	expiresIn   int    `json:"expires_in"`
}

type AccessToken struct {
	AccessTokenRepository interfaces.AmadeusRepository
	TokenBufferTime       time.Duration
	AccessToken           *string
	expiresAt             *time.Time
}

func (at *AccessToken) getBearerToken(clientId, clientSecret string) (string, error) {
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
func (at *AccessToken) needsRefresh() bool {
	if at.AccessToken == nil {
		return true
	}
	return time.Now().Add(at.TokenBufferTime).After(*at.expiresAt)
}

func (at *AccessToken) getUpdatedAccessToken(clientId, clientSecret string) (*string, *time.Time, error) {
	_, responseBody, err := at.fetchAccessToken(clientId, clientSecret)
	if err != nil {
		return nil, nil, err
	}
	accessToken, expiresIn := at.parseAccessToken(responseBody)
	return &accessToken, &expiresIn, nil
}
func (at *AccessToken) fetchAccessToken(clientId, clientSecret string) (int, []byte, error) {
	params := params.
		With("grant_type", "client_credentials").
		And("client_id", clientId).
		And("client_secret", clientSecret)
	return at.AccessTokenRepository.Create(params)
}

func (at *AccessToken) parseAccessToken(responseStr []byte) (string, time.Time) {
	responseObj := new(accessTokenResponse)
	json.Unmarshal(responseStr, responseObj)
	accessToken := responseObj.AccessToken
	expiresIn := responseObj.expiresIn
	expiresAt := time.Now().Add(time.Duration(expiresIn) * time.Second)
	return accessToken, expiresAt
}
