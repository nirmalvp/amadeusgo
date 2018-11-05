package main

import (
	"fmt"

	"github.com/nirmalvp/amadeusgo/api"
	"github.com/nirmalvp/amadeusgo/api/params"
)

func main() {
	client := api.NewClientBuilder().Build()

	//urlParams := params.With("IATACode", "AI,3H")
	//statusCode, response, err := client.ReferencedData.Airlines.GetWithParams(urlParams)
	//urlParams = params.With("keyword", "lon").And("subType", "AIRPORT,CITY")
	//statusCode, response, err := client.ReferencedData.Locations.GetWithParams(urlParams)
	//urlParams := params.With("airline", "AF")
	//statusCode, response, err := client.ReferencedData.Urls.CheckinLinks.GetWithParams(urlParams)
	_ = params.With("IATACode", "AI,3H")
	statusCode, response, err := client.ReferencedData.Location("ALHR").Get()

	if err != nil {
		/*if responseError, ok := err.(errors.ResponseError); ok {
			// handle API errors
			fmt.Println(amadeusError)
		} else {*/
		//handle internal errors
		fmt.Println("An errors occured")
		fmt.Println(err)
		//}
		return
	}
	fmt.Println(statusCode)
	fmt.Printf("%+v\n", response)
}
