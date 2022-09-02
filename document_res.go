package ofd

import (
	"archive/zip"
	"encoding/xml"
)

type DocumentRes struct {
	DocumentResXml
	pwd string
	rc  *zip.ReadCloser
}
type DocumentResXml struct {
	XMLName     xml.Name `xml:"Res"`
	Text        string   `xml:",chardata"`
	Ofd         string   `xml:"ofd,attr"`
	BaseLoc     string   `xml:"BaseLoc,attr"`
	MultiMedias struct {
		Text       string `xml:",chardata"`
		MultiMedia []struct {
			Text      string `xml:",chardata"`
			ID        string `xml:"ID,attr"`
			Type      string `xml:"Type,attr"`
			MediaFile struct {
				Text string `xml:",chardata"`
			} `xml:"MediaFile"`
		} `xml:"MultiMedia"`
	} `xml:"MultiMedias"`
}

func (docRes *DocumentRes) GetFileContent(path string) ([]byte, error) {
	return LoadZipFileContent(docRes.rc, path)
}
