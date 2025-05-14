package siribus

type SiriBus struct {
	Siri struct {
		ServiceDelivery struct {
			ResponseTimestamp      string `json:"ResponseTimestamp"`
			StopMonitoringDelivery []struct {
				MonitoredStopVisit []struct {
					MonitoredVehicleJourney struct {
						LineRef                 string `json:"LineRef"`
						DirectionRef            string `json:"DirectionRef"`
						FramedVehicleJourneyRef struct {
							DataFrameRef           string `json:"DataFrameRef"`
							DatedVehicleJourneyRef string `json:"DatedVehicleJourneyRef"`
						} `json:"FramedVehicleJourneyRef"`
						JourneyPatternRef        string   `json:"JourneyPatternRef"`
						PublishedLineName        []string `json:"PublishedLineName"`
						OperatorRef              string   `json:"OperatorRef"`
						OriginRef                string   `json:"OriginRef"`
						DestinationRef           string   `json:"DestinationRef"`
						DestinationName          []string `json:"DestinationName"`
						OriginAimedDepartureTime string   `json:"OriginAimedDepartureTime"`
						Monitored                bool     `json:"Monitored"`
						VehicleLocation          struct {
							Longitude float64 `json:"Longitude"`
							Latitude  float64 `json:"Latitude"`
						} `json:"VehicleLocation"`
						Bearing        float64  `json:"Bearing"`
						ProgressRate   string   `json:"ProgressRate"`
						ProgressStatus []string `json:"ProgressStatus"`
						BlockRef       string   `json:"BlockRef"`
						VehicleRef     string   `json:"VehicleRef"`
						MonitoredCall  struct {
							AimedArrivalTime      string   `json:"AimedArrivalTime"`
							ExpectedArrivalTime   string   `json:"ExpectedArrivalTime"`
							ArrivalProximityText  string   `json:"ArrivalProximityText"`
							ExpectedDepartureTime string   `json:"ExpectedDepartureTime"`
							DistanceFromStop      int      `json:"DistanceFromStop"`
							NumberOfStopsAway     int      `json:"NumberOfStopsAway"`
							StopPointRef          string   `json:"StopPointRef"`
							VisitNumber           int      `json:"VisitNumber"`
							StopPointName         []string `json:"StopPointName"`
						} `json:"MonitoredCall"`
					} `json:"MonitoredVehicleJourney"`
					RecordedAtTime string `json:"RecordedAtTime"`
				} `json:"MonitoredStopVisit"`
				ResponseTimestamp string `json:"ResponseTimestamp"`
				ValidUntil        string `json:"ValidUntil"`
				ErrorCondition    struct {
					OtherError struct {
						ErrorText string `json:"ErrorText"`
					} `json:"OtherError"`
					Description string `json:"Description"`
				} `json:"ErrorCondition"`
			} `json:"StopMonitoringDelivery"`
		} `json:"ServiceDelivery"`
	} `json:"Siri"`
}
