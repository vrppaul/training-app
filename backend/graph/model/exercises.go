package model

type Exercise struct {
	ID          string  `json:"_id" bson:"_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
