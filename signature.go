package ofd

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"strings"

	"golang.org/x/crypto/cryptobyte"
)

type Signature struct {
	SignatureXml
	pwd string
	rc  *zip.ReadCloser
	Validator
	Category
	Content []byte
}
type SignatureXml struct {
	XMLName    xml.Name `xml:"Signature"`
	Text       string   `xml:",chardata"`
	Ofd        string   `xml:"ofd,attr"`
	SignedInfo struct {
		Text     string `xml:",chardata"`
		Provider struct {
			Text         string `xml:",chardata"`
			ProviderName string `xml:"ProviderName,attr"`
			Version      string `xml:"Version,attr"`
			Company      string `xml:"Company,attr"`
		} `xml:"Provider"`
		SignatureMethod struct {
			Text string `xml:",chardata"`
		} `xml:"SignatureMethod"`
		SignatureDateTime struct {
			Text string `xml:",chardata"`
		} `xml:"SignatureDateTime"`
		Parameters struct {
			Text      string `xml:",chardata"`
			Parameter struct {
				Text string `xml:",chardata"`
				Name string `xml:"Name,attr"`
			} `xml:"Parameter"`
		} `xml:"Parameters"`
		References struct {
			Text        string `xml:",chardata"`
			CheckMethod string `xml:"CheckMethod,attr"`
			Reference   []struct {
				Text       string `xml:",chardata"`
				FileRef    string `xml:"FileRef,attr"`
				CheckValue struct {
					Text string `xml:",chardata"`
				} `xml:"CheckValue"`
			} `xml:"Reference"`
		} `xml:"References"`
		StampAnnot struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"ID,attr"`
			PageRef  string `xml:"PageRef,attr"`
			Boundary string `xml:"Boundary,attr"`
		} `xml:"StampAnnot"`
	} `xml:"SignedInfo"`
	SignedValue struct {
		Text string `xml:",chardata"`
	} `xml:"SignedValue"`
}

func (signature Signature) GetFileContent(path string) ([]byte, error) {
	return LoadZipFileContent(signature.rc, path)
}
func (signature Signature) VerifyDigest() (bool, error) {
	checkMethod := signature.SignedInfo.References.CheckMethod
	if !(strings.EqualFold(checkMethod, SM3_OID) || strings.EqualFold(checkMethod, "sm3")) {
		return false, fmt.Errorf("oid(%s) not support ", checkMethod)
	}
	ref := signature.SignedInfo.References.Reference
	for _, value := range ref {
		filePath := value.FileRef
		if filePath[0] == '/' {
			filePath = filePath[1:]
		}
		content, err := signature.GetFileContent(filePath)
		if err != nil {
			break
		} else {
			validator := signature.Validator
			hashed := validator.Digest(content)
			encodeString := base64.StdEncoding.EncodeToString(hashed)
			if !strings.EqualFold(encodeString, value.CheckValue.Text) {
				return false, fmt.Errorf("file %s checkValue mistake", filePath)
			} else {
				continue
			}
		}
	}
	return true, nil

}
func (signature Signature) Verify() (bool, error) {
	switch signature.Category {
	case SIGN:
		return signature.verifySign()
	case SEAL:
		return signature.verifySeal()
	default:
		return signature.verifySeal()
	}
}

func (signature Signature) verifySign() (bool, error) {

	signatureMethod := signature.SignedInfo.SignatureMethod.Text
	if !strings.EqualFold(SM3WITHSM2_OID, signatureMethod) {
		return false, fmt.Errorf("oid(%s) not support ", signatureMethod)
	}

	signDatPath := signature.pwd + "/" + signature.SignedValue.Text
	content, err := signature.GetFileContent(signDatPath)

	if err != nil {
		return false, err
	}

	if contentInfo, err := NewContentInfo(cryptobyte.String(content)); err != nil {
		return false, err
	} else {
		content := contentInfo.GetContent()
		signedData, err := NewSignedData(content)
		if err != nil {
			return false, err
		}

		signerInfoes, err := signedData.GetSignerInfos()
		if err != nil {
			return false, err
		}
		cert := signedData.GetCertificates()
		for _, sign := range signerInfoes {
			msg := sign.GetAuthenticatedAttributes()
			encryptedDigest := sign.GetEncryptedDigest()
			validator := signature.Validator
			result, err := validator.Verify(cert, msg, encryptedDigest)
			if err != nil {
				break
			} else {
				fmt.Printf("%v  ", result)

			}
		}

	}
	return true, nil
}
func (signature Signature) verifySeal() (bool, error) {
	signatureMethod := signature.SignedInfo.SignatureMethod.Text
	if !strings.EqualFold(SM3WITHSM2_OID, signatureMethod) {
		return false, fmt.Errorf("oid(%s) not support ", signatureMethod)
	}

	signDatPath := signature.SignedValue.Text
	content, err := signature.GetFileContent(signDatPath)
	if err != nil {
		signDatPath = signature.pwd + "/" + signature.SignedValue.Text
		content, err = signature.GetFileContent(signDatPath)
		if err != nil {
			return false, err
		}

	}
	validator := signature.Validator
	if ses, err := New_SES_Signature(content); err != nil {
		return false, err
	} else {
		tbs_sign, err := ses.Get_TBS_Sign()
		if err != nil {
			return false, err
		} else {
			dataDash := validator.Digest(signature.Content)
			if !bytes.Equal(tbs_sign.Get_DataHash(), dataDash) {
				return false, fmt.Errorf("%s", "Signatures.xml Hash Value Mistake")
			}
			seal, err := tbs_sign.Get_Seal()
			if err != nil {
				return false, err
			}

			if _, err := validator.Verify(seal.cert, seal.eSealInfo, seal.signature.Bytes); err != nil {
				return false, err
			}

		}
		if _, err := validator.Verify(ses.cert, ses.to_sign, ses.signature.Bytes); err != nil {
			return false, err
		} else {
			return true, nil
		}
	}
}
