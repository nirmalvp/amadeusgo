package service

import (
	"fmt"

	"github.com/nirmalvp/amadeusgo/api/params"
	"github.com/nirmalvp/amadeusgo/api/request"
)

type UnAuthenticatedRequestCreator struct {
	scheme          string
	host            string
	clientVersion   string
	languageVersion string
	appId           *string
	appVersion      *string
	port            int
	userAgent       string
	accept          string
}

func NewUnAuthenticatedRequestCreator(configuration Configuration) *UnAuthenticatedRequestCreator {
	aurc := UnAuthenticatedRequestCreator{
		host:            configuration.host,
		accept:          configuration.accept,
		languageVersion: configuration.languageVersion,
		clientVersion:   configuration.clientVersion,
		port:            configuration.port,
	}
	if configuration.ssl {
		aurc.scheme = "https"
	} else {
		aurc.scheme = "http"
	}
	aurc.userAgent = fmt.Sprintf("amadeus-go/%s go/%s", aurc.clientVersion, aurc.languageVersion)
	if configuration.appId != nil && configuration.appVersion != nil {
		aurc.userAgent = aurc.userAgent + fmt.Sprintf(" %s/%s", *configuration.appId, *configuration.appVersion)
	}
	return &aurc
}

func (aurc *UnAuthenticatedRequestCreator) Create(verb request.Verb, pathUrl string, params params.Params) request.AmadeusRequestData {
	requestData := request.AmadeusRequestData{
		Verb:   verb,
		Params: params,
	}
	requestData.URI = fmt.Sprintf("%s://%s:%d%s", aurc.scheme, aurc.host, aurc.port, pathUrl)
	requestData.Headers = make(request.Header)
	requestData.Headers["User-Agent"] = aurc.userAgent
	requestData.Headers["Accept"] = aurc.accept
	return requestData
}
