package domain

import ()

type Poi struct {
	IdPoi int    `json:"id_poi,omitempty"`
	Nome  string `json:"nome"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
}
