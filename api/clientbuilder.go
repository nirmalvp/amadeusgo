package api

import (
	"log"
	"os"
	"runtime"
)

// Hosts represent the test and prod api of Amadeus. The user can switch to
// the Prod Api by calling Production() on the client builder
var hosts map[string]string = map[string]string{"test": "test.api.amadeus.com",
	"production": "api.amadeus.com",
}

// Client Builder takes all the customizable elements that can be
// provided by the user and uses the information to build the client
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

//NewClientBuilder returnsn a new client builder element to the user. The user
//can set customizable elements calling functions on the client builder and
//finally create the client using the build() function
func NewClientBuilder() clientBuilder {
	clientId := os.Getenv("AMADEUS_CLIENT_ID")
	clientSecret := os.Getenv("AMADEUS_CLIENT_SECRET")
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

// Production() forces the client to switch to production mode.
func (cb clientBuilder) Production() clientBuilder {
	cb.host = hosts["production"]
	return cb
}

// SetAuth() forces the client to use the user provided credentials to authenticate
// api calls
func (cb clientBuilder) SetAuth(clientId, clientSecret string) clientBuilder {
	cb.clientId = clientId
	cb.clientSecret = clientSecret
	return cb
}

// DisableSSL() downgrades the connection from HTTPS to HTTP
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
