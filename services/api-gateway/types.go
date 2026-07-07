package main

import "ride-sharing/shared/types"

//go usa la primera letra para indicar si es public o private, en muchos lenguajes se usa la palabra especfica, por eso jsonmarshall es lo que se va a enviar y aqui hacemos la traduccion
type previewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}
