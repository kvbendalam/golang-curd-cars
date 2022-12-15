package models

type Car struct {
	ID    string `json:"id,omitempty"`
	Model string `json:"model,omitempty"`
	Price int    `json:"price,omitempty"`
}
