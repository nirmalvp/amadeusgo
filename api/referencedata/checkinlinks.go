package referencedata

import (
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

func (checkinLinks *checkinLinks) Get() (response.CheckinLinks, error) {
	return checkinLinks.GetWithParams(nil)

}

func (checkinLinks *checkinLinks) GetWithParams(params params.Params) (response.CheckinLinks, error) {
	request, authenticationErr := checkinLinks.AuthenticatedRequestCreator.Create(request.GET, checkinLinks.PathUrl, params)
	if authenticationErr != nil {
		return response.CheckinLinks{}, authenticationErr
	}
	statusCode, responseBody, err := checkinLinks.RestClient.Send(request)
	if err != nil {
		return response.CheckinLinks{}, err
	}
	return response.NewCheckinLinksResponse(statusCode, responseBody, request), nil
}

func NewCheckinLinks(restClient interfaces.AmadeusRest, authenticatedRequestCreator *service.AuthenticatedRequestCreator) *checkinLinks {
	return &checkinLinks{
		PathUrl:                     "/v2/reference-data/urls/checkin-links",
		RestClient:                  restClient,
		AuthenticatedRequestCreator: authenticatedRequestCreator,
	}
}
