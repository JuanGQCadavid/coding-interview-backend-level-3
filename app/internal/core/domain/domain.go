package domain

type Item struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Price int    `json:"price,omitempty"` // TODO Could not be float  64?
}
