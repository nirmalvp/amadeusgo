package referencedata

import (
	"github.com/nirmalvp/amadeusgo/api/interfaces"
	"github.com/nirmalvp/amadeusgo/api/params"
	"github.com/nirmalvp/amadeusgo/api/request"
	"github.com/nirmalvp/amadeusgo/api/response"
	"github.com/nirmalvp/amadeusgo/api/service"
)

type airlines struct {
	PathUrl                     string
	RestClient                  interfaces.AmadeusRest
	AuthenticatedRequestCreator *service.AuthenticatedRequestCreator
}

func (airlines *airlines) Get() (response.Airlines, error) {
	return airlines.GetWithParams(nil)

}

func (airlines *airlines) GetWithParams(params params.Params) (response.Airlines, error) {
	request, authenticationErr := airlines.AuthenticatedRequestCreator.Create(request.GET, airlines.PathUrl, params)
	if authenticationErr != nil {
		return response.Airlines{}, authenticationErr
	}
	statusCode, responseBody, err := airlines.RestClient.Send(request)
	if err != nil {
		return response.Airlines{}, err
	}
	return response.NewAirlineResponse(statusCode, responseBody, request), nil
}

func NewAirlines(restClient interfaces.AmadeusRest, authenticatedRequestCreator *service.AuthenticatedRequestCreator) *airlines {
	return &airlines{
		PathUrl:                     "/v1/reference-data/airlines",
		RestClient:                  restClient,
		AuthenticatedRequestCreator: authenticatedRequestCreator,
	}
}
