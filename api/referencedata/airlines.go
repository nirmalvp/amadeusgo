package referencedata

import (
	"encoding/json"

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

func (airlines *airlines) Get() (int, response.Airlines, error) {
	return airlines.GetWithParams(nil)

}

func (airlines *airlines) GetWithParams(params params.Params) (int, response.Airlines, error) {
	request, authenticationErr := airlines.AuthenticatedRequestCreator.Create(request.GET, airlines.PathUrl, params)
	if authenticationErr != nil {
		return 0, response.Airlines{}, authenticationErr
	}
	statusCode, responseBody, err := airlines.RestClient.Send(request)
	if err != nil {
		return 0, response.Airlines{}, err
	}
	var formatedRestResponse response.AirlineRest
	err = json.Unmarshal(responseBody, &formatedRestResponse)
	formatedClientResponse := response.NewAirlineResponse(statusCode, formatedRestResponse, request, err == nil)
	return statusCode, formatedClientResponse, err
}

func NewAirlines(restClient interfaces.AmadeusRest, authenticatedRequestCreator *service.AuthenticatedRequestCreator) *airlines {
	return &airlines{
		PathUrl:                     "/v1/reference-data/airlines",
		RestClient:                  restClient,
		AuthenticatedRequestCreator: authenticatedRequestCreator,
	}
}
