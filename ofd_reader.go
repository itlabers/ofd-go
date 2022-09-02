package ofd

import (
	"archive/zip"
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
)

const (
	OFD_PATH = "OFD.xml"
)

type OFDReader struct {
	validator Validator
	filepath  string
	rc        *zip.ReadCloser
	ofd       *OFD
}

type Option func(opt *Options)

type Options struct {
	validator Validator
}

var defaultOption = Options{
	validator: &CommonValidator{},
}

func WithValidator(validator Validator) Option {
	return func(opt *Options) {
		opt.validator = validator
	}
}

func NewOFDReader(path string, opts ...Option) (*OFDReader, error) {
	opt := defaultOption
	for _, o := range opts {
		o(&opt)
	}
	rc, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	return &OFDReader{
		rc:        rc,
		validator: opt.validator,
		filepath:  path,
	}, nil
}
func (ofdReader *OFDReader) Close() error {
	if ofdReader.rc != nil {
		return ofdReader.rc.Close()
	}
	return nil
}

func (ofdReader *OFDReader) OFD() (*OFD, error) {
	if ofdReader.ofd == nil {
		content, err := ofdReader.GetFileContent(OFD_PATH)
		if err != nil {
			return nil, err
		}
		var ofd OFD
		if err := xml.Unmarshal(content, &ofd); err != nil {
			return nil, err
		} else {
			ofdReader.ofd = &ofd
			return &ofd, nil
		}
	} else {
		return ofdReader.ofd, nil
	}
}

type Resource int

const (
	DOCUMENT Resource = iota
	SIGNATURES
)

func (ofdReader *OFDReader) getResourceById(docId string, resType Resource) (interface{}, error) {
	if docId == "" {
		return nil, errors.New("docId is empty")
	}
	ofd, err := ofdReader.OFD()
	if err != nil {
		return nil, err
	}
	docs := ofd.DocBody
	var path string
	for _, doc := range docs {
		if !strings.EqualFold(doc.DocInfo.DocID.Text, docId) {
			continue
		}
		switch resType {
		case DOCUMENT:
			path = doc.DocRoot.Text
		case SIGNATURES:
			path = doc.Signatures.Text
		default:
			return nil, fmt.Errorf("resType[%v] is invalid", resType)
		}
		if path[0] == '/' {
			path = path[1:]
		}
		pos := strings.LastIndexByte(path, '/')
		pwd := path[0:pos]
		content, err := ofdReader.GetFileContent(path)
		if err != nil {
			return nil, err
		}
		switch resType {
		case DOCUMENT:
			var doc Document
			if err := xml.Unmarshal(content, &doc); err != nil {
				return nil, err
			} else {
				doc.pwd = pwd
				doc.rc = ofdReader.rc
			}
			return &doc, nil
		case SIGNATURES:
			var signatures Signatures
			if err := xml.Unmarshal(content, &signatures); err != nil {
				return nil, err
			} else {
				signatures.pwd = pwd
				signatures.rc = ofdReader.rc
			}
			return &signatures, nil
		default:
			return nil, fmt.Errorf("resType[%v] is invalid", resType)

		}
	}
	return nil, fmt.Errorf("docId [%v] not found", docId)
}
func (ofdReader *OFDReader) GetDocumentById(docId string) (*Document, error) {
	res, err := ofdReader.getResourceById(docId, DOCUMENT)
	if err != nil {
		return nil, err
	}
	return res.(*Document), nil
}

func (ofdReader *OFDReader) GetSignaturesById(docId string) (*Signatures, error) {
	res, err := ofdReader.getResourceById(docId, SIGNATURES)
	if err != nil {
		return nil, err
	}
	signatures := res.(*Signatures)
	signatures.Validator = ofdReader.validator
	return signatures, nil
}
func (ofdReader *OFDReader) GetFileContent(path string) ([]byte, error) {
	return LoadZipFileContent(ofdReader.rc, path)
}
