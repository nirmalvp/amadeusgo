package api

import (
	"fmt"

	"github.com/nirmalvp/amadeusgo/api/params"
	"github.com/nirmalvp/amadeusgo/api/types"
)

type AmadeusHTTPRepository struct {
	AccessToken *AccessToken
	URLPaths    map[types.Action]string
	Configuration
}

func (amRepo AmadeusHTTPRepository) Get(params params.Params) (int, []byte, error) {
	if _, ok := amRepo.URLPaths[types.READ]; !ok {
		return 0, nil, fmt.Errorf("Action not supported")
	}
	urlPath := amRepo.URLPaths[types.READ]
	request := NewAmadeusRequest(amRepo.useSSL, amRepo.host, amRepo.port, urlPath)
	request = request.
		Get(params).
		SetHeaders(amRepo.accept, amRepo.clientVersion, amRepo.languageVersion)
	if amRepo.customAppId != nil && amRepo.customAppVersion != nil {
		request = request.SetAppID(*amRepo.customAppId, *amRepo.customAppVersion)
	}
	if amRepo.AccessToken != nil {
		bearerToken, err := amRepo.AccessToken.getBearerToken(amRepo.clientId, amRepo.clientSecret)
		if err != nil {
			return 0, nil, err
		}
		request = request.SetAuth(bearerToken)
	}
	return request.Do()
}

func (amRepo AmadeusHTTPRepository) Create(params params.Params) (int, []byte, error) {
	if _, ok := amRepo.URLPaths[types.CREATE]; !ok {
		return 0, nil, fmt.Errorf("Action not supported")
	}
	urlPath := amRepo.URLPaths[types.CREATE]
	request := NewAmadeusRequest(amRepo.useSSL, amRepo.host, amRepo.port, urlPath)
	request.
		Post(params).
		SetHeaders(amRepo.accept, amRepo.clientVersion, amRepo.languageVersion)
	if amRepo.customAppId != nil && amRepo.customAppVersion != nil {
		request = request.SetAppID(*amRepo.customAppId, *amRepo.customAppVersion)
	}
	if amRepo.AccessToken != nil {
		bearerToken, err := amRepo.AccessToken.getBearerToken(amRepo.clientId, amRepo.clientSecret)
		if err != nil {
			return 0, nil, err
		}
		request = request.SetAuth(bearerToken)
	}
	return request.Do()
}
