package model

import "time"

type Shipment struct {
	Id          string     `json:"id"`
	DateCreated time.Time  `json:"date_created"`
	LastUpdated *time.Time `json:"last_updated"`
	OrderId     string     `json:"order_id"`
	Mode        string     `json:"mode"`
	Distributor string     `json:"distributor"`
	Cost        float64    `json:"cost"`
	Payer       Payer      `json:"payer"`
	Seller      Seller     `json:"seller"`
}

type Address struct {
	StreetName   string `json:"street_name"`
	StreetNumber int    `json:"street_number"`
	State        string `json:"state"`
}

type Seller struct {
	Id         int64  `json:"id"`
}

type Payer struct {
	Address     Address `json:"address"`
	Name        string  `json:"name"`
	LastName    string  `json:"last_name"`
	PhoneNumber string  `json:"phone_number"`
}
