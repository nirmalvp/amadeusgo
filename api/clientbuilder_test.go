package api

import (
	"log"
	"os"
	"reflect"
	"runtime"
	"testing"
)

func init() {
	os.Setenv("AMADEUS_CLIENT_ID", "AMADEUS_CLIENT_ID")
	os.Setenv("AMADEUS_CLIENT_SECRET", "AMADEUS_CLIENT_SECRET")
}

func TestNewClientBuilder(t *testing.T) {
	languageVersion := runtime.Version()
	gotClientBuilder := NewClientBuilder()
	expectedClientBuilder := clientBuilder{
		clientId:        "AMADEUS_CLIENT_ID",
		clientSecret:    "AMADEUS_CLIENT_SECRET",
		logger:          log.New(os.Stdout, "Amadeus", 0),
		logLevel:        "debug",
		host:            "test.api.amadeus.com",
		useSSL:          true,
		port:            443,
		languageVersion: languageVersion,
	}
	if !reflect.DeepEqual(expectedClientBuilder, gotClientBuilder) {
		t.Errorf("NewClientBuilder, got: %+v, want: %+v", gotClientBuilder, expectedClientBuilder)
	}

}

func TestProduction(t *testing.T) {
	languageVersion := runtime.Version()
	gotClientBuilder := NewClientBuilder().Production()
	expectedClientBuilder := clientBuilder{
		clientId:        "AMADEUS_CLIENT_ID",
		clientSecret:    "AMADEUS_CLIENT_SECRET",
		logger:          log.New(os.Stdout, "Amadeus", 0),
		logLevel:        "debug",
		host:            "api.amadeus.com",
		useSSL:          true,
		port:            443,
		languageVersion: languageVersion,
	}
	if !reflect.DeepEqual(expectedClientBuilder, gotClientBuilder) {
		t.Errorf("NewClientBuilder, got: %+v, want: %+v", gotClientBuilder, expectedClientBuilder)
	}
}

func TestSetAuth(t *testing.T) {
	languageVersion := runtime.Version()
	gotClientBuilder := NewClientBuilder().SetAuth("customId", "customSecret")
	expectedClientBuilder := clientBuilder{
		clientId:        "customId",
		clientSecret:    "customSecret",
		logger:          log.New(os.Stdout, "Amadeus", 0),
		logLevel:        "debug",
		host:            "test.api.amadeus.com",
		useSSL:          true,
		port:            443,
		languageVersion: languageVersion,
	}
	if !reflect.DeepEqual(expectedClientBuilder, gotClientBuilder) {
		t.Errorf("NewClientBuilder, got: %+v, want: %+v", gotClientBuilder, expectedClientBuilder)
	}
}

func TestDisableSSL(t *testing.T) {
	languageVersion := runtime.Version()
	gotClientBuilder := NewClientBuilder().DisableSSL()
	expectedClientBuilder := clientBuilder{
		clientId:        "AMADEUS_CLIENT_ID",
		clientSecret:    "AMADEUS_CLIENT_SECRET",
		logger:          log.New(os.Stdout, "Amadeus", 0),
		logLevel:        "debug",
		host:            "test.api.amadeus.com",
		useSSL:          false,
		port:            80,
		languageVersion: languageVersion,
	}
	if !reflect.DeepEqual(expectedClientBuilder, gotClientBuilder) {
		t.Errorf("NewClientBuilder, got: %+v, want: %+v", gotClientBuilder, expectedClientBuilder)
	}
}

func TestSetCustomApp(t *testing.T) {
	gotClientBuilder := NewClientBuilder().SetCustomApp("CUSTOM_APP_ID", "CUSTOM_APP_VERSION")
	expectedCustomAppId := "CUSTOM_APP_ID"
	expectedCustomAppVersion := "CUSTOM_APP_VERSION"

	if *gotClientBuilder.customAppId != expectedCustomAppId {
		t.Errorf("NewClientBuilder, got customAppId: %s, want: %s", *gotClientBuilder.customAppId, expectedCustomAppId)
	}

	if *gotClientBuilder.customAppVersion != expectedCustomAppVersion {
		t.Errorf("NewClientBuilder, got customAppId: %s, want: %s", *gotClientBuilder.customAppVersion, expectedCustomAppVersion)
	}
}

func TestBuild(t *testing.T) {
	gotType := reflect.TypeOf(NewClientBuilder().Build()).String()
	expectedType := "api.client"

	if gotType != expectedType {
		t.Errorf("Build, got type: %s, want: %s", gotType, expectedType)
	}
}
