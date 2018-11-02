package service

import (
	"fmt"

	"github.com/nirmalvp/amadeusgo/api/params"
	"github.com/nirmalvp/amadeusgo/api/request"
	"github.com/nirmalvp/amadeusgo/api/response"
)

type AuthenticatedRequestCreator struct {
	clientId           string
	clientSecret       string
	scheme             string
	host               string
	clientVersion      string
	languageVersion    string
	appId              *string
	appVersion         *string
	ssl                bool
	port               int
	userAgent          string
	accept             string
	accessTokenService *accessTokenService
}

func NewAuthenticatedRequestCreator(configuration Configuration, accessTokenService *accessTokenService) *AuthenticatedRequestCreator {
	arc := AuthenticatedRequestCreator{
		clientId:           configuration.clientId,
		clientSecret:       configuration.clientSecret,
		host:               configuration.host,
		languageVersion:    configuration.languageVersion,
		clientVersion:      configuration.clientVersion,
		appId:              configuration.appId,
		appVersion:         configuration.appVersion,
		port:               configuration.port,
		ssl:                configuration.ssl,
		accept:             configuration.accept,
		accessTokenService: accessTokenService,
	}
	if arc.ssl {
		arc.scheme = "https"
	} else {
		arc.scheme = "http"
	}
	arc.userAgent = fmt.Sprintf("amadeus-go/%s go/%s", arc.clientVersion, arc.languageVersion)
	if arc.appId != nil && arc.appVersion != nil {
		arc.userAgent = arc.userAgent + fmt.Sprintf(" %s/%s", arc.appId, arc.appVersion)
	}
	return &arc
}

func (authenticatedRequestCreator *AuthenticatedRequestCreator) Create(verb request.Verb, pathUrl string, params params.Params) (request.AmadeusRequestData, error) {
	requestData := request.AmadeusRequestData{
		Verb:            verb,
		Host:            authenticatedRequestCreator.host,
		Path:            pathUrl,
		Params:          params,
		LanguageVersion: authenticatedRequestCreator.languageVersion,
		ClientVersion:   authenticatedRequestCreator.clientVersion,
		AppId:           authenticatedRequestCreator.appId,
		AppVersion:      authenticatedRequestCreator.appVersion,
		Port:            authenticatedRequestCreator.port,
		SSL:             authenticatedRequestCreator.ssl,
		Scheme:          authenticatedRequestCreator.scheme,
		UserAgent:       authenticatedRequestCreator.userAgent,
		Accept:          authenticatedRequestCreator.accept,
	}
	requestData.URI = fmt.Sprintf("%s://%s:%d%s", requestData.Scheme, requestData.Host, requestData.Port, requestData.Path)
	requestData.Headers = make(request.Header)
	requestData.Headers["User-Agent"] = requestData.UserAgent
	requestData.Headers["Accept"] = requestData.Accept
	bearerToken, err := authenticatedRequestCreator.accessTokenService.getBearerToken(
		authenticatedRequestCreator.clientId,
		authenticatedRequestCreator.clientSecret)
	if err != nil {
		if respErr, ok := err.(response.ResponseError); ok {
			respErr.AmadeusResponse.Request = requestData
			return request.AmadeusRequestData{}, respErr
		}
		return request.AmadeusRequestData{}, err
	}
	requestData.BearerToken = &bearerToken
	requestData.Headers["Authorization"] = *requestData.BearerToken
	return requestData, nil
}
