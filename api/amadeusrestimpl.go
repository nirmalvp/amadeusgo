package api

import (
	"fmt"

	"github.com/nirmalvp/amadeusgo/api/request"
	"github.com/nirmalvp/amadeusgo/api/response"
	"github.com/parnurzeal/gorequest"
)

type AmadeusRequester struct {
	requester *gorequest.SuperAgent
}

type AmadeusRestClient struct{}

func (amRepo *AmadeusRestClient) Send(amRequest request.AmadeusRequestData) (int, []byte, error) {
	amadeusRequester := AmadeusRequester{gorequest.New()}
	switch amRequest.Verb {
	case request.GET:
		amadeusRequester.requester.Get(amRequest.URI)
		for key, value := range amRequest.Params {
			amadeusRequester.requester.Param(key, value)
		}
	case request.POST:
		amadeusRequester.requester.Post(amRequest.URI)
		for key, value := range amRequest.Params {
			amadeusRequester.requester.Send(fmt.Sprintf("%s=%s", key, value))
		}
	}
	for key, val := range amRequest.Headers {
		amadeusRequester.requester.Set(key, val)
	}
	fmt.Println(amadeusRequester.requester.AsCurlCommand())
	resp, body, errs := amadeusRequester.requester.End()
	if len(errs) > 0 {
		err := errs[0]
		return 0, nil, err
	}
	if resp == nil {
		err := fmt.Errorf("Unable to Connect")
		return 0, nil, err
	}
	httpresponse := make([]byte, len(body))
	httpresponse = []byte(body)
	statusCode := resp.StatusCode
	if statusCode >= 400 {
		return 0, nil, response.NewResponseError(statusCode, httpresponse, amRequest)
	}
	return statusCode, httpresponse, nil
}
