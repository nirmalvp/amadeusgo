package service

import (
	"fmt"

	"github.com/nirmalvp/amadeusgo/api/params"
	"github.com/nirmalvp/amadeusgo/api/request"
)

type UnAuthenticatedRequestCreator struct {
	clientId        string
	clientSecret    string
	scheme          string
	host            string
	clientVersion   string
	languageVersion string
	appId           *string
	appVersion      *string
	ssl             bool
	port            int
	userAgent       string
	accept          string
}

func NewUnAuthenticatedRequestCreator(configuration Configuration) *UnAuthenticatedRequestCreator {
	aurc := UnAuthenticatedRequestCreator{
		clientId:        configuration.clientId,
		clientSecret:    configuration.clientSecret,
		host:            configuration.host,
		languageVersion: configuration.languageVersion,
		clientVersion:   configuration.clientVersion,
		appId:           configuration.appId,
		appVersion:      configuration.appVersion,
		port:            configuration.port,
		ssl:             configuration.ssl,
		accept:          configuration.accept,
	}
	if aurc.ssl {
		aurc.scheme = "https"
	} else {
		aurc.scheme = "http"
	}
	aurc.userAgent = fmt.Sprintf("amadeus-go/%s go/%s", aurc.clientVersion, aurc.languageVersion)
	if aurc.appId != nil && aurc.appVersion != nil {
		aurc.userAgent = aurc.userAgent + fmt.Sprintf(" %s/%s", aurc.appId, aurc.appVersion)
	}
	return &aurc
}

func (aurc *UnAuthenticatedRequestCreator) Create(verb request.Verb, pathUrl string, params params.Params) request.AmadeusRequestData {
	requestData := request.AmadeusRequestData{
		Verb:            verb,
		Host:            aurc.host,
		Path:            pathUrl,
		Params:          params,
		LanguageVersion: aurc.languageVersion,
		ClientVersion:   aurc.clientVersion,
		AppId:           aurc.appId,
		AppVersion:      aurc.appVersion,
		Port:            aurc.port,
		SSL:             aurc.ssl,
		Scheme:          aurc.scheme,
		UserAgent:       aurc.userAgent,
		Accept:          aurc.accept,
	}
	requestData.URI = fmt.Sprintf("%s://%s:%d%s", requestData.Scheme, requestData.Host, requestData.Port, requestData.Path)
	requestData.Headers = make(request.Header)
	requestData.Headers["User-Agent"] = requestData.UserAgent
	requestData.Headers["Accept"] = requestData.Accept
	return requestData
}
