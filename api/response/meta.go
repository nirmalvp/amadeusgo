package response

// Meta represents the Meta part included in every amadeus API response
type Meta struct {
	Count int
	Links struct {
		Self string
		Next string
		Last string
	}
}
