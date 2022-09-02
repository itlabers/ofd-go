package ofd

import "encoding/xml"

type Page struct {
	PageXml
	pwd  string
}
type PageXml struct {
	XMLName xml.Name `xml:"Page"`
	Text    string   `xml:",chardata"`
	Ofd     string   `xml:"ofd,attr"`
	Area    struct {
		Text        string `xml:",chardata"`
		PhysicalBox struct {
			Text string `xml:",chardata"`
		} `xml:"PhysicalBox"`
	} `xml:"Area"`
	Content struct {
		Text  string `xml:",chardata"`
		Layer struct {
			Text       string `xml:",chardata"`
			ID         string `xml:"ID,attr"`
			TextObject []struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"ID,attr"`
				Boundary string `xml:"Boundary,attr"`
				Font     string `xml:"Font,attr"`
				Size     string `xml:"Size,attr"`
				Weight   string `xml:"Weight,attr"`
				TextCode struct {
					Text   string `xml:",chardata"`
					X      string `xml:"X,attr"`
					Y      string `xml:"Y,attr"`
					DeltaX string `xml:"DeltaX,attr"`
				} `xml:"TextCode"`
			} `xml:"TextObject"`
			PathObject []struct {
				Text            string `xml:",chardata"`
				ID              string `xml:"ID,attr"`
				Boundary        string `xml:"Boundary,attr"`
				LineWidth       string `xml:"LineWidth,attr"`
				AbbreviatedData struct {
					Text string `xml:",chardata"`
				} `xml:"AbbreviatedData"`
			} `xml:"PathObject"`
			ImageObject []struct {
				Text       string `xml:",chardata"`
				ID         string `xml:"ID,attr"`
				Boundary   string `xml:"Boundary,attr"`
				CTM        string `xml:"CTM,attr"`
				ResourceID string `xml:"ResourceID,attr"`
			} `xml:"ImageObject"`
		} `xml:"Layer"`
	} `xml:"Content"`
}
