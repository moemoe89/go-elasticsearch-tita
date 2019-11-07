package models

type Destination struct {
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	Photo            string  `json:"photo"`
	Overview         string  `json:"overview"`
	TouristAgreement string  `json:"tourist_agreement"`
	Duration         int     `json:"duration"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	MarkerIcon       string  `json:"marker_icon"`
	Categories       string  `json:"categories"`
	Rate             float64 `json:"rate"`
	Active           int     `json:"active"`
	City             City    `json:"city"`
}

type City struct {
	ID            int    `json:"id"`
	City          string `json:"city"`
	Photo         string `json:"photo"`
	GooglePlaceID string `json:"google_place_id"`
}
