package referencedata

import (
	"encoding/json"

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
}

func (locations *locations) Get() (int, response.Locations, error) {
	return locations.GetWithParams(nil)

}

func (locations *locations) GetWithParams(params params.Params) (int, response.Locations, error) {
	request, authenticationErr := locations.AuthenticatedRequestCreator.Create(request.GET, locations.PathUrl, params)
	if authenticationErr != nil {
		return 0, response.Locations{}, authenticationErr
	}
	statusCode, responseBody, err := locations.RestClient.Send(request)
	if err != nil {
		return 0, response.Locations{}, err
	}
	var formatedRestResponse response.LocationsRest
	err = json.Unmarshal(responseBody, &formatedRestResponse)
	formatedClientResponse := response.NewLocationsResponse(statusCode, formatedRestResponse, request, err == nil)
	return statusCode, formatedClientResponse, err
}

func NewLocations(restClient interfaces.AmadeusRest, authenticatedRequestCreator *service.AuthenticatedRequestCreator) *locations {
	return &locations{
		PathUrl:                     "/v1/reference-data/locations",
		RestClient:                  restClient,
		AuthenticatedRequestCreator: authenticatedRequestCreator,
	}
}
