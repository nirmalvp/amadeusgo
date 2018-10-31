package referencedata

type ReferenceData struct {
	Urls      *urls
	Locations *locations
	Airlines  *airlines
}

func NewReferenceData(urls *urls, locations *locations, airlines *airlines) *ReferenceData {
	return &ReferenceData{
		Urls:      urls,
		Locations: locations,
		Airlines:  airlines,
	}
}

func (referenceData *ReferenceData) Location(locationId string) *Location {
	return NewLocation(locationId)
}
