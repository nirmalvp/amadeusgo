package response

import "github.com/nirmalvp/amadeusgo/api/request"

type LocationsRest struct {
	Meta Meta
	Data []LocationsData
}

type LocationsData struct {
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
	Data   []LocationsData
}

func NewLocationsResponse(statusCode int, locationRestResp LocationsRest, request request.AmadeusRequestData, isParsed bool) Locations {
	return Locations{
		AmadeusResponse: AmadeusResponse{
			StatusCode: statusCode,
			Request:    request,
			Parsed:     isParsed,
		},
		Result: locationRestResp,
		Data:   locationRestResp.Data,
	}

}
