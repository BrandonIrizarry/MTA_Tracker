package onebusaway

type OneBusAway struct {
	Code        int   `json:"code"`
	CurrentTime int64 `json:"currentTime"`
	Data        struct {
		Entry struct {
			Polylines     []any  `json:"polylines"`
			RouteID       string `json:"routeId"`
			StopGroupings []struct {
				Ordered    bool `json:"ordered"`
				StopGroups []struct {
					ID   string `json:"id"`
					Name struct {
						Name  string   `json:"name"`
						Names []string `json:"names"`
						Type  string   `json:"type"`
					} `json:"name"`
					Polylines []any    `json:"polylines"`
					StopIds   []string `json:"stopIds"`
					SubGroups []any    `json:"subGroups"`
				} `json:"stopGroups"`
				Type string `json:"type"`
			} `json:"stopGroupings"`
			StopIds []string `json:"stopIds"`
		} `json:"entry"`
		References struct {
			Agencies []struct {
				Disclaimer     string `json:"disclaimer"`
				ID             string `json:"id"`
				Lang           string `json:"lang"`
				Name           string `json:"name"`
				Phone          string `json:"phone"`
				PrivateService bool   `json:"privateService"`
				Timezone       string `json:"timezone"`
				URL            string `json:"url"`
			} `json:"agencies"`
			Routes []struct {
				AgencyID    string `json:"agencyId"`
				Color       string `json:"color"`
				Description string `json:"description"`
				ID          string `json:"id"`
				LongName    string `json:"longName"`
				ShortName   string `json:"shortName"`
				TextColor   string `json:"textColor"`
				Type        int    `json:"type"`
				URL         string `json:"url"`
			} `json:"routes"`
			Situations []any `json:"situations"`
			Stops      []struct {
				Code               string   `json:"code"`
				Direction          string   `json:"direction"`
				ID                 string   `json:"id"`
				Lat                float64  `json:"lat"`
				LocationType       int      `json:"locationType"`
				Lon                float64  `json:"lon"`
				Name               string   `json:"name"`
				RouteIds           []string `json:"routeIds"`
				WheelchairBoarding string   `json:"wheelchairBoarding"`
			} `json:"stops"`
			Trips []any `json:"trips"`
		} `json:"references"`
	} `json:"data"`
	Text    string `json:"text"`
	Version int    `json:"version"`
}
