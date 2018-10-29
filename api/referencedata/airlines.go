package referencedata

import (
	"encoding/json"

	"github.com/nirmalvp/amadeusgo/api/interfaces"
	"github.com/nirmalvp/amadeusgo/api/params"
	"github.com/nirmalvp/amadeusgo/api/response"
)

type airlines struct {
	AirlinesRepository interfaces.AmadeusRepository
}

func (airlines *airlines) Get() (int, response.Airlines, error) {
	return airlines.GetWithParams(nil)

}

func (airlines *airlines) GetWithParams(params params.Params) (int, response.Airlines, error) {
	var formatedResponse response.Airlines
	statusCode, responseBody, err := airlines.AirlinesRepository.Get(params)
	if err != nil {
		return 0, formatedResponse, err
	}
	err = json.Unmarshal(responseBody, &formatedResponse)
	return statusCode, formatedResponse, err
}

func NewAirlines(AirlinesRepository interfaces.AmadeusRepository) *airlines {
	return &airlines{AirlinesRepository}
}
