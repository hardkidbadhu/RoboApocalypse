package model

import (
	"fmt"
)

type Response struct {
	Message string `json:"message"`
}

type SurvivorPayload struct {
	Name            string    `json:"name"`
	Gender          string    `json:"gender"`
	Age             int       `json:"age"`
	InfectionStatus string    `json:"infectionStatus,omitempty"`
	Location        Location  `json:"location,omitempty"`
	Inventory       Inventory `json:"inventory,omitempty"`
}
type Survivor struct {
	Id               int    `json:"id"`
	CoOrdinates      string `json:"coOrdinates"`
	*SurvivorPayload `json:",inline,omitempty"`
}

type Inventory struct {
	Water      int `json:"water,omitempty"`
	Food       int `json:"food,omitempty"`
	Medication int `json:"medication,omitempty"`
	Ammunition int `json:"ammunition,omitempty"`
}

type PercentageResp struct {
	Percentage string `json:"percentage"`
	InfectionStatus string `json:"infectionStatus"`
}

type Location struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

type FlagPayload struct {
	InfectedUser    int `json:"infectedUser" binding:"required"`
	User            int `json:"user" binding:"required"`
	InfectionStatus int `json:"infectionStatus" binding:"required"`
}

//ValidateInfectionStatus validates whether the infection status is valid
func (fPayload *FlagPayload) ValidateInfectionStatus() bool {
	if fPayload.InfectionStatus == 0 || fPayload.InfectionStatus == 1 {
		return true
	}
	return false
}

func (l *Location) LatLongToPointStr() string {
	return fmt.Sprintf("Point(%f %f)", l.Latitude, l.Longitude)
}
