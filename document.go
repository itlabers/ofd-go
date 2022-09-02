package ofd

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"strings"
)

type DocumentXml struct {
	XMLName     xml.Name   `xml:"Document"`
	Text        string     `xml:",chardata"`
	Ofd         string     `xml:"ofd,attr"`
	CommonData  CommonData `xml:"CommonData"`
	Pages       Pages      `xml:"Pages"`
	Attachments struct {
		Text string `xml:",chardata"`
	} `xml:"Attachments"`
	CustomTags struct {
		Text string `xml:",chardata"`
	} `xml:"CustomTags"`
}
type Document struct {
	DocumentXml
	pwd string
	rc  *zip.ReadCloser
}

func (doc *Document) GetCommonData() (*CommonData, error) {
	commonData := doc.CommonData
	commonData.rc = doc.rc
	commonData.pwd = doc.pwd
	return &commonData, nil
}
func (doc *Document) GetPages() (*Pages, error) {
	pages := doc.Pages
	pages.rc = doc.rc
	pages.pwd = doc.pwd
	return &pages, nil
}

func (doc *Document) GetAttachments() (*Attachments, error) {
	res, err := doc.getResource(ATTACHMENTS)
	if err != nil {
		return nil, err
	}
	return res.(*Attachments), nil
}

func (doc *Document) GetCustomTags() (*CustomTags, error) {
	res, err := doc.getResource(CUSTOMTAGS)
	if err != nil {
		return nil, err
	}
	return res.(*CustomTags), nil
}

type DocResource int

const (
	ATTACHMENTS DocResource = iota
	CUSTOMTAGS
)

func (document *Document) GetFileContent(path string) ([]byte, error) {
	return LoadZipFileContent(document.rc, path)
}
func (document *Document) getResource(resType DocResource) (interface{}, error) {
	var path string
	switch resType {
	case ATTACHMENTS:
		path = document.Attachments.Text
	case CUSTOMTAGS:
		path = document.CustomTags.Text
	default:
		return nil, fmt.Errorf("resType[%v] is invalid", resType)
	}
	if strings.EqualFold(path, "") {
		return nil, fmt.Errorf("Resource is null")
	}
	pos := strings.LastIndexByte(path, '/')
	pwd := path[0:pos]
	content, err := document.GetFileContent(path)
	if err != nil {
		if path[0] == '/' {
			path = document.pwd + path
		} else {
			path = document.pwd + "/" + path
		}

		pos := strings.LastIndexByte(path, '/')
		pwd = path[0:pos]
		content, err = document.GetFileContent(path)
		if err != nil {
			return nil, err
		}
	}

	switch resType {
	case ATTACHMENTS:
		var attachments Attachments
		if err := xml.Unmarshal(content, &attachments); err != nil {
			return nil, err
		} else {

			attachments.pwd = pwd
			attachments.rc = document.rc
		}
		return &attachments, nil
	case CUSTOMTAGS:
		var customTag CustomTags
		if err := xml.Unmarshal(content, &customTag); err != nil {
			return nil, err
		} else {
			customTag.pwd = pwd
			customTag.rc = document.rc
		}
		return &customTag, nil
	default:
		return nil, fmt.Errorf("resType[%v] is invalid", resType)

	}
}
