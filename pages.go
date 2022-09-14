package ofd

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"strings"
)

type Pages struct {
	PagesXml
	pwd string
	rc  *zip.ReadCloser
}
type PagesXml struct {
	Text string `xml:",chardata"`
	Page []struct {
		Text    string `xml:",chardata"`
		ID      string `xml:"ID,attr"`
		BaseLoc string `xml:"BaseLoc,attr"`
	} `xml:"Page"`
}

func (pages *Pages) GetFileContent(path string) ([]byte, error) {
	return LoadZipFileContent(pages.rc, path)
}
func (pages *Pages) GetPageById(pageId string) (*Page, error) {
	if strings.EqualFold(pageId, "") {
		return nil, fmt.Errorf("pageId is empty")
	}
	for _, item := range pages.Page {
		if !strings.EqualFold(item.ID, pageId) {
			continue
		}
		path := item.BaseLoc
		pos := strings.LastIndexByte(path, '/')
		pwd := path[0:pos]
		content, err := pages.GetFileContent(path)
		if err != nil {
			path = pages.pwd + "/" + item.BaseLoc
			pos := strings.LastIndexByte(path, '/')
			pwd = path[0:pos]
			content, err = pages.GetFileContent(path)
			if err != nil {
				return nil, err
			}
		}
		var page Page
		if err := xml.Unmarshal(content, &page); err != nil {
			return nil, err
		} else {
			page.pwd = pwd
			return &page, nil
		}
	}
	return nil, fmt.Errorf("pageId [%v] is invalid", pageId)

}

func (pages *Pages) GetPages() ([]Page, error) {
	var items []Page
	for _, item := range pages.Page {
		path := item.BaseLoc
		pos := strings.LastIndexByte(path, '/')
		pwd := path[0:pos]
		content, err := pages.GetFileContent(path)
		if err != nil {
			path = pages.pwd + "/" + item.BaseLoc
			pos := strings.LastIndexByte(path, '/')
			pwd = path[0:pos]
			content, err = pages.GetFileContent(path)
			if err != nil {
				return nil, err
			}
		}

		var page Page
		if err := xml.Unmarshal(content, &page); err != nil {
			return nil, err
		} else {
			page.pwd = pwd
		}
		items = append(items, page)
	}
	return items, nil
}
