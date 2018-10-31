package response

type Meta struct {
	Count int
	Links struct {
		Self string
		Next string
		Last string
	}
}
