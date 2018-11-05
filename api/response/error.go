package response

import (
	"fmt"
	"strconv"

	"github.com/nirmalvp/amadeusgo/api/request"
)

// ResponseErrorData represents the body of the error responses produced by
// Amadeus API
type ResponseErrorData struct {
	Status int
	Code   int
	Title  string
	Detail string
	Source struct {
		Parameter string
	}
}

// ResponseErrorDataFallback represents the body of the error responses produced by
// Amadeus API. This is used as a fallback in case the error body doesnt match with ResponseErrorData.
// THis is required to compensate for the inconsitency of error message body in Amadeus API
type ResponseErrorDataFallback struct {
	Status string
	Code   string
	Title  string
	Detail string
	Source struct {
		Parameter string
	}
}

type RestErrorResponse struct {
	Errors           []ResponseErrorData
	Error            string
	ErrorDescription string `json:"error_description"`
	Code             int
	Title            string
}

type RestErrorResponseFallback struct {
	Errors           []ResponseErrorDataFallback
	Error            string
	ErrorDescription string `json:"error_description"`
	Code             int
	Title            string
}

type ResponseError struct {
	AmadeusResponse
	Code        string
	Description string
	Result      []ResponseErrorData
}

func (re ResponseError) Error() string {
	return fmt.Sprintf("AmadeusResponse = %+v\n Code = %s\n Description = %s\n Result = %s", re.AmadeusResponse, re.Code, re.Description, re.Result)

}

func NewResponseError(statusCode int, restResp interface{}, request request.AmadeusRequestData, isParsed bool) ResponseError {
	var restErrorResp RestErrorResponse
	// If the response is of the fallback format, convert it in to the actual format
	if restErrFallBack, ok := restResp.(RestErrorResponseFallback); ok {
		restErrorResp.Error = restErrFallBack.Error
		restErrorResp.ErrorDescription = restErrFallBack.ErrorDescription
		restErrorResp.Code = restErrFallBack.Code
		restErrorResp.Title = restErrFallBack.Title
		if len(restErrFallBack.Errors) > 0 {
			restErrorResp.Errors = make([]ResponseErrorData, 0)
			for _, errData := range restErrFallBack.Errors {
				red := ResponseErrorData{Title: errData.Title, Detail: errData.Detail, Source: errData.Source}
				status, err := strconv.Atoi(errData.Status)
				isParsed = isParsed && err != nil
				red.Status = status
				code, err := strconv.Atoi(errData.Code)
				red.Code = code
				isParsed = isParsed && err != nil
				restErrorResp.Errors = append(restErrorResp.Errors, red)
			}
		}
	} else {
		restErrorResp = restResp.(RestErrorResponse)
	}

	var result []ResponseErrorData
	// Some error responses have an array of errors, where as some have a single error
	// In case of a single error, convert it into a single element array anyway to provide
	// consistency
	if restErrorResp.Errors != nil {
		result = restErrorResp.Errors
	} else {
		result = append(result,
			ResponseErrorData{
				Status: statusCode,
				Code:   restErrorResp.Code,
				Title:  restErrorResp.Title,
				Detail: restErrorResp.ErrorDescription,
			},
		)
	}
	return ResponseError{
		AmadeusResponse: AmadeusResponse{
			StatusCode: statusCode,
			Request:    request,
			Parsed:     isParsed,
		},
		Code:        "blah", // replace me
		Description: "blah", // replace me
		Result:      result,
	}

}
