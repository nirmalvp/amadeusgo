package service

import "log"

type Configuration struct {
	clientId        string
	clientSecret    string
	logger          *log.Logger
	logLevel        string
	accept          string
	host            string
	ssl             bool
	port            int
	languageVersion string
	appId           *string
	appVersion      *string
	clientVersion   string
}

func NewConfiguration(
	clientId string,
	clientSecret string,
	logger *log.Logger,
	logLevel string,
	accept string,
	host string,
	ssl bool,
	port int,
	languageVersion string,
	appId *string,
	appVersion *string,
	clientVersion string) Configuration {
	return Configuration{
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

}
