package ofd

import "encoding/xml"

type OFD struct {
	XMLName xml.Name `xml:"OFD"`
	Text    string   `xml:",chardata"`
	Ofd     string   `xml:"ofd,attr"`
	Version string   `xml:"Version,attr"`
	DocBody []struct {
		Text    string `xml:",chardata"`
		DocInfo struct {
			Text  string `xml:",chardata"`
			DocID struct {
				Text string `xml:",chardata"`
			} `xml:"DocID"`
		} `xml:"DocInfo"`
		DocRoot struct {
			Text string `xml:",chardata"`
		} `xml:"DocRoot"`
		Signatures struct {
			Text string `xml:",chardata"`
		} `xml:"Signatures"`
	} `xml:"DocBody"`
}
