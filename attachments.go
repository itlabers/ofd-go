package ofd

import (
	"archive/zip"
	"encoding/xml"
)

type Attachments struct {
	AttachmentsXml
	pwd string
	rc  *zip.ReadCloser
}
type AttachmentsXml struct {
	XMLName    xml.Name `xml:"Attachments"`
	Text       string   `xml:",chardata"`
	Ofd        string   `xml:"ofd,attr"`
	Attachment []struct {
		Text         string `xml:",chardata"`
		ID           string `xml:"ID,attr"`
		Name         string `xml:"Name,attr"`
		Format       string `xml:"Format,attr"`
		Size         string `xml:"Size,attr"`
		Visible      string `xml:"Visible,attr"`
		CreationDate string `xml:"CreationDate,attr"`
	} `xml:"Attachment"`
}

func (attachments *Attachments) GetFileContent(path string) ([]byte, error) {
	return LoadZipFileContent(attachments.rc, path)

}
