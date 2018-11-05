package referencedata

import (
	"github.com/nirmalvp/amadeusgo/api/interfaces"
	"github.com/nirmalvp/amadeusgo/api/params"
	"github.com/nirmalvp/amadeusgo/api/request"
	"github.com/nirmalvp/amadeusgo/api/response"
	"github.com/nirmalvp/amadeusgo/api/service"
)

type airports struct {
	PathUrl                     string
	RestClient                  interfaces.AmadeusRest
	AuthenticatedRequestCreator *service.AuthenticatedRequestCreator
}

func (airports *airports) Get() (response.Airports, error) {
	return airports.GetWithParams(nil)

}

func (airports *airports) GetWithParams(params params.Params) (response.Airports, error) {
	request, authenticationErr := airports.AuthenticatedRequestCreator.Create(request.GET, airports.PathUrl, params)
	if authenticationErr != nil {
		return response.Airports{}, authenticationErr
	}
	statusCode, responseBody, err := airports.RestClient.Send(request)
	if err != nil {
		return response.Airports{}, err
	}
	return response.NewAirportsResponse(statusCode, responseBody, request), nil
}

func NewAirports(restClient interfaces.AmadeusRest, authenticatedRequestCreator *service.AuthenticatedRequestCreator) *airports {
	return &airports{
		PathUrl:                     "/v1/reference-data/locations/airports",
		RestClient:                  restClient,
		AuthenticatedRequestCreator: authenticatedRequestCreator,
	}
}
