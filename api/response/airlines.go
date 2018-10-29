package response

type Airlines struct {
	Meta Meta
	Data []struct {
		Type         string
		IataCode     string
		BusinessName string
		CommonName   string
	}
}
