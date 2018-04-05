package b2bclient

import "time"

type AuthResp struct {
	AccessToken string `json:"access_token"`
	CreatedAt   int64  `json:"created_at"`
	ExpiresIn   int64  `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type Rider struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type Pickup struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"address"`
}
type Destination struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"address"`
}

type RideRequest struct {
	ProductID    string      `json:"product_id"`
	Rider        Rider       `json:"rider"`
	Pickup       Pickup      `json:"pickup"`
	Destination  Destination `json:"destination"`
	NoteToDriver string      `json:"note_to_driver"`
}

type CreateRideResponse struct {
	RideID    string `json:"ride_id"`
	ProductID string `json:"product_id"`
	Status    string `json:"status"`
	Rider     struct {
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
		ImageURL    string `json:"image_url"`
	} `json:"rider"`
	Pickup struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Address   string  `json:"address"`
	} `json:"pickup"`
	Destination struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Address   string  `json:"address"`
	} `json:"destination"`
	NoteToDriver string `json:"note_to_driver"`
	Reference    string `json:"reference"`
	StopPoints   []struct {
		Name      string  `json:"name"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Address   string  `json:"address"`
	} `json:"stop_points"`
	RequestedAt time.Time `json:"requested_at"`
	ScheduledAt time.Time `json:"scheduled_at"`
}
