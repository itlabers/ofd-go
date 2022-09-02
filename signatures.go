package ofd

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"strings"
)

type SignaturesXml struct {
	XMLName   xml.Name `xml:"Signatures"`
	Text      string   `xml:",chardata"`
	Ofd       string   `xml:"ofd,attr"`
	MaxSignId struct {
		Text string `xml:",chardata"`
	} `xml:"MaxSignId"`
	Signature []struct {
		Text    string `xml:",chardata"`
		ID      string `xml:"ID,attr"`
		Type    string `xml:"Type,attr"`
		BaseLoc string `xml:"BaseLoc,attr"`
	} `xml:"Signature"`
}

type Signatures struct {
	SignaturesXml
	pwd string
	rc  *zip.ReadCloser
	Validator
}

func (signatures Signatures) GetFileContent(path string) ([]byte, error) {
	return LoadZipFileContent(signatures.rc, path)
}
func (signatures Signatures) GetSignatureById(signId string) (*Signature, error) {
	if strings.EqualFold(signId, "") {
		return nil, fmt.Errorf("signId is empty")
	}
	for _, sign := range signatures.Signature {
		if !strings.EqualFold(signId, sign.ID) {
			continue
		}
		path := sign.BaseLoc
		pos := strings.LastIndexByte(path, '/')
		pwd := path[0:pos]
		content, err := signatures.GetFileContent(path)
		if err != nil {
			if path[0] == '/' {
				path = signatures.pwd + path
			} else {
				path = signatures.pwd + "/" + path
			}
			pos = strings.LastIndexByte(path, '/')
			pwd = path[0:pos]
			content, err = signatures.GetFileContent(path)
			if err != nil {
				return nil, err
			}
		}

		var signature Signature
		if err := xml.Unmarshal(content, &signature); err != nil {
			return nil, err
		} else {
			signature.pwd = pwd
			signature.rc = signatures.rc
			signature.Validator = signatures.Validator
			signature.Content = content
			switch sign.Type {
			case "Seal":
				signature.Category = SEAL
			case "Sign":
				signature.Category = SIGN
			default:
				signature.Category = SEAL
			}
		}
		return &signature, nil
	}
	return nil, fmt.Errorf("signId [%v] not found", signId)
}
