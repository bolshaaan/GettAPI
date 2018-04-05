package b2bclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	BaseURL string = "https://api.gett.com/"
	//BaseURL        string = "https://rides.gett.com/api/"
	//BaseURL string = "http://localhost:8087/"

	//CreateRidePath string = "v1/rides"
	CreateRidePath string = "v1/business/rides"

	//CreateRidePath string = "v1/business/rides"

	GetProductsPath string = "v1/business/products"
	AuthURL         string = "v1/oauth/token"

	TokenDataPath string = "tokenFile"
)

const testHeaders = false

type B2bClient struct {
	ClientID     string
	ClientSecret string
	GrantType    string
	Scope        string
	AuthData     *AuthResp

	aToken string
}

func NewB2bClient(clientID, cSecret, gType, scope string) (*B2bClient, error) {
	cl := &B2bClient{ClientID: clientID, ClientSecret: cSecret, GrantType: gType, Scope: scope}

	//if osk.IsExist

	// DEBUG
	//cl.AuthData = &AuthResp{AccessToken: "AAA"}
	if err := cl.Auth(); err != nil {
		return nil, err
	}
	return cl, nil
}

func (c *B2bClient) GetProducts() error {

	return nil
}

func (c *B2bClient) CreateRide(rr *RideRequest, bID string) error {
	b, err := json.Marshal(rr)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", BaseURL+CreateRidePath+"?business_id="+bID, bytes.NewReader(b))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer "+c.AuthData.AccessToken)

	// lets directly go to b2brides
	if testHeaders {
		req.Header.Add("Public-API-GW-Scopes", "business")
		//req.Header.Add("Public-API-GW-App",  "business")
		req.Header.Add("Public-API-GW-BusinessIds", bID)
	}

	cl := &http.Client{}
	resp, err := cl.Do(req)
	//resp, err := http.Post(BaseURL+CreateRidePath, "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	testPrintOut(bytes.NewReader(buf))

	return nil
}

func (c *B2bClient) LoadAuth() error {

	return nil
}

func (c *B2bClient) CheckAuthExpired() error {

	return nil
}

func (c *B2bClient) Auth() error {
	urlValues := make(url.Values)
	urlValues.Add("client_id", c.ClientID)
	urlValues.Add("client_secret", c.ClientSecret)
	urlValues.Add("grant_type", c.GrantType)
	urlValues.Add("scope", "business")

	resp, err := http.PostForm(BaseURL+AuthURL, urlValues)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	c.AuthData = &AuthResp{}
	if err := json.Unmarshal(buf, c.AuthData); err != nil {
		return err
	}

	if c.AuthData.AccessToken == "" {
		testPrintOut(bytes.NewReader(buf))
	}

	//writeGob(TokenDataPath)

	return nil
}

func testPrintOut(r io.Reader) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("(%s) -- Output: %s\n", MyCaller(), string(b))
}
