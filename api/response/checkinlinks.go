package response

import "github.com/nirmalvp/amadeusgo/api/request"

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

func NewCheckinLinksResponse(statusCode int, checkinLinksRest CheckinLinksRest, request request.AmadeusRequestData, isParsed bool) CheckinLinks {
	return CheckinLinks{
		AmadeusResponse: AmadeusResponse{
			StatusCode: statusCode,
			Request:    request,
			Parsed:     isParsed,
		},
		Result: checkinLinksRest,
		Data:   checkinLinksRest.Data,
	}

}
