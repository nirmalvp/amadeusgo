package service

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/nirmalvp/amadeusgo/api/interfaces"
	"github.com/nirmalvp/amadeusgo/api/params"
	"github.com/nirmalvp/amadeusgo/api/request"
)

type testAccessTokenRestClient struct{}

func (rc *testAccessTokenRestClient) Send(httpRequest request.AmadeusRequestData) (int, []byte, error) {
	if httpRequest.Verb == request.POST &&
		(httpRequest.URI == "https://host:443/v1/security/oauth2/token" ||
			httpRequest.URI == "https://host:80/v1/security/oauth2/token") &&
		reflect.DeepEqual(httpRequest.Params,
			params.With("grant_type", "client_credentials").
				And("client_id", "client_id").
				And("client_secret", "client_secret")) &&
		httpRequest.BearerToken == nil &&
		reflect.DeepEqual(httpRequest.Headers, request.Header{
			"User-Agent": "amadeus-go/clientVersion go/languageVersion",
			"Accept":     "accept",
		}) {
		return 200, []byte(`{"access_token": "access_token","expires_in": 10}`), nil
	}
	return 0, nil, errors.New("An error occured")
}

type testTimeGetter struct {
	currentTime time.Time
}

func (ttg testTimeGetter) GetCurrentTime() time.Time {
	return ttg.currentTime
}

func TestNewAccessTokenService(t *testing.T) {
	timeGetter := testTimeGetter{time.Now()}
	testCases := []struct {
		restClient                    interfaces.AmadeusRest
		unAuthenticatedRequestCreator *UnAuthenticatedRequestCreator
		tokenBufferTime               time.Duration
		expected                      *accessTokenService
	}{
		{
			restClient: &testAccessTokenRestClient{},
			unAuthenticatedRequestCreator: &UnAuthenticatedRequestCreator{
				host:            "host",
				languageVersion: "languageVersion",
				clientVersion:   "clientVersion",
				port:            443,
				accept:          "accept",
				scheme:          "https",
				userAgent:       "amadeus-go/clientVersion go/languageVersion",
			},
			tokenBufferTime: time.Duration(10),
			expected: &accessTokenService{
				pathUrl:    "/v1/security/oauth2/token",
				restClient: &testAccessTokenRestClient{},
				unAuthenticatedRequestCreator: &UnAuthenticatedRequestCreator{
					host:            "host",
					languageVersion: "languageVersion",
					clientVersion:   "clientVersion",
					port:            443,
					accept:          "accept",
					scheme:          "https",
					userAgent:       "amadeus-go/clientVersion go/languageVersion",
				},
				tokenBufferTime: time.Duration(10),
				timeGetter:      timeGetter,
			},
		},
	}
	for _, testCase := range testCases {
		gotAccessTokenService := NewAccessTokenService(testCase.restClient,
			testCase.unAuthenticatedRequestCreator,
			testCase.tokenBufferTime,
			timeGetter,
		)
		if !reflect.DeepEqual(gotAccessTokenService, testCase.expected) {
			t.Errorf("TestNewAccessTokenService, got: %v, want: %v.", gotAccessTokenService, testCase.expected)
		}
	}
}

func TestNeedsRefresh(t *testing.T) {
	currentTime := time.Now()
	expiresAt := currentTime.Add(time.Duration(100))
	testCases := []struct {
		accessTokenService *accessTokenService
		expected           bool
	}{
		{
			accessTokenService: &accessTokenService{},
			expected:           true,
		},
		{
			accessTokenService: &accessTokenService{
				accessToken:     new(string),
				tokenBufferTime: time.Duration(9),
				expiresAt:       &expiresAt,
				timeGetter:      testTimeGetter{expiresAt.Add(-10)},
			},
			expected: false,
		},
		{
			accessTokenService: &accessTokenService{
				accessToken:     new(string),
				tokenBufferTime: time.Duration(11),
				expiresAt:       &expiresAt,
				timeGetter:      testTimeGetter{expiresAt.Add(-10)},
			},
			expected: true,
		},
	}
	for _, testCase := range testCases {
		gotNeedsRefresh := testCase.accessTokenService.needsRefresh()
		if gotNeedsRefresh != testCase.expected {
			t.Errorf("TestNeedsRefresh, got: %v \n want: %v \n TestCase : %v", gotNeedsRefresh, testCase.expected, testCase)
		}
	}
}

func TestFetchAccessToken(t *testing.T) {
	unAuthenticatedRequestCreator := &UnAuthenticatedRequestCreator{
		host:            "host",
		languageVersion: "languageVersion",
		clientVersion:   "clientVersion",
		port:            443,
		accept:          "accept",
		scheme:          "https",
		userAgent:       "amadeus-go/clientVersion go/languageVersion",
	}
	restClient := new(testAccessTokenRestClient)
	accessTokenService := NewAccessTokenService(restClient,
		unAuthenticatedRequestCreator,
		time.Duration(1),
		testTimeGetter{},
	)
	testCases := []struct {
		clientId           string
		clientSecret       string
		expectedStatusCode int
		expectedBody       []byte
		expectedError      error
	}{
		{
			clientId:           "client_id",
			clientSecret:       "client_secret",
			expectedStatusCode: 200,
			expectedBody:       []byte(`{"access_token": "access_token","expires_in": 10}`),
			expectedError:      nil,
		},
		{
			clientId:           "wrong_client_id",
			clientSecret:       "wrong_client_secret",
			expectedStatusCode: 0,
			expectedBody:       nil,
			expectedError:      errors.New("An error occured"),
		},
	}
	for _, testCase := range testCases {
		gotStatusCode, gotBody, gotError := accessTokenService.fetchAccessToken(testCase.clientId, testCase.clientSecret)
		if gotStatusCode != testCase.expectedStatusCode ||
			!reflect.DeepEqual(gotBody, testCase.expectedBody) ||
			!reflect.DeepEqual(gotError, testCase.expectedError) {
			t.Errorf("TestFetchAccessToken, got: (%d, %v, %v) \n want: (%d, %v, %v) \n TestCase : %v",
				gotStatusCode, string(gotBody), gotError,
				testCase.expectedStatusCode, string(testCase.expectedBody), testCase.expectedError,
				testCase,
			)
		}
	}
}

func TestGetUpdatedAccessToken(t *testing.T) {
	currentTime := time.Now()
	unAuthenticatedRequestCreator := &UnAuthenticatedRequestCreator{
		host:            "host",
		languageVersion: "languageVersion",
		clientVersion:   "clientVersion",
		port:            443,
		accept:          "accept",
		scheme:          "https",
		userAgent:       "amadeus-go/clientVersion go/languageVersion",
	}
	restClient := new(testAccessTokenRestClient)
	accessTokenService := NewAccessTokenService(restClient,
		unAuthenticatedRequestCreator,
		time.Duration(1),
		testTimeGetter{currentTime},
	)
	testCases := []struct {
		clientId            string
		clientSecret        string
		expectedAccessToken string
		expectedExpiresAt   time.Time
		expectedIsError     bool
	}{
		{
			clientId:            "client_id",
			clientSecret:        "client_secret",
			expectedAccessToken: "access_token",
			expectedExpiresAt:   currentTime.Add(10 * time.Second),
			expectedIsError:     false,
		},
		{
			clientId:        "wrong_client_id",
			clientSecret:    "wrong_client_secret",
			expectedIsError: true,
		},
	}
	for _, testCase := range testCases {
		gotAccessToken, gotExpiresAt, gotError := accessTokenService.getUpdatedAccessToken(testCase.clientId, testCase.clientSecret)
		gotIsError := (gotError != nil)
		if testCase.expectedIsError == gotIsError && gotIsError == true {
			return
		}
		if testCase.expectedIsError != gotIsError {
			t.Errorf("TestGetUpdatedAccessToken, got unexpected error : %v \n TestCase : %v", gotError, testCase)
			return
		}
		if *gotAccessToken != testCase.expectedAccessToken ||
			*gotExpiresAt != testCase.expectedExpiresAt {
			t.Errorf("TestGetUpdatedAccessToken, got: (%v, %v, %v) \n want: (%v, %v, %v) \n TestCase : %+v",
				*gotAccessToken, *gotExpiresAt, gotIsError,
				testCase.expectedAccessToken, testCase.expectedExpiresAt, testCase.expectedIsError,
				testCase,
			)
		}
	}
}
