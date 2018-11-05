package response

import (
	"encoding/json"
	"github.com/nirmalvp/amadeusgo/api/request"
)

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

func NewLocationsResponse(statusCode int, responseBody []byte, request request.AmadeusRequestData) Locations {
	var formatedRestResponse LocationsRest
	parseError := json.Unmarshal(responseBody, &formatedRestResponse)
	return Locations{
		AmadeusResponse: AmadeusResponse{
			StatusCode: statusCode,
			Body:       string(responseBody),
			Request:    request,
			Parsed:     parseError == nil,
		},
		Result: formatedRestResponse,
		Data:   formatedRestResponse.Data,
	}

}

func NewLocationResponse(statusCode int, responseBody []byte, request request.AmadeusRequestData) Location {
	var formatedRestResponse LocationRest
	parseError := json.Unmarshal(responseBody, &formatedRestResponse)
	return Location{
		AmadeusResponse: AmadeusResponse{
			StatusCode: statusCode,
			Body:       string(responseBody),
			Request:    request,
			Parsed:     parseError == nil,
		},
		Result: formatedRestResponse,
		Data:   formatedRestResponse.Data,
	}

}
