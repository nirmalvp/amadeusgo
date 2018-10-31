package response

type Locations struct {
	Meta Meta
	Data []struct {
		Type         string
		SubType      string
		Name         string
		DetailedName string
		Id           string
		Self         struct {
			Href    string
			methods []string
		}
		TimeZoneOffset string
		IataCode       string
		GeoCode        struct {
			Latitude  float64
			Longitude float64
		}
		Address struct {
			CityName    string
			CityCode    string
			CountryName string
			CountryCode string
			RegionCode  string
		}
		Analytics struct {
			Travelers struct {
				Score int
			}
		}
	}
}
