package model

type Genre struct {
	ID int `from:"ID" json:"id"`
	Name string `from:"Name" json:"name"`
}
