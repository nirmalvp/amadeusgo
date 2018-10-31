package api

import (
	"log"

	"github.com/nirmalvp/amadeusgo/api/referencedata"
	"github.com/nirmalvp/amadeusgo/api/types"
)

var clientVersion string = "1.1.1"

type client struct {
	ReferencedData *referencedata.ReferenceData
}

type Configuration struct {
	clientId         string
	clientSecret     string
	logger           *log.Logger
	logLevel         string
	accept           string
	host             string
	useSSL           bool
	port             int
	languageVersion  string
	customAppId      *string
	customAppVersion *string
	clientVersion    string
}

func newClient(referencedataObj *referencedata.ReferenceData) client {
	return client{referencedataObj}
}

func generateClient(cb clientBuilder) client {
	config := Configuration{
		clientId:         cb.clientId,
		clientSecret:     cb.clientSecret,
		logger:           cb.logger,
		logLevel:         cb.logLevel,
		host:             cb.host,
		useSSL:           cb.useSSL,
		port:             cb.port,
		languageVersion:  cb.languageVersion,
		customAppId:      cb.customAppId,
		customAppVersion: cb.customAppVersion,
		clientVersion:    clientVersion,
		accept:           "application/json, application/vnd.amadeus+json",
	}

	//Create AccessToken Service here
	accessTokenUrlPaths := map[types.Action]string{
		types.CREATE: "/v1/security/oauth2/token",
	}
	accessTokenRepository := AmadeusHTTPRepository{nil, accessTokenUrlPaths, config}
	accessToken := AccessToken{AccessTokenRepository: accessTokenRepository}

	// Initialize Airlines Service here
	airlinesUrlPaths := map[types.Action]string{
		types.READ: "/v1/reference-data/airlines",
	}

	//airlineRepository is not an open resource and hence need a access token service
	airlinesRepository := AmadeusHTTPRepository{&accessToken, airlinesUrlPaths, config}
	airlines := referencedata.NewAirlines(airlinesRepository)

	// Initialize locations Service here
	locationsUrlPaths := map[types.Action]string{
		types.READ: "/v1/reference-data/locations",
	}

	//airlineRepository is not an open resource and hence need a access token service
	locationsRepository := AmadeusHTTPRepository{&accessToken, locationsUrlPaths, config}
	locations := referencedata.NewLocations(locationsRepository)

	// Initialize locations Service here
	checkinLinksUrlPaths := map[types.Action]string{
		types.READ: "/v2/reference-data/urls/checkin-links",
	}

	//airlineRepository is not an open resource and hence need a access token service
	checkinlinksRepository := AmadeusHTTPRepository{&accessToken, checkinLinksUrlPaths, config}
	checkinlinks := referencedata.NewCheckinLinks(checkinlinksRepository)

	urls := referencedata.NewUrls(checkinlinks)

	// Create referencedData here. Only airlines implemented as of now
	referencedataObj := referencedata.NewReferenceData(urls, locations, airlines)
	return newClient(referencedataObj)
}
