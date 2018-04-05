package main

import (
	"fmt"
	"time"

	"github.com/bolshaaan/testb2brides/b2bclient"
)

//curl -F 'client_id=003caed893af3b44ba8f5986f9ac7272930444636d11eb5c8eff9085871ede98' \
//-F 'client_secret=f9b944c49b1988c8c0f133799eefd442bca4b3e45e66b1c33f53d3bf12c95cfb' \
//-F 'grant_type=client_credentials' \
//-F 'scope=business' \
//'https://api.gett.com/v1/oauth/token'
func main() {
	// auth

	// scrum50
	//cl, err := b2bclient.NewB2bClient("529c8c3d3ad9e0c90e795446b2d7c0cf55b3b29af29720cc6001183117008162",
	//	"f64d376a0ff156efa4aef290e30c7e5db4af391e327565b314420f6c7ef93bf2", "client_credentials", "business")

	// prod
	cl, err := b2bclient.NewB2bClient("003caed893af3b44ba8f5986f9ac7272930444636d11eb5c8eff9085871ede98",
		"f9b944c49b1988c8c0f133799eefd442bca4b3e45e66b1c33f53d3bf12c95cfb", "client_credentials", "business")

	if err != nil {
		panic(err)
	}

	//fmt.Println("TOKEN: ", cl.AuthData.AccessToken)
	// create ride

	start := time.Now()

	rr := &b2bclient.RideRequest{
		//ProductID:    "2f810d02-ddfa-47aa-8fb2-423e4ac2bcba", // GT prod
		ProductID:    "2f810d02-ddfa-47aa-8fb2-423e4ac2bcba", // scrum50
		NoteToDriver: "This is a test ride - DO NOT ACCEPT !",
		Rider: b2bclient.Rider{
			Name:        "Alexander Voloshin",
			PhoneNumber: "+79197707092",
		},
		Pickup: b2bclient.Pickup{
			Latitude:  55.724086,
			Longitude: 37.653638,
		},
		//Destination: b2bclient.Destination{
		//	Latitude:  55.724086 + 1,
		//	Longitude: 37.653638 + 1,
		//},
	}

	// RU-1999 -- prod
	// RU-7631 -- scrum50
	if err := cl.CreateRide(rr, "RU-1999"); err != nil {
		panic(err)
	}

	fmt.Println("TIME: ", time.Since(start))

	return
}
