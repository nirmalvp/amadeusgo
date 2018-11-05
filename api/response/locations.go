package response

import "github.com/nirmalvp/amadeusgo/api/request"

// LocationsRest represents the response from a location BULK GET
type LocationsRest struct {
	Meta Meta
	Data []locationData
}

// LocationRest represents the response from a location by Id GET
type LocationRest struct {
	Meta Meta
	Data locationData
}

type locationData struct {
	Type         string
	SubType      string
	Name         string
	DetailedName string
	Id           string
	Self         struct {
		Href    string
		methods []string
	}
	TimeZoneOffset string
	IataCode       string
	GeoCode        struct {
		Latitude  float64
		Longitude float64
	}
	Address struct {
		CityName    string
		CityCode    string
		CountryName string
		CountryCode string
		RegionCode  string
	}
	Analytics struct {
		Travelers struct {
			Score int
		}
	}
}

type Locations struct {
	AmadeusResponse
	Result LocationsRest
	Data   []locationData
}

type Location struct {
	AmadeusResponse
	Result LocationRest
	Data   locationData
}

func NewLocationsResponse(statusCode int, locationsRestResp LocationsRest, request request.AmadeusRequestData, isParsed bool) Locations {
	return Locations{
		AmadeusResponse: AmadeusResponse{
			StatusCode: statusCode,
			Request:    request,
			Parsed:     isParsed,
		},
		Result: locationsRestResp,
		Data:   locationsRestResp.Data,
	}

}

func NewLocationResponse(statusCode int, locationRestResp LocationRest, request request.AmadeusRequestData, isParsed bool) Location {
	return Location{
		AmadeusResponse: AmadeusResponse{
			StatusCode: statusCode,
			Request:    request,
			Parsed:     isParsed,
		},
		Result: locationRestResp,
		Data:   locationRestResp.Data,
	}

}
