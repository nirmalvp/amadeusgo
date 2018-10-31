package response

type CheckinLinks struct {
	Meta Meta
	Data []struct {
		Type    string
		Id      string
		Href    string
		Channel string
	}
}
