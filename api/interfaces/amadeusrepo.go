package interfaces

import "github.com/nirmalvp/amadeusgo/api/params"

type AmadeusRepository interface {
	Get(params.Params) (int, []byte, error)
	Create(params.Params) (int, []byte, error)
}
