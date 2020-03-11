package trainannouncment

import (
	"time"

	"github.com/coral/trafikverket/responses/apierror"
)

type Root struct {
	RESPONSE RESPONSE `json:"RESPONSE"`
}
type FromLocation struct {
	LocationName string `json:"LocationName"`
	Priority     int    `json:"Priority"`
	Order        int    `json:"Order"`
}
type ToLocation struct {
	LocationName string `json:"LocationName"`
	Priority     int    `json:"Priority"`
	Order        int    `json:"Order"`
}
type ViaToLocation struct {
	LocationName string `json:"LocationName"`
	Priority     int    `json:"Priority"`
	Order        int    `json:"Order"`
}
type TrainAnnouncement struct {
	ActivityID                            string          `json:"ActivityId"`
	ActivityType                          string          `json:"ActivityType"`
	Advertised                            bool            `json:"Advertised"`
	AdvertisedTimeAtLocation              time.Time       `json:"AdvertisedTimeAtLocation"`
	AdvertisedTrainIdent                  string          `json:"AdvertisedTrainIdent"`
	Canceled                              bool            `json:"Canceled"`
	EstimatedTimeIsPreliminary            bool            `json:"EstimatedTimeIsPreliminary"`
	FromLocation                          []FromLocation  `json:"FromLocation"`
	InformationOwner                      string          `json:"InformationOwner"`
	LocationSignature                     string          `json:"LocationSignature"`
	ModifiedTime                          time.Time       `json:"ModifiedTime"`
	NewEquipment                          int             `json:"NewEquipment"`
	PlannedEstimatedTimeAtLocationIsValid bool            `json:"PlannedEstimatedTimeAtLocationIsValid"`
	ProductInformation                    []string        `json:"ProductInformation"`
	ScheduledDepartureDateTime            time.Time       `json:"ScheduledDepartureDateTime"`
	TechnicalTrainIdent                   string          `json:"TechnicalTrainIdent"`
	ToLocation                            []ToLocation    `json:"ToLocation"`
	TrackAtLocation                       string          `json:"TrackAtLocation"`
	TypeOfTraffic                         string          `json:"TypeOfTraffic"`
	ViaToLocation                         []ViaToLocation `json:"ViaToLocation,omitempty"`
	WebLink                               string          `json:"WebLink"`
	WebLinkName                           string          `json:"WebLinkName"`
}
type RESULT struct {
	TrainAnnouncement []TrainAnnouncement `json:"TrainAnnouncement"`
	ERROR             apierror.ERROR      `json:"ERROR,omitempty"`
}
type RESPONSE struct {
	RESULT []RESULT `json:"RESULT"`
}
