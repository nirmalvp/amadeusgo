package interfaces

import (
	"github.com/nirmalvp/amadeusgo/api/request"
)

type AmadeusRest interface {
	Send(request.AmadeusRequestData) (int, []byte, error)
}
