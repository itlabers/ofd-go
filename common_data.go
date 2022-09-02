package ofd

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"strings"
)

type CommonDataXml struct {
	Text      string `xml:",chardata"`
	MaxUnitID struct {
		Text string `xml:",chardata"`
	} `xml:"MaxUnitID"`
	PageArea struct {
		Text        string `xml:",chardata"`
		PhysicalBox struct {
			Text string `xml:",chardata"`
		} `xml:"PhysicalBox"`
	} `xml:"PageArea"`
	PublicRes struct {
		Text string `xml:",chardata"`
	} `xml:"PublicRes"`
	DocumentRes struct {
		Text string `xml:",chardata"`
	} `xml:"DocumentRes"`
	TemplatePage struct {
		Text    string `xml:",chardata"`
		ID      string `xml:"ID,attr"`
		BaseLoc string `xml:"BaseLoc,attr"`
	} `xml:"TemplatePage"`
}
type CommonData struct {
	CommonDataXml
	pwd string
	rc  *zip.ReadCloser
}

func (commonData *CommonData) GetPublicRes() (*PublicRes, error) {
	res, err := commonData.getResource(PUBLICRES)
	if err != nil {
		return nil, err
	}
	return res.(*PublicRes), nil
}
func (commonData *CommonData) GetDocumentRes() (*DocumentRes, error) {
	res, err := commonData.getResource(DOCUMENTRES)
	if err != nil {
		return nil, err
	}
	return res.(*DocumentRes), nil
}
func (commonData *CommonData) GetTemplatePage() (*Page, error) {

	res, err := commonData.getResource(TEMPLATEPAGE)
	if err != nil {
		return nil, err
	}
	return res.(*Page), nil
}

type CommonDataResource int

const (
	PUBLICRES CommonDataResource = iota

	DOCUMENTRES

	TEMPLATEPAGE
)

func (commonData *CommonData) getFileContent(path string) ([]byte, error) {
	return LoadZipFileContent(commonData.rc, path)
}
func (commonData *CommonData) getResource(resType CommonDataResource) (interface{}, error) {
	var path string
	switch resType {
	case PUBLICRES:
		path = commonData.PublicRes.Text
	case DOCUMENTRES:
		path = commonData.DocumentRes.Text
	case TEMPLATEPAGE:
		path = commonData.TemplatePage.BaseLoc
	default:
		return nil, fmt.Errorf("resType[%v] is invalid", resType)
	}
	if strings.EqualFold(path, "") {
		return nil, fmt.Errorf("Resource is null")
	}
	if path[0] == '/' {
		path = commonData.pwd + path
	} else {
		path = commonData.pwd + "/" + path
	}
	pos := strings.LastIndexByte(path, '/')
	pwd := path[0:pos]
	content, err := commonData.getFileContent(path)
	if err != nil {
		return nil, err
	}
	switch resType {
	case PUBLICRES:
		var publicRes PublicRes
		err := xml.Unmarshal(content, &publicRes)
		if err != nil {
			return nil, err
		} else {
			publicRes.pwd = pwd
			publicRes.rc = commonData.rc
		}
		return &publicRes, nil
	case DOCUMENTRES:
		var documentRes DocumentRes
		err := xml.Unmarshal(content, &documentRes)
		if err != nil {
			return nil, err
		} else {
			documentRes.pwd = pwd
			documentRes.rc = commonData.rc
		}
		return &documentRes, nil

	case TEMPLATEPAGE:
		var page Page
		err := xml.Unmarshal(content, &page)
		if err != nil {
			return nil, err
		} else {
			page.pwd = pwd
		}
		return &page, nil
	default:
		return nil, fmt.Errorf("resType[%v] is invalid", resType)

	}
}
