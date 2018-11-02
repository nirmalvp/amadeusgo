package referencedata

type ReferenceData struct {
	Urls      string
	Locations *locations
	Airlines  *airlines
}

func NewReferenceData(urls string, locations *locations, airlines *airlines) *ReferenceData {
	return &ReferenceData{
		Urls:      urls,
		Locations: locations,
		Airlines:  airlines,
	}
}

/*func (referenceData *ReferenceData) Location(locationId string) *Location {
	return NewLocation(locationId)
}*/
