package api

import (
	"encoding/json"
	"fmt"

	"github.com/nirmalvp/amadeusgo/api/errors"
	"github.com/nirmalvp/amadeusgo/api/params"
	"github.com/parnurzeal/gorequest"
)

type AmadeusRequest struct {
	requester   *gorequest.SuperAgent
	scheme      string
	host        string
	port        int
	accept      string
	userAgent   string
	fullUrl     string
	bearerToken string
}

func (amRequest AmadeusRequest) SetAppID(appId, appVersion string) AmadeusRequest {
	userAgent := amRequest.userAgent + fmt.Sprintf(" %s/%s", appId, appVersion)
	amRequest.userAgent = userAgent
	amRequest.requester.Set("User-Agent", userAgent)
	return amRequest
}

func (amRequest AmadeusRequest) SetHeaders(accept, clientVersion, languageVersion string) AmadeusRequest {
	amRequest.accept = accept
	amRequest.requester.Set("Accept", accept)
	userAgent := fmt.Sprintf("amadeusgo/%s go/%s", clientVersion, languageVersion)
	amRequest.userAgent = userAgent
	amRequest.requester.Set("User-Agent", userAgent)
	return amRequest
}

func (amRequest AmadeusRequest) SetAuth(bearerToken string) AmadeusRequest {
	amRequest.bearerToken = bearerToken
	amRequest.requester.Set("Authorization", bearerToken)
	return amRequest
}

func (amRequest AmadeusRequest) Get(params params.Params) AmadeusRequest {
	amRequest.requester.Get(amRequest.fullUrl)
	for key, value := range params {
		amRequest.requester.Param(key, value)
	}
	return amRequest
}

func (amRequest AmadeusRequest) Post(params params.Params) AmadeusRequest {
	amRequest.requester.Post(amRequest.fullUrl)
	for key, value := range params {
		amRequest.requester.Send(fmt.Sprintf("%s=%s", key, value))
	}
	return amRequest
}

func (amRequest AmadeusRequest) Do() (int, []byte, error) {

	resp, body, errs := amRequest.requester.End()

	if len(errs) > 0 {
		err := errs[0]
		return 0, nil, err
	}
	if resp == nil {
		err := fmt.Errorf("Unable to Connect")
		return 0, nil, err
	}
	response := make([]byte, len(body))
	response = []byte(body)
	statusCode := resp.StatusCode
	if statusCode >= 400 {
		amadeusError := errors.AmadeusError{}
		_ = json.Unmarshal(response, &amadeusError)
		return 0, nil, amadeusError
	}
	return statusCode, response, nil
}

func NewAmadeusRequest(useSSL bool, host string, port int, urlPath string) AmadeusRequest {
	amRequest := AmadeusRequest{requester: gorequest.New(),
		host: host,
		port: port,
	}
	if useSSL {
		amRequest.scheme = "https"
	} else {
		amRequest.scheme = "http"
	}
	amRequest.fullUrl = fmt.Sprintf("%s://%s:%d%s", amRequest.scheme, amRequest.host, amRequest.port, urlPath)
	return amRequest
}
