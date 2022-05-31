package app

import "encoding/xml"

type Product struct {
	Base
	XMLName  xml.Name  `xml:"product" json:"-"`
	Price    int64     `xml:"price" json:"price"`
	year     int       `xml:"year" json:"year"`
	Category *Category `xml:"category" json:"category"`
}

func (p *Product) SetProduct(id int, name string, price int64) {
	p.ID = id
	p.Name = name
	p.Price = price
}
