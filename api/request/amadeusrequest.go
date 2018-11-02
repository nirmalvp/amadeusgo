package request

import (
	"github.com/nirmalvp/amadeusgo/api/params"
)

type Verb string
type Header map[string]string

const (
	GET  Verb = "GET"
	POST Verb = "POST"
)

type AmadeusRequestData struct {
	Verb            Verb
	Scheme          string
	Host            string
	Path            string
	Params          params.Params
	BearerToken     *string
	ClientVersion   string
	LanguageVersion string
	UserAgent       string
	AppId           *string
	AppVersion      *string
	SSL             bool
	Port            int
	Accept          string
	Headers         Header
	URI             string
}
