package main

import (
	"fmt"

	"github.com/nirmalvp/amadeusgo/api"
	"github.com/nirmalvp/amadeusgo/api/errors"
	"github.com/nirmalvp/amadeusgo/api/params"
)

func main() {
	client := api.NewClientBuilder("clientId", "clientSecret").Build()

	//params := params.With("IATACode", "AI,3H")
	//statusCode, response, err := client.ReferencedData.Airlines.GetWithParams(params)
	//params := params.With("keyword", "lon").And("subType", "AIRPORT,CITY")
	//statusCode, response, err := client.ReferencedData.Locations.GetWithParams(params)
	params := params.With("airline", "AF")
	statusCode, response, err := client.ReferencedData.Urls.CheckinLinks.GetWithParams(params)

	if err != nil {
		if amadeusError, ok := err.(errors.AmadeusError); ok {
			// handle API errors
			fmt.Println(amadeusError)
		} else {
			//handle internal errors
			fmt.Println(err)
		}
		return
	}
	fmt.Println(statusCode)
	fmt.Printf("%+v\n", response)
}
