package b2bclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	//BaseURL string = "https://api.gett.com/"
	//BaseURL string = "https://publicapi-scrum12.gtforge.com/"
	//BaseURL        string = "https://rides.gett.com/api/"
	//BaseURL string = "http://localhost:8087/"

	//CreateRidePath string = "v1/rides"
	CreateRidePath string = "v1/business/rides"
	GetRidePath    string = "v1/business/rides"

	//CreateRidePath string = "v1/business/rides"

	GetProductsPath string = "v1/business/products"
	AuthURL         string = "v1/oauth/token"

	TokenDataPath string = "tokenFile"
)

const testHeaders = false

type B2bClient struct {
	ClientID     string   `yaml:"client_id"`
	ClientSecret string   `yaml:"client_secret"`
	GrantType    string   `yaml:"grant_type"`
	Scope        string   `yaml:"scope"`
	BusinessIDs  []string `yaml:"business_ids"`
	BaseURL      string   `yaml:"base_url"`

	AuthData *AuthResp
	aToken   string
}

func LoadB2bClientFromFile(confFileName string) (*B2bClient, error) {
	cl := &B2bClient{}
	b, err := ioutil.ReadFile(confFileName)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(b, cl); err != nil {
		return nil, err
	}

	isExpired, err := cl.LoadAuth()
	if !isExpired && err == nil {
		return cl, nil
	}

	if err := cl.Auth(); err != nil {
		return nil, err
	}

	return cl, nil
}

func NewB2bClient(clientID, cSecret, gType, scope string) (*B2bClient, error) {
	cl := &B2bClient{ClientID: clientID, ClientSecret: cSecret, GrantType: gType, Scope: scope}

	//if osk.IsExist

	// DEBUG
	//cl.AuthData = &AuthResp{AccessToken: "AAA"}

	isExpired, err := cl.LoadAuth()

	if !isExpired && err == nil {
		return cl, nil
	}

	if err := cl.Auth(); err != nil {
		return nil, err
	}
	return cl, nil
}

func (c *B2bClient) GetProducts(bID string, lat, lon float64) (*GetProductsResp, error) {
	vals := make(url.Values)
	vals.Add("business_id", bID)
	vals.Add("latitude", fmt.Sprintf("%f", lat))
	vals.Add("longitude", fmt.Sprintf("%f", lon))

	req, err := http.NewRequest("GET", c.BaseURL+GetProductsPath+"?"+vals.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{}
	req.Header.Add("Authorization", "Bearer "+c.AuthData.AccessToken)

	cl := &http.Client{}
	resp, err := cl.Do(req)
	if err != nil {
		return nil, nil
	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	prodResp := &GetProductsResp{}
	if err := json.Unmarshal(buf, prodResp); err != nil {
		testPrintOut(bytes.NewReader(buf))
		return nil, err
	}

	return prodResp, nil
}

func (c *B2bClient) GetRideDetails(rideID, bID string) (*GetRideDetailResponse, error) {
	params := make(url.Values)
	params.Add("business_id", bID)

	fmt.Println(rideID)
	req, err := http.NewRequest("GET", c.BaseURL+GetRidePath+"/"+rideID+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{}
	req.Header.Add("Authorization", "Bearer "+c.AuthData.AccessToken)

	cl := &http.Client{}
	resp, err := cl.Do(req)
	if err != nil {
		return nil, nil
	}

	testPrintOut(resp.Body)

	return nil, nil
}

func (c *B2bClient) CreateRide(rr *RideRequest, bID string) (*RideRequestResponse, error) {
	b, err := json.Marshal(rr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.BaseURL+CreateRidePath+"?business_id="+bID, bytes.NewReader(b))
	if err != nil {
		return nil, err
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
		return nil, err
	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	respRide := &RideRequestResponse{}
	if err := json.Unmarshal(buf, respRide); err != nil {
		return nil, err
	}

	//testPrintOut(bytes.NewReader(buf))

	return respRide, nil
}

func (c *B2bClient) LoadAuth() (bool, error) {
	c.AuthData = &AuthResp{}

	if err := readGob(TokenDataPath, c.AuthData); err != nil {
		return true, err
	}

	return c.IsAuthExpired(), nil
}

func (c *B2bClient) IsAuthExpired() bool {
	return c.AuthData.CreatedAt+c.AuthData.ExpiresIn < time.Now().Unix()
}

func (c *B2bClient) Auth() error {
	urlValues := make(url.Values)
	urlValues.Add("client_id", c.ClientID)
	urlValues.Add("client_secret", c.ClientSecret)
	urlValues.Add("grant_type", c.GrantType)
	urlValues.Add("scope", "business")

	resp, err := http.PostForm(c.BaseURL+AuthURL, urlValues)
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

	if err := writeGob(TokenDataPath, c.AuthData.AccessToken); err != nil {
		return err
	}

	return nil
}

func testPrintOut(r io.Reader) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("(%s) -- Output: %s\n", MyCaller(), string(b))
}
