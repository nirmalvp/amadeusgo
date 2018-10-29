package types

//Action Has the set of actions possible
type Action uint8

const (
	//CREATE create resource request
	CREATE Action = 0X01
	//BULKREAD get request on a resource
	//READ is by default bulk
	BULKREAD Action = 0X02
	//READ flag on action
	READ Action = 0x04
	//UPDATE update request on a resource
	UPDATE Action = 0X08
	//DELETE delete request on a resource
	DELETE Action = 0X16
)
