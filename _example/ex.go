package main

import (
	"flag"
	"fmt"

	"github.com/coral/trafikverket"
)

func main() {

	var apikey = flag.String("apikey", "", "apikey")
	var searchstation = flag.String("station", "Reftele", "What station you are querying")
	flag.Parse()

	if *apikey == "" {
		panic("Please provide your API key")
	}

	tf := trafikverket.New(*apikey)

	station, err := tf.LookupStation(*searchstation)
	if err != nil {
		panic(err)
	}

	trann, err := tf.QueryTrainAnnouncementsAtLocationSignature(station.LocationSignature)
	if err != nil {
		panic(err)
	}

	for _, k := range trann.RESPONSE.RESULT[0].TrainAnnouncement {

		from, err := tf.LookupLocationSignature(k.FromLocation[0].LocationName)
		if err != nil {
			panic(err)
		}

		to, err := tf.LookupLocationSignature(k.ToLocation[0].LocationName)
		if err != nil {
			panic(err)
		}

		formatted := fmt.Sprintf("%02d:%02d",
			k.AdvertisedTimeAtLocation.Hour(),
			k.AdvertisedTimeAtLocation.Minute())

		fmt.Println(formatted+
			"    Operator: "+k.InformationOwner,
			"    From: "+from.AdvertisedLocationName+
				"    To: "+to.AdvertisedLocationName+
				"    Track: "+k.TrackAtLocation)
	}

}
