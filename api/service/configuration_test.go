package service

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func TestNewConfiguration(t *testing.T) {
	var (
		clientId                = "clientId"
		clientSecret            = "clientSecret"
		logger                  = log.New(os.Stdout, "Amadeus", 0)
		logLevel                = "debug"
		accept                  = "accept"
		host                    = "host"
		ssl                     = true
		port                    = 443
		languageVersion         = "languageVersion"
		appId           *string = nil
		appVersion      *string = nil
		clientVersion           = "clientVersion"
	)

	expectedConfiguration := Configuration{
		clientId:        clientId,
		clientSecret:    clientSecret,
		logger:          logger,
		logLevel:        logLevel,
		accept:          accept,
		host:            host,
		ssl:             ssl,
		port:            port,
		languageVersion: languageVersion,
		appId:           appId,
		appVersion:      appVersion,
		clientVersion:   clientVersion,
	}

	gotConfiguration := NewConfiguration(clientId,
		clientSecret,
		logger,
		logLevel,
		accept,
		host,
		ssl,
		port,
		languageVersion,
		appId,
		appVersion,
		clientVersion)

	if !reflect.DeepEqual(expectedConfiguration, gotConfiguration) {
		t.Errorf("NewClientBuilder, got: %+v, want: %+v", gotConfiguration, expectedConfiguration)
	}
}
