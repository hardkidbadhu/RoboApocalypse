package model

import "time"

type Robots struct {
	Model            string `json:"model"`
	SerialNumber     string `json:"serialNumber"`
	ManufacturedDate time.Time `json:"manufacturedDate"`
	Category         string `json:"category"`
}

type RoboList []Robots

func (e RoboList) Len() int {
	return len(e)
}

func (e RoboList) Less(i, j int) bool {
	return e[i].ManufacturedDate.Before(e[j].ManufacturedDate)
}

func (e RoboList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}