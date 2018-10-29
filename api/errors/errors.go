package errors

import "fmt"

/*type AmadeusError struct {
	Errors []struct {
		Code   int
		Title  string
		Detail string
		Status int
	}
}

func (ae AmadeusError) Error() string {
	return fmt.Sprint(ae.Errors)

}*/

type AmadeusError map[string]interface{}

func (ae AmadeusError) Error() string {
	errorString := ""
	for key, value := range ae {
		errorString = errorString + fmt.Sprintf("%s : %s, ", key, value)
	}
	return errorString
}
