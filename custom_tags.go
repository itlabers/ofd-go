package ofd

import (
	"archive/zip"
	"encoding/xml"
)

type CustomTags struct {
	CustomTagsXml
	pwd string
	rc  *zip.ReadCloser
}

type CustomTagsXml struct {
	XMLName   xml.Name `xml:"CustomTags"`
	Text      string   `xml:",chardata"`
	Ofd       string   `xml:"ofd,attr"`
	CustomTag struct {
		Text    string `xml:",chardata"`
		TypeID  string `xml:"TypeID,attr"`
		FileLoc struct {
			Text string `xml:",chardata"`
		} `xml:"FileLoc"`
	} `xml:"CustomTag"`
}
func (attachments *CustomTags) GetFileContent(path string) ([]byte, error) {
	return LoadZipFileContent(attachments.rc, path)

}