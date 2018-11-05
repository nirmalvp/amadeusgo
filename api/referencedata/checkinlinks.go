package referencedata

import (
	"encoding/json"

	"github.com/nirmalvp/amadeusgo/api/interfaces"
	"github.com/nirmalvp/amadeusgo/api/params"
	"github.com/nirmalvp/amadeusgo/api/request"
	"github.com/nirmalvp/amadeusgo/api/response"
	"github.com/nirmalvp/amadeusgo/api/service"
)

type checkinLinks struct {
	PathUrl                     string
	RestClient                  interfaces.AmadeusRest
	AuthenticatedRequestCreator *service.AuthenticatedRequestCreator
}

func (checkinLinks *checkinLinks) Get() (int, response.CheckinLinks, error) {
	return checkinLinks.GetWithParams(nil)

}

func (checkinLinks *checkinLinks) GetWithParams(params params.Params) (int, response.CheckinLinks, error) {
	request, authenticationErr := checkinLinks.AuthenticatedRequestCreator.Create(request.GET, checkinLinks.PathUrl, params)
	if authenticationErr != nil {
		return 0, response.CheckinLinks{}, authenticationErr
	}
	statusCode, responseBody, err := checkinLinks.RestClient.Send(request)
	if err != nil {
		return 0, response.CheckinLinks{}, err
	}
	var formatedRestResponse response.CheckinLinksRest
	err = json.Unmarshal(responseBody, &formatedRestResponse)
	formatedClientResponse := response.NewCheckinLinksResponse(statusCode, formatedRestResponse, request, err == nil)
	return statusCode, formatedClientResponse, err
}

func NewCheckinLinks(restClient interfaces.AmadeusRest, authenticatedRequestCreator *service.AuthenticatedRequestCreator) *checkinLinks {
	return &checkinLinks{
		PathUrl:                     "/v2/reference-data/urls/checkin-links",
		RestClient:                  restClient,
		AuthenticatedRequestCreator: authenticatedRequestCreator,
	}
}
