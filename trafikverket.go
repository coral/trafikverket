package trafikverket

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/coral/trafikverket/responses/apierror"
	"github.com/coral/trafikverket/responses/trainannouncement"
	"github.com/coral/trafikverket/responses/trainstation"
)

type TrafikVerket struct {
	apikey    string
	Endpoint  string
	locsigndb map[string]trainstation.TrainStation
}

//New takes an api key and creates the Trafikverket API wrapper
func New(apikey string) *TrafikVerket {
	return &TrafikVerket{
		apikey:    apikey,
		locsigndb: make(map[string]trainstation.TrainStation),
		Endpoint:  "https://api.trafikinfo.trafikverket.se/v2/data.json",
	}
}

func (tf *TrafikVerket) createRequest(rq REQUEST) ([]byte, error) {
	res, err := xml.MarshalIndent(rq, "", "	")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (tf *TrafikVerket) ExecuteRequest(rq REQUEST) (*http.Response, error) {

	rq.Login.Authenticationkey = tf.apikey

	t, _ := tf.createRequest(rq)

	return http.Post(tf.Endpoint, "application/xml", bytes.NewBuffer(t))

}

func (tf *TrafikVerket) parseError(a apierror.ERROR) error {

	if a.MESSAGE != "" {
		return fmt.Errorf(a.SOURCE + " - " + a.MESSAGE)
	}
	return nil
}

func (tf *TrafikVerket) LookupLocationSignature(locsig string) (trainstation.TrainStation, error) {

	if len(tf.locsigndb) == 0 {
		err := tf.initLocSignDb()
		if err != nil {
			return trainstation.TrainStation{}, err
		}
	}

	if station, ok := tf.locsigndb[locsig]; ok {
		return station, nil
	}

	return trainstation.TrainStation{}, fmt.Errorf("Could not find station")
}

func (tf *TrafikVerket) LookupStation(stationName string) (trainstation.TrainStation, error) {
	if len(tf.locsigndb) == 0 {
		err := tf.initLocSignDb()
		if err != nil {
			return trainstation.TrainStation{}, err
		}
	}

	for _, station := range tf.locsigndb {
		if station.AdvertisedLocationName == stationName {
			return station, nil
		}
	}

	return trainstation.TrainStation{}, fmt.Errorf("Could not find station")
}

func (tf *TrafikVerket) initLocSignDb() error {
	trs, err := tf.QueryTrainStations()
	if err != nil {
		return err
	}
	for _, station := range trs {
		tf.locsigndb[station.LocationSignature] = station
	}

	return nil
}

func (tf *TrafikVerket) QueryTrainStations() ([]trainstation.TrainStation, error) {

	rq := REQUEST{
		Query: QUERY{
			Objecttype:    "TrainStation",
			Schemaversion: "1",
			Filter: FILTER{
				EQ: []EQ{
					EQ{
						Name:  "Advertised",
						Value: "true",
					},
				},
			},
			Include: []string{
				"AdvertisedLocationName",
				"LocationSignature",
				"LocationInformationText",
				"Geometry.WGS84",
				"PlatformLine",
			},
		},
	}

	t, _ := tf.ExecuteRequest(rq)

	defer t.Body.Close()

	b, err := ioutil.ReadAll(t.Body)
	if err != nil {
		return nil, err
	}

	var tfr trainstation.Root

	if err := json.NewDecoder(bytes.NewBuffer(b)).Decode(&tfr); err != nil {
		return nil, err
	}

	if len(tfr.RESPONSE.RESULT) > 0 {
		err = tf.parseError(tfr.RESPONSE.RESULT[0].ERROR)
		if err != nil {
			return nil, err
		}
	}

	return tfr.RESPONSE.RESULT[0].TrainStation, nil

}

func (tf *TrafikVerket) QueryTrainAnnouncementsAtLocationSignature(LocationSignature string) ([]trainannouncement.TrainAnnouncement, error) {

	rq := REQUEST{
		Query: QUERY{
			Objecttype:    "TrainAnnouncement",
			Schemaversion: "1.3",
			Orderby:       "AdvertisedTimeAtLocation",
			Filter: FILTER{
				And: &AND{
					EQ: []EQ{
						EQ{
							Name:  "ActivityType",
							Value: "Avgang",
						},
						EQ{
							Name:  "LocationSignature",
							Value: LocationSignature,
						},
					},
					Or: &OR{
						And: []AND{
							AND{
								EQ: []EQ{
									EQ{
										XMLName: xml.Name{Local: "GT"},
										Name:    "AdvertisedTimeAtLocation",
										Value:   "$dateadd(-00:15:00)",
									},
									EQ{
										XMLName: xml.Name{Local: "LT"},
										Name:    "AdvertisedTimeAtLocation",
										Value:   "$dateadd(23:30:00)",
									},
								},
							},
							AND{
								EQ: []EQ{
									EQ{
										XMLName: xml.Name{Local: "LT"},
										Name:    "AdvertisedTimeAtLocation",
										Value:   "$dateadd(00:30:00)",
									},
									EQ{
										XMLName: xml.Name{Local: "GT"},
										Name:    "EstimatedTimeAtLocation",
										Value:   "$dateadd(-00:15:00)",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	t, _ := tf.ExecuteRequest(rq)

	defer t.Body.Close()

	b, err := ioutil.ReadAll(t.Body)
	if err != nil {
		return nil, err
	}

	var tfr trainannouncement.Root

	if err := json.NewDecoder(bytes.NewBuffer(b)).Decode(&tfr); err != nil {
		return nil, err
	}

	if len(tfr.RESPONSE.RESULT) > 0 {
		err = tf.parseError(tfr.RESPONSE.RESULT[0].ERROR)
		if err != nil {
			return nil, err
		}
	}

	return tfr.RESPONSE.RESULT[0].TrainAnnouncement, nil
}
