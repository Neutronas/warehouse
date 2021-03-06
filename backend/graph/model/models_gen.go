// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Location struct {
	Name *string  `json:"name"`
	Lat  *float64 `json:"lat"`
	Lng  *float64 `json:"lng"`
}

type NewWarehouse struct {
	Name         string   `json:"name"`
	Description  *string  `json:"description"`
	LocationName *string  `json:"locationName"`
	Lat          *float64 `json:"lat"`
	Lng          *float64 `json:"lng"`
}

type Warehouse struct {
	ID          string    `json:"_id" bson:"_id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Location    *Location `json:"location"`
}
