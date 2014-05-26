package drive

// Traffic represents the data structure returned from a 511.org pathlist (traffic conditions) API call
type Traffic struct {
	Routes []struct {
		TravelTime        int      `xml:"currentTravelTime"`
		TypicalTravelTime int      `xml:"typicalTravelTime"`
		Miles             float64  `xml:"miles"`
		Segments          []string `xml:"segments>segment>road"`
		Incidents         []string `xml:"incidents>incident"`
	} `xml:"path"`
}
