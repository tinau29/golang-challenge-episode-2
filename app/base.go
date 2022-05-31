package app

import "encoding/xml"

type Base struct {
	XMLName xml.Name `xml:"base" json:"-"`
	ID      int      `xml:"id" json:"id, omitempty"`
	Name    string   `xml:"name" json:"name, omitempty"`
}
