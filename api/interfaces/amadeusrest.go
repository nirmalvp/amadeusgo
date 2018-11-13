package interfaces

import (
	"time"

	"github.com/nirmalvp/amadeusgo/api/request"
)

type AmadeusRest interface {
	Send(request.AmadeusRequestData) (int, []byte, error)
}

type TimeGetter interface {
	GetCurrentTime() time.Time
}
