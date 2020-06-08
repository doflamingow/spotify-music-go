package model

type Artist struct {
	ID       int    `from:"ID" json:"id,omitempty"`
	Name     string `from:"Name" json:"name,omitempty"`
	Debut    string `from:"Debut" json:"debut,omitempty"`
	Category string `from:"Category" json:"category,omitempty"`
	Genre    string `from:"Genre" json:"genre,omitempty"`
	ImageUrl string `from:"ImageUrl" json:"image,omitempty"`
}
