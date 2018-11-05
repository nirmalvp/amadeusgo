package referencedata

import (
	"github.com/nirmalvp/amadeusgo/api/interfaces"
	"github.com/nirmalvp/amadeusgo/api/service"
)

type ReferenceData struct {
	Urls                        *urls
	Locations                   *locations
	Airlines                    *airlines
	RestClient                  interfaces.AmadeusRest
	AuthenticatedRequestCreator *service.AuthenticatedRequestCreator
}

func NewReferenceData(urls *urls,
	locations *locations,
	airlines *airlines,
	restClient interfaces.AmadeusRest,
	authenticatedRequestCreator *service.AuthenticatedRequestCreator) *ReferenceData {
	return &ReferenceData{
		Urls:                        urls,
		Locations:                   locations,
		Airlines:                    airlines,
		RestClient:                  restClient,
		AuthenticatedRequestCreator: authenticatedRequestCreator,
	}
}

func (referenceData *ReferenceData) Location(locationId string) *location {
	return NewLocation(locationId, referenceData.RestClient, referenceData.AuthenticatedRequestCreator)
}
