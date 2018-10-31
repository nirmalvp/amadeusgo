package api

import (
	"log"
	"os"
	"runtime"
)

var hosts map[string]string = map[string]string{"test": "test.api.amadeus.com",
	"production": "test.api.amadeus.com",
}

type clientBuilder struct {
	clientId         string
	clientSecret     string
	logger           *log.Logger
	logLevel         string
	host             string
	useSSL           bool
	port             int
	languageVersion  string
	customAppId      *string
	customAppVersion *string
}

func NewClientBuilder(clientId, clientSecret string) clientBuilder {
	clientId = os.Getenv("AMADEUS_CLIENT_ID")
	clientSecret = os.Getenv("AMADEUS_CLIENT_SECRET")
	return clientBuilder{
		clientId:        clientId,
		clientSecret:    clientSecret,
		logger:          log.New(os.Stdout, "Amadeus", 0),
		logLevel:        "debug",
		host:            hosts["test"],
		useSSL:          true,
		port:            443,
		languageVersion: runtime.Version(),
	}
}

func (cb clientBuilder) Production() clientBuilder {
	cb.host = hosts["production"]
	return cb
}

func (cb clientBuilder) DisableSSL() clientBuilder {
	cb.port = 80
	cb.useSSL = false
	return cb
}

func (cb clientBuilder) SetCustomApp(appId, appVersion string) clientBuilder {
	cb.customAppId = &appId
	cb.customAppVersion = &appVersion
	return cb
}

func (cb clientBuilder) Build() client {
	return generateClient(cb)
}
