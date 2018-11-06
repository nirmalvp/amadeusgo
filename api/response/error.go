package response

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/nirmalvp/amadeusgo/api/request"
)

// ResponseErrorData represents the body of the error responses produced by
// Amadeus API
type responseErrorData struct {
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
type responseErrorDataFallback struct {
	Status string
	Code   string
	Title  string
	Detail string
	Source struct {
		Parameter string
	}
}

type restErrorResponse struct {
	Errors           []responseErrorData
	Error            string
	ErrorDescription string `json:"error_description"`
	Code             int
	Title            string
}

type restErrorResponseFallback struct {
	Errors           []responseErrorDataFallback
	Error            string
	ErrorDescription string `json:"error_description"`
	Code             int
	Title            string
}

type responseErrorResponse struct {
	AmadeusResponse
	Result restErrorResponse
}
type ResponseError struct {
	Response    responseErrorResponse
	Code        string
	Description string
}

func (re ResponseError) Error() string {
	return fmt.Sprintf("Response = %+v\n Code = %s\n Description = %s\n", re.Response, re.Code, re.Description)

}

func convertRestErrorResponseFallback(restErrFallBack restErrorResponseFallback) (restErrorResponse, bool) {
	isSuccess := true
	var restErrorResponse restErrorResponse
	restErrorResponse.Error = restErrFallBack.Error
	restErrorResponse.ErrorDescription = restErrFallBack.ErrorDescription
	restErrorResponse.Code = restErrFallBack.Code
	restErrorResponse.Title = restErrFallBack.Title
	if len(restErrFallBack.Errors) > 0 {
		restErrorResponse.Errors = make([]responseErrorData, 0)
		for _, errData := range restErrFallBack.Errors {
			red := responseErrorData{Title: errData.Title, Detail: errData.Detail, Source: errData.Source}
			status, err := strconv.Atoi(errData.Status)
			isSuccess = isSuccess && err != nil
			red.Status = status
			code, err := strconv.Atoi(errData.Code)
			red.Code = code
			isSuccess = isSuccess && err != nil
			restErrorResponse.Errors = append(restErrorResponse.Errors, red)
		}
	}
	return restErrorResponse, isSuccess
}

func NewResponseError(statusCode int, errorResponseBody []byte, request request.AmadeusRequestData) ResponseError {
	var formatedErrorRestResponse restErrorResponse
	parseError := json.Unmarshal(errorResponseBody, &formatedErrorRestResponse)
	// Try two different error responses, falling back if the first doesnt work
	// This is done due to a inconsistency in the error body returned
	// from amadeus API
	if parseError != nil {
		var formatedErrorRestResponseFallback restErrorResponseFallback
		parseFallbackError := json.Unmarshal(errorResponseBody, &formatedErrorRestResponseFallback)
		if parseFallbackError != nil {
			return ResponseError{
				Response: responseErrorResponse{
					AmadeusResponse: AmadeusResponse{
						StatusCode: statusCode,
						Request:    request,
						Body:       string(errorResponseBody),
						Parsed:     false,
					},
				},
				Code:        "blah", // replace me
				Description: "blah", // replace me
			}
		}
		var isSuccess bool
		formatedErrorRestResponse, isSuccess = convertRestErrorResponseFallback(formatedErrorRestResponseFallback)
		if !isSuccess {
			return ResponseError{
				Response: responseErrorResponse{
					AmadeusResponse: AmadeusResponse{
						StatusCode: statusCode,
						Request:    request,
						Body:       string(errorResponseBody),
						Parsed:     false,
					},
				},
				Code:        "blah", // replace me
				Description: "blah", // replace me
			}
		}
	}
	return ResponseError{
		Response: responseErrorResponse{
			AmadeusResponse: AmadeusResponse{
				StatusCode: statusCode,
				Request:    request,
				Body:       string(errorResponseBody),
				Parsed:     true,
			},
			Result: formatedErrorRestResponse,
		},
		Code:        "blah", // replace me
		Description: "blah", // replace me
	}

}
