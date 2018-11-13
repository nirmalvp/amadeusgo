package api

import (
	"time"

	"github.com/nirmalvp/amadeusgo/api/referencedata"
	"github.com/nirmalvp/amadeusgo/api/service"
)

var clientVersion string = "1.1.1"

type client struct {
	ReferencedData *referencedata.ReferenceData
}

//generateClient takes in a clientBuilder obj and created a new client using
//the information set in the client builder
func generateClient(cb clientBuilder) client {
	config := service.NewConfiguration(
		cb.clientId,
		cb.clientSecret,
		cb.logger,
		cb.logLevel,
		"application/json, application/vnd.amadeus+json",
		cb.host,
		cb.useSSL,
		cb.port,
		cb.languageVersion,
		cb.customAppId,
		cb.customAppVersion,
		clientVersion,
	)

	restClient := &AmadeusRestClient{}
	unAuthenticatedRequestCreator := service.NewUnAuthenticatedRequestCreator(config)
	bufferTime := time.Duration(10) * time.Second
	timeGetter := service.SystemTimeGetter{}
	accessTokenService := service.NewAccessTokenService(restClient,
		unAuthenticatedRequestCreator,
		bufferTime,
		timeGetter,
	)
	authenticatedRequestCreator := service.NewAuthenticatedRequestCreator(config, accessTokenService)
	airlines := referencedata.NewAirlines(restClient, authenticatedRequestCreator)
	airports := referencedata.NewAirports(restClient, authenticatedRequestCreator)
	locations := referencedata.NewLocations(restClient, authenticatedRequestCreator, airports)
	checkinlinks := referencedata.NewCheckinLinks(restClient, authenticatedRequestCreator)
	urls := referencedata.NewUrls(checkinlinks)

	/*// Initialize locations Service here
	checkinLinksUrlPaths := map[types.Action]string{
		types.READ: "/v2/reference-data/urls/checkin-links",
	}

	//airlineRepository is not an open resource and hence need a access token service
	checkinlinksRepository := AmadeusHTTPRepository{&accessToken, checkinLinksUrlPaths, config}
	checkinlinks := referencedata.NewCheckinLinks(checkinlinksRepository)

	urls := referencedata.NewUrls(checkinlinks)*/

	// Create referencedData here. Only airlines implemented as of now*/
	referencedataObj := referencedata.NewReferenceData(urls, locations, airlines, restClient, authenticatedRequestCreator)
	return client{referencedataObj}
}
