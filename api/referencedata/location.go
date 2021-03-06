package referencedata

import (
	"fmt"

	"github.com/nirmalvp/amadeusgo/api/interfaces"
	"github.com/nirmalvp/amadeusgo/api/params"
	"github.com/nirmalvp/amadeusgo/api/request"
	"github.com/nirmalvp/amadeusgo/api/response"
	"github.com/nirmalvp/amadeusgo/api/service"
)

type location struct {
	PathUrl                     string
	RestClient                  interfaces.AmadeusRest
	AuthenticatedRequestCreator *service.AuthenticatedRequestCreator
}

func (location *location) Get() (response.Location, error) {
	return location.GetWithParams(nil)

}

func (location *location) GetWithParams(params params.Params) (response.Location, error) {
	request, authenticationErr := location.AuthenticatedRequestCreator.Create(request.GET, location.PathUrl, params)
	if authenticationErr != nil {
		return response.Location{}, authenticationErr
	}
	statusCode, responseBody, err := location.RestClient.Send(request)
	if err != nil {
		return response.Location{}, err
	}
	return response.NewLocationResponse(statusCode, responseBody, request), nil
}

func NewLocation(locationId string, restClient interfaces.AmadeusRest, authenticatedRequestCreator *service.AuthenticatedRequestCreator) *location {
	return &location{
		PathUrl:                     fmt.Sprintf("/v1/reference-data/locations/%s", locationId),
		RestClient:                  restClient,
		AuthenticatedRequestCreator: authenticatedRequestCreator,
	}
}
