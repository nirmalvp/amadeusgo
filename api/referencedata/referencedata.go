package referencedata

type ReferenceData struct {
	Urls      string
	Locations string
	Airlines  *airlines
}

func NewReferenceData(urls, locations string, airlines *airlines) *ReferenceData {
	return &ReferenceData{
		Urls:      urls,
		Locations: locations,
		Airlines:  airlines,
	}
}
