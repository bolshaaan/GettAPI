package main

import (
	"fmt"
	"time"

	"os"

	"regexp"

	"github.com/bolshaaan/testb2brides/b2bclient"
)

import _ "github.com/motemen/go-loghttp/global"

func testCall() {
	// prod
	//cl, err := b2bclient.NewB2bClient("003caed893af3b44ba8f5986f9ac7272930444636d11eb5c8eff9085871ede98",
	//	"f9b944c49b1988c8c0f133799eefd442bca4b3e45e66b1c33f53d3bf12c95cfb", "client_credentials", "business")

	//cl, err := b2bclient.LoadB2bClientFromFile("config/scrum51_ru12319.yml")
	cl, err := b2bclient.LoadB2bClientFromFile("config/prod.yml")
	if err != nil {
		panic(err)
	}
	//
	businessID := cl.BusinessIDs[0]
	fmt.Println("BusinessID: ", businessID)
	//
	//cl.GetRideDetails("183202874", businessID)
	//os.Exit(0)

	cl.BaseURL += "sandbox/" // testing sendbox

	fmt.Println(cl.GetRideDetails("1753873348", businessID))
	os.Exit(0)

	prodRes, err := cl.GetProducts(businessID, 55.724086, 37.653638)
	if err != nil {
		panic(err)
	}

	productID := prodRes.Products[0].ID
	for _, v := range prodRes.Products {
		if matched, _ := regexp.MatchString("Эконом", v.DisplayName); matched {
			fmt.Println("Found: ", v.DisplayName)
			fmt.Println(v)
			productID = v.ID
			break
		}
	}
	//
	//productID = "163eb7f7-8c1f-4509-8de6-5a04f657331c"

	productID = "6d89725f-f4a1-4fba-9952-a33e836fb78f"
	fmt.Println("productID: ", productID)
	fmt.Println("TOKEN: ", cl.AuthData.AccessToken)
	//create ride

	start := time.Now()

	rr := &b2bclient.RideRequest{
		//ProductID:    "2f810d02-ddfa-47aa-8fb2-423e4ac2bcba", // GT prod
		//ProductID:    "2f810d02-ddfa-47aa-8fb2-423e4ac2bcba", // scrum50
		ProductID:    productID,
		NoteToDriver: "This is a test ride - DO NOT ACCEPT !",
		Rider: b2bclient.Rider{
			Name:        "Alexander Voloshin",
			PhoneNumber: "+79197707092",
		},
		Pickup: b2bclient.Pickup{
			Latitude:  55.724086,
			Longitude: 37.653638,
		},
		//Pickup: b2bclient.Pickup{
		//	Latitude:  57.724086,
		//	Longitude: 37.653638,
		//},
		PaymentType: "voucher",
		//ExtraFields: map[string]interface{}{
		//	"Smoking Malboro": "Yes, please, more Malboro",
		//	//"company_field_#363": "Yes, please",
		//	//"company_field_#364": "Only Nikes :)",
		//},
		//Destination: b2bclient.Destination{
		//	Latitude:  55.724086 + 1,
		//	Longitude: 37.653638 + 1,
		//},
	}

	// RU-1999 -- prod
	// RU-7631 -- scrum50
	resp, err := cl.CreateRide(rr, businessID)
	if err != nil {
		panic(err)
	}

	os.Exit(0)

	cl.GetRideDetails(resp.RideID, businessID)

	fmt.Println("TIME: ", time.Since(start))
}

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

	//for i := 0; i < 10; i++ {
	testCall()
	//}

	return
}
