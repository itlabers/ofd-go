package ofd

import (
	std "encoding/asn1"

	"golang.org/x/crypto/cryptobyte"
	"golang.org/x/crypto/cryptobyte/asn1"
)

type ContentInfo struct {
	contentType std.ObjectIdentifier
	content     cryptobyte.String
}

func NewContentInfo(input cryptobyte.String) (*ContentInfo, error) {

	var contentInfo cryptobyte.String
	input.ReadASN1(&contentInfo, asn1.SEQUENCE)

	var contentType std.ObjectIdentifier
	contentInfo.ReadASN1ObjectIdentifier(&contentType)

	var content cryptobyte.String
	var tag asn1.Tag
	contentInfo.ReadAnyASN1(&content, &tag)
	return &ContentInfo{
		contentType: contentType,
		content:     content,
	}, nil
}

func (contentInfo *ContentInfo) GetContent() []byte {
	return contentInfo.content
}

type SignedData struct {
	version                             int64
	digestAlgorithms                    cryptobyte.String
	contentInfo                         cryptobyte.String
	certificates                        cryptobyte.String
	extendedCertificatesAndCertificates cryptobyte.String
	crls                                cryptobyte.String
	signerInfos                         cryptobyte.String
}

func NewSignedData(content cryptobyte.String) (*SignedData, error) {

	var signedData cryptobyte.String
	content.ReadASN1(&signedData, asn1.SEQUENCE)
	var version int64
	signedData.ReadASN1Integer(&version)

	var digestAlgorithms cryptobyte.String
	signedData.ReadASN1(&digestAlgorithms, asn1.SET)

	var contentInfo cryptobyte.String
	signedData.ReadASN1(&contentInfo, asn1.SEQUENCE)

	var certificates cryptobyte.String
	signedData.ReadASN1(&certificates, asn1.SEQUENCE)

	var extendedCertificatesAndCertificates cryptobyte.String
	var exit bool
	signedData.ReadOptionalASN1(&extendedCertificatesAndCertificates, &exit, asn1.SEQUENCE)

	var crls cryptobyte.String
	signedData.ReadOptionalASN1(&crls, &exit, asn1.SEQUENCE)

	var signerInfos cryptobyte.String
	signedData.ReadASN1(&signerInfos, asn1.SET)

	return &SignedData{
		version:                             version,
		digestAlgorithms:                    digestAlgorithms,
		contentInfo:                         content,
		certificates:                        certificates,
		extendedCertificatesAndCertificates: extendedCertificatesAndCertificates,
		crls:                                crls,
		signerInfos:                         signerInfos,
	}, nil
}
func (signedData *SignedData) GetCertificates() []byte {
	return signedData.certificates
}
func (signedData *SignedData) GetSignerInfos() ([]SignerInfo, error) {
	var signerInfoSet []SignerInfo
	signerInfos := signedData.signerInfos

	exist := true
	for exist {
		var signerInfo cryptobyte.String
		exist = signerInfos.ReadASN1(&signerInfo, asn1.SEQUENCE)
		if exist {
			item, err := NewSignerInfo(signerInfo)
			if err != nil {
				break
			}
			signerInfoSet = append(signerInfoSet, *item)
		}
	}
	return signerInfoSet, nil
}

type SignerInfo struct {
	version                   int64
	issuerAndSerialNumber     cryptobyte.String
	digestAlgorithm           cryptobyte.String
	authenticatedAttributes   cryptobyte.String
	digestEncryptionAlgorithm cryptobyte.String
	encryptedDigest           cryptobyte.String
	unauthenticatedAttributes cryptobyte.String
}

func NewSignerInfo(signInfo cryptobyte.String) (*SignerInfo, error) {
	var version int64
	signInfo.ReadASN1Integer(&version)
	var issuerAndSerialNumber cryptobyte.String
	signInfo.ReadASN1(&issuerAndSerialNumber, asn1.SEQUENCE)
	var digestAlgorithm cryptobyte.String
	signInfo.ReadASN1(&digestAlgorithm, asn1.SEQUENCE)
	var authenticatedAttributes cryptobyte.String
	var tag asn1.Tag
	signInfo.ReadAnyASN1Element(&authenticatedAttributes, &tag)

	var digestEncryptionAlgorithm cryptobyte.String
	signInfo.ReadASN1(&digestEncryptionAlgorithm, asn1.SEQUENCE)

	var encryptedDigest cryptobyte.String
	signInfo.ReadASN1(&encryptedDigest, asn1.SEQUENCE)

	var unauthenticatedAttributes cryptobyte.String
	var exit bool
	signInfo.ReadOptionalASN1(&unauthenticatedAttributes, &exit, asn1.SEQUENCE)

	return &SignerInfo{
		version:                   version,
		issuerAndSerialNumber:     issuerAndSerialNumber,
		digestAlgorithm:           digestAlgorithm,
		authenticatedAttributes:   authenticatedAttributes,
		digestEncryptionAlgorithm: digestEncryptionAlgorithm,
		encryptedDigest:           encryptedDigest,
		unauthenticatedAttributes: unauthenticatedAttributes,
	}, nil
}

func (signerInfo *SignerInfo) GetVersion() int64 {
	return signerInfo.version
}
func (signerInfo *SignerInfo) GetIssuerAndSerialNumber() cryptobyte.String {
	return signerInfo.issuerAndSerialNumber
}
func (signerInfo *SignerInfo) GetDigestAlgorithm() string {
	digestAlgorithm := signerInfo.digestAlgorithm
	var digestAlgorithmOid std.ObjectIdentifier
	digestAlgorithm.ReadASN1ObjectIdentifier(&digestAlgorithmOid)
	return digestAlgorithmOid.String()
}

func (signerInfo *SignerInfo) GetAuthenticatedAttributes() cryptobyte.String {
	authenticatedAttributes := signerInfo.authenticatedAttributes
	tag := authenticatedAttributes[0]
	if tag == 0xA0 {
		authenticatedAttributes[0] = 0x31
	}
	return authenticatedAttributes
}

func (signerInfo *SignerInfo) GetDigestEncryptionAlgorithm() string {
	digestEncryptionAlgorithm := signerInfo.digestEncryptionAlgorithm
	var digestEncryptionAlgorithmOid std.ObjectIdentifier
	digestEncryptionAlgorithm.ReadASN1ObjectIdentifier(&digestEncryptionAlgorithmOid)
	return digestEncryptionAlgorithmOid.String()
}

func (signerInfo *SignerInfo) GetEncryptedDigest() cryptobyte.String {
	return signerInfo.encryptedDigest
}

func (signerInfo *SignerInfo) GetUnauthenticatedAttributes() cryptobyte.String {
	unauthenticatedAttributes := signerInfo.unauthenticatedAttributes
	tag := unauthenticatedAttributes[0]
	if tag == 0xA1 {
		unauthenticatedAttributes[0] = 0x31
	}
	return unauthenticatedAttributes
}
