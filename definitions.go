package trafikverket

import (
	"encoding/xml"
	"time"
)

//XML SHITSHOW

type REQUEST struct {
	XMLName xml.Name `xml:"REQUEST"`
	Text    string   `xml:",chardata"`
	Login   LOGIN    `xml:"LOGIN"`
	Query   QUERY    `xml:"QUERY"`
}

type LOGIN struct {
	Text              string `xml:",chardata"`
	Authenticationkey string `xml:"authenticationkey,attr"`
}

type QUERY struct {
	Text          string   `xml:",chardata"`
	Objecttype    string   `xml:"objecttype,attr"`
	Schemaversion string   `xml:"schemaversion,attr"`
	Orderby       string   `xml:"orderby,attr,omitempty"`
	Filter        FILTER   `xml:"FILTER"`
	Include       []string `xml:"INCLUDE,omitempty"`
}

type FILTER struct {
	Text string `xml:",chardata"`
	EQ   []EQ   `xml:"EQ"`
	And  *AND   `xml:"AND,omitempty"`
}

type AND struct {
	Text string `xml:",chardata"`
	EQ   []EQ   `xml:"EQ"`
	Or   *OR    `xml:"OR,omitempty"`
}

type OR struct {
	Text string `xml:",chardata"`
	And  []AND  `xml:"AND"`
}

type EQ struct {
	XMLName xml.Name
	Text    string `xml:",chardata"`
	Name    string `xml:"name,attr"`
	Value   string `xml:"value,attr"`
}

//JSON SHITSHOW

type TrainAnnouncementResponse struct {
	RESPONSE struct {
		RESULT []struct {
			TrainAnnouncement []struct {
				ActivityID                 string    `json:"ActivityId"`
				ActivityType               string    `json:"ActivityType"`
				Advertised                 bool      `json:"Advertised"`
				AdvertisedTimeAtLocation   time.Time `json:"AdvertisedTimeAtLocation"`
				AdvertisedTrainIdent       string    `json:"AdvertisedTrainIdent"`
				Canceled                   bool      `json:"Canceled"`
				EstimatedTimeIsPreliminary bool      `json:"EstimatedTimeIsPreliminary"`
				FromLocation               []struct {
					LocationName string `json:"LocationName"`
					Priority     int    `json:"Priority"`
					Order        int    `json:"Order"`
				} `json:"FromLocation"`
				InformationOwner                      string    `json:"InformationOwner"`
				LocationSignature                     string    `json:"LocationSignature"`
				ModifiedTime                          time.Time `json:"ModifiedTime"`
				NewEquipment                          int       `json:"NewEquipment"`
				PlannedEstimatedTimeAtLocationIsValid bool      `json:"PlannedEstimatedTimeAtLocationIsValid"`
				ProductInformation                    []string  `json:"ProductInformation"`
				ScheduledDepartureDateTime            time.Time `json:"ScheduledDepartureDateTime"`
				TechnicalTrainIdent                   string    `json:"TechnicalTrainIdent"`
				ToLocation                            []struct {
					LocationName string `json:"LocationName"`
					Priority     int    `json:"Priority"`
					Order        int    `json:"Order"`
				} `json:"ToLocation"`
				TrackAtLocation string `json:"TrackAtLocation"`
				TypeOfTraffic   string `json:"TypeOfTraffic"`
				ViaToLocation   []struct {
					LocationName string `json:"LocationName"`
					Priority     int    `json:"Priority"`
					Order        int    `json:"Order"`
				} `json:"ViaToLocation,omitempty"`
				WebLink     string `json:"WebLink"`
				WebLinkName string `json:"WebLinkName"`
			} `json:"TrainAnnouncement"`
		} `json:"RESULT"`
	} `json:"RESPONSE"`
}
