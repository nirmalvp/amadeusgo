package response

import (
	"encoding/json"
	"github.com/nirmalvp/amadeusgo/api/request"
)

type checkinLinksData struct {
	Type    string
	Id      string
	Href    string
	Channel string
}

type CheckinLinksRest struct {
	Meta Meta
	Data []checkinLinksData
}

type CheckinLinks struct {
	AmadeusResponse
	Result CheckinLinksRest
	Data   []checkinLinksData
}

func NewCheckinLinksResponse(statusCode int, responseBody []byte, request request.AmadeusRequestData) CheckinLinks {
	var formatedRestResponse CheckinLinksRest
	parseError := json.Unmarshal(responseBody, &formatedRestResponse)
	return CheckinLinks{
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
