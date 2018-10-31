package referencedata

type urls struct {
	CheckinLinks *checkinLinks
}

func NewUrls(checkinLinks *checkinLinks) *urls {
	return &urls{CheckinLinks: checkinLinks}
}
