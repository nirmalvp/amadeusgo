package referencedata

import (
	"encoding/json"

	"github.com/nirmalvp/amadeusgo/api/interfaces"
	"github.com/nirmalvp/amadeusgo/api/params"
	"github.com/nirmalvp/amadeusgo/api/response"
)

type locations struct {
	LocationsRepository interfaces.AmadeusRepository
}

func (locations *locations) Get() (int, response.Locations, error) {
	return locations.GetWithParams(nil)

}

func (locations *locations) GetWithParams(params params.Params) (int, response.Locations, error) {
	var formatedResponse response.Locations
	statusCode, responseBody, err := locations.LocationsRepository.Get(params)
	if err != nil {
		return 0, response.Locations{}, err
	}
	err = json.Unmarshal(responseBody, &formatedResponse)
	return statusCode, formatedResponse, err
}

func NewLocations(LocationsRepository interfaces.AmadeusRepository) *locations {
	return &locations{LocationsRepository}
}
