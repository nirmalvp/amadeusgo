package service

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/nirmalvp/amadeusgo/api/params"
	"github.com/nirmalvp/amadeusgo/api/request"
)

func TestNewUnAuthenticatedRequestCreator(t *testing.T) {
	testAppId := "testAppId"
	testAppVersion := "testAppVersion"

	testCases := []struct {
		configuration Configuration
		expected      *UnAuthenticatedRequestCreator
	}{
		{
			configuration: Configuration{
				clientId:        "clientId",
				clientSecret:    "clientSecret",
				logger:          log.New(os.Stdout, "Amadeus", 0),
				logLevel:        "logLevel",
				accept:          "accept",
				host:            "host",
				ssl:             true,
				port:            443,
				languageVersion: "languageVersion",
				appId:           nil,
				appVersion:      nil,
				clientVersion:   "clientVersion",
			},
			expected: &UnAuthenticatedRequestCreator{
				clientId:        "clientId",
				clientSecret:    "clientSecret",
				host:            "host",
				languageVersion: "languageVersion",
				clientVersion:   "clientVersion",
				port:            443,
				accept:          "accept",
				scheme:          "https",
				userAgent:       "amadeus-go/clientVersion go/languageVersion",
			},
		},
		{
			configuration: Configuration{
				clientId:        "clientId",
				clientSecret:    "clientSecret",
				logger:          log.New(os.Stdout, "Amadeus", 0),
				logLevel:        "logLevel",
				accept:          "accept",
				host:            "host",
				ssl:             false,
				port:            80,
				languageVersion: "languageVersion",
				appId:           nil,
				appVersion:      nil,
				clientVersion:   "clientVersion",
			},
			expected: &UnAuthenticatedRequestCreator{
				clientId:        "clientId",
				clientSecret:    "clientSecret",
				host:            "host",
				languageVersion: "languageVersion",
				clientVersion:   "clientVersion",
				port:            80,
				accept:          "accept",
				scheme:          "http",
				userAgent:       "amadeus-go/clientVersion go/languageVersion",
			},
		},
		{
			configuration: Configuration{
				clientId:        "clientId",
				clientSecret:    "clientSecret",
				logger:          log.New(os.Stdout, "Amadeus", 0),
				logLevel:        "logLevel",
				accept:          "accept",
				host:            "host",
				ssl:             false,
				port:            80,
				languageVersion: "languageVersion",
				appId:           &testAppId,
				appVersion:      &testAppVersion,
				clientVersion:   "clientVersion",
			},
			expected: &UnAuthenticatedRequestCreator{
				clientId:        "clientId",
				clientSecret:    "clientSecret",
				host:            "host",
				languageVersion: "languageVersion",
				clientVersion:   "clientVersion",
				port:            80,
				accept:          "accept",
				scheme:          "http",
				userAgent:       "amadeus-go/clientVersion go/languageVersion testAppId/testAppVersion",
			},
		},
	}
	for _, testCase := range testCases {
		gotRequestCreator := NewUnAuthenticatedRequestCreator(testCase.configuration)
		if !reflect.DeepEqual(gotRequestCreator, testCase.expected) {
			t.Errorf("TestNewUnAuthenticatedRequestCreator, got: %v, want: %v.", gotRequestCreator, testCase.expected)
		}
	}
}

func TestCreate(t *testing.T) {
	testCases := []struct {
		requestCreator *UnAuthenticatedRequestCreator
		verb           request.Verb
		pathUrl        string
		params         params.Params
		expected       request.AmadeusRequestData
	}{
		{
			requestCreator: &UnAuthenticatedRequestCreator{
				clientId:        "clientId",
				clientSecret:    "clientSecret",
				host:            "host",
				languageVersion: "languageVersion",
				clientVersion:   "clientVersion",
				port:            443,
				accept:          "accept",
				scheme:          "https",
				userAgent:       "userAgent",
			},
			verb:    request.GET,
			pathUrl: "/pathUrl",
			params:  params.With("key", "value"),
			expected: request.AmadeusRequestData{
				Verb:   request.GET,
				Params: params.With("key", "value"),
				URI:    "https://host:443/pathUrl",
				Headers: map[string]string{
					"User-Agent": "userAgent",
					"Accept":     "accept",
				},
			},
		},
		{
			requestCreator: &UnAuthenticatedRequestCreator{
				clientId:        "clientId",
				clientSecret:    "clientSecret",
				host:            "host",
				languageVersion: "languageVersion",
				clientVersion:   "clientVersion",
				port:            443,
				accept:          "accept",
				scheme:          "https",
				userAgent:       "userAgent",
			},
			verb:    request.POST,
			pathUrl: "/pathUrl",
			params:  params.With("key", "value"),
			expected: request.AmadeusRequestData{
				Verb:   request.POST,
				Params: params.With("key", "value"),
				URI:    "https://host:443/pathUrl",
				Headers: map[string]string{
					"User-Agent": "userAgent",
					"Accept":     "accept",
				},
			},
		},
	}
	for _, testCase := range testCases {
		gotRequest := testCase.requestCreator.Create(testCase.verb, testCase.pathUrl, testCase.params)
		if !reflect.DeepEqual(gotRequest, testCase.expected) {
			t.Errorf("TestCreate, got: %v, want: %v.", gotRequest, testCase.expected)
		}
	}
}
