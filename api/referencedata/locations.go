package referencedata

import (
	"github.com/nirmalvp/amadeusgo/api/interfaces"
	"github.com/nirmalvp/amadeusgo/api/params"
	"github.com/nirmalvp/amadeusgo/api/request"
	"github.com/nirmalvp/amadeusgo/api/response"
	"github.com/nirmalvp/amadeusgo/api/service"
)

type locations struct {
	PathUrl                     string
	RestClient                  interfaces.AmadeusRest
	AuthenticatedRequestCreator *service.AuthenticatedRequestCreator
	Airports                    *airports
}

func (locations *locations) Get() (response.Locations, error) {
	return locations.GetWithParams(nil)

}

func (locations *locations) GetWithParams(params params.Params) (response.Locations, error) {
	request, authenticationErr := locations.AuthenticatedRequestCreator.Create(request.GET, locations.PathUrl, params)
	if authenticationErr != nil {
		return response.Locations{}, authenticationErr
	}
	statusCode, responseBody, err := locations.RestClient.Send(request)
	if err != nil {
		return response.Locations{}, err
	}
	return response.NewLocationsResponse(statusCode, responseBody, request), nil
}

func NewLocations(restClient interfaces.AmadeusRest, authenticatedRequestCreator *service.AuthenticatedRequestCreator, airports *airports) *locations {
	return &locations{
		PathUrl:                     "/v1/reference-data/locations",
		RestClient:                  restClient,
		AuthenticatedRequestCreator: authenticatedRequestCreator,
		Airports:                    airports,
	}
}
