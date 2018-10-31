package referencedata

import (
	"encoding/json"

	"github.com/nirmalvp/amadeusgo/api/interfaces"
	"github.com/nirmalvp/amadeusgo/api/params"
	"github.com/nirmalvp/amadeusgo/api/response"
)

type checkinLinks struct {
	CheckinlinksRepository interfaces.AmadeusRepository
}

func (checkinLinks *checkinLinks) Get() (int, response.CheckinLinks, error) {
	return checkinLinks.GetWithParams(nil)

}

func (checkinLinks *checkinLinks) GetWithParams(params params.Params) (int, response.CheckinLinks, error) {
	var formatedResponse response.CheckinLinks
	statusCode, responseBody, err := checkinLinks.CheckinlinksRepository.Get(params)
	if err != nil {
		return 0, response.CheckinLinks{}, err
	}
	err = json.Unmarshal(responseBody, &formatedResponse)
	return statusCode, formatedResponse, err
}

func NewCheckinLinks(checkinLinksRepository interfaces.AmadeusRepository) *checkinLinks {
	return &checkinLinks{checkinLinksRepository}
}
