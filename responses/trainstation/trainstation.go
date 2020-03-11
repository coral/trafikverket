package trainstation

import "github.com/coral/trafikverket/responses/apierror"

type Root struct {
	RESPONSE RESPONSE `json:"RESPONSE"`
}
type Geometry struct {
	WGS84 string `json:"WGS84"`
}
type TrainStation struct {
	AdvertisedLocationName  string   `json:"AdvertisedLocationName"`
	Geometry                Geometry `json:"Geometry"`
	LocationSignature       string   `json:"LocationSignature"`
	PlatformLine            []string `json:"PlatformLine,omitempty"`
	LocationInformationText string   `json:"LocationInformationText,omitempty"`
}
type RESULT struct {
	TrainStation []TrainStation `json:"TrainStation"`
	ERROR        apierror.ERROR `json:"ERROR,omitempty"`
}
type RESPONSE struct {
	RESULT []RESULT `json:"RESULT"`
}
