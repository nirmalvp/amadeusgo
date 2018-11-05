package main

import (
	"fmt"

	"github.com/nirmalvp/amadeusgo/api"
	"github.com/nirmalvp/amadeusgo/api/params"
)

func main() {
	client := api.NewClientBuilder().Build()

	//urlParams := params.With("IATACode", "AI,3H")
	//response, err := client.ReferencedData.Airlines.GetWithParams(urlParams)
	//urlParams = params.With("keyword", "lon").And("subType", "AIRPORT,CITY")
	//response, err := client.ReferencedData.Locations.GetWithParams(urlParams)
	//urlParams := params.With("airline", "AF")
	//response, err := client.ReferencedData.Urls.CheckinLinks.GetWithParams(urlParams)
	urlParams := params.With("latitude", 49.0000).And("longitude", 2.55)
	//response, err := client.ReferencedData.Location("ALHR").Get()
	response, err := client.ReferencedData.Locations.Airports.GetWithParams(urlParams)

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
	fmt.Printf("%d\n", response.StatusCode)
	fmt.Printf("%s\n", response.Body)
	fmt.Printf("%+s\n", response.Request)
	fmt.Printf("%s\n", response.Parsed)
	fmt.Printf("%+s\n", response.Result)
	fmt.Printf("%+s\n", response.Data)

}
