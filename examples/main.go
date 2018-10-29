package main

import (
	"fmt"

	"github.com/nirmalvp/amadeusgo/api"
	"github.com/nirmalvp/amadeusgo/api/errors"
	"github.com/nirmalvp/amadeusgo/api/params"
)

func main() {
	client := api.NewClientBuilder("YOUR_ID_HERE", "YOUR_KEY_HERE").Build()

	params := params.With("IATACode", "AI,3H")
	statusCode, response, err := client.ReferencedData.Airlines.GetWithParams(params)
	//statusCode, response, err := client.ReferencedData.Airlines.Get()

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
