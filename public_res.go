package ofd

import (
	"archive/zip"
	"encoding/xml"
)

type PublicRes struct {
	PublicResXml
	pwd string
	rc  *zip.ReadCloser
}
type PublicResXml struct {
	XMLName xml.Name `xml:"Res"`
	Text    string   `xml:",chardata"`
	Ofd     string   `xml:"ofd,attr"`
	BaseLoc string   `xml:"BaseLoc,attr"`
	Fonts   struct {
		Text string `xml:",chardata"`
		Font []struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"ID,attr"`
			FontName string `xml:"FontName,attr"`
			Bold     string `xml:"Bold,attr"`
		} `xml:"Font"`
	} `xml:"Fonts"`
}

func (publicRes *PublicRes) GetFileContent(path string) ([]byte, error) {
	return LoadZipFileContent(publicRes.rc, path)
}
