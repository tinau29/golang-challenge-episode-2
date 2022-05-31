package app

import "encoding/xml"

type Category struct {
	Base
	XMLName xml.Name `xml:"category" json:"-"`
}

func (c *Category) SetCategory(id int, name string) {
	c.ID = id
	c.Name = name
}
