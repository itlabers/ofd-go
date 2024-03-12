package ofd

import (
	"errors"
	"time"

	std "encoding/asn1"

	"golang.org/x/crypto/cryptobyte"
	"golang.org/x/crypto/cryptobyte/asn1"
)

type SES_Signature struct {
	to_sign        cryptobyte.String
	cert           []byte
	signatureAlgId std.ObjectIdentifier
	signature      std.BitString
	timestamp      std.BitString
}

func New_SES_Signature(input cryptobyte.String) (*SES_Signature, error) {
	var ses_signature cryptobyte.String
	input.ReadASN1(&ses_signature, asn1.SEQUENCE)

	var to_sign cryptobyte.String
	ses_signature.ReadASN1Element(&to_sign, asn1.SEQUENCE)

	var cert cryptobyte.String
	ses_signature.ReadASN1(&cert, asn1.OCTET_STRING)

	var signatureAlgId std.ObjectIdentifier
	ses_signature.ReadASN1ObjectIdentifier(&signatureAlgId)

	var signature std.BitString
	ses_signature.ReadASN1BitString(&signature)

	var signTime std.BitString
	ses_signature.ReadASN1BitString(&signTime)

	return &SES_Signature{
		to_sign:        to_sign,
		cert:           cert,
		signatureAlgId: signatureAlgId,
		signature:      signature,
		timestamp:      signTime,
	}, nil
}

func (ses *SES_Signature) Get_TBS_Sign() (*TBS_Sign, error) {
	to_sign, err := New_TBS_Sign(ses.to_sign)
	if err != nil {
		return nil, err
	}
	return to_sign, nil
}

func (ses *SES_Signature) Get_Cert() ([]byte, error) {
	return ses.cert, nil
}

func (ses *SES_Signature) Get_SignatureAlgId() string {
	return ses.signatureAlgId.String()
}

func (ses *SES_Signature) Get_Signature() []byte {
	return ses.signature.Bytes

}

type TBS_Sign struct {
	version      int64
	eSeal        cryptobyte.String // SESeal
	timeInfo     time.Time         // timeInfo
	dataHash     std.BitString     // dataHash
	propertyInfo string            // propertyInfo
	extDatas     cryptobyte.String // ExtensionDatas
}

func (tbs *TBS_Sign) Get_Version() int64 {
	return tbs.version
}
func (tbs *TBS_Sign) Get_DataHash() []byte {
	return tbs.dataHash.Bytes
}
func (tbs *TBS_Sign) Get_Seal() (*SESeal, error) {
	eSeal, err := New_SESeal(tbs.eSeal)
	if err != nil {
		return nil, err
	}
	return eSeal, nil
}

type ExtensionDatas struct {
	ExtnId    std.ObjectIdentifier
	Critical  bool
	ExtnValue []byte
}

func New_TBS_Sign(input cryptobyte.String) (*TBS_Sign, error) {
	var tbs_sign cryptobyte.String
	input.ReadASN1(&tbs_sign, asn1.SEQUENCE)

	var version int64
	tbs_sign.ReadASN1Int64WithTag(&version, asn1.INTEGER)

	var eSeal cryptobyte.String
	tbs_sign.ReadASN1(&eSeal, asn1.SEQUENCE)

	var timeInfo time.Time
	tbs_sign.ReadASN1GeneralizedTime(&timeInfo)

	var dataHash std.BitString
	tbs_sign.ReadASN1BitString(&dataHash)

	var propertyInfo cryptobyte.String
	tbs_sign.ReadASN1(&propertyInfo, asn1.IA5String)

	var extDatas cryptobyte.String
	tbs_sign.ReadASN1(&extDatas, asn1.SEQUENCE)

	return &TBS_Sign{
		version:      version,
		eSeal:        eSeal,
		timeInfo:     timeInfo,
		dataHash:     dataHash,
		propertyInfo: string(propertyInfo),
		extDatas:     extDatas,
	}, nil
}

type SESeal struct {
	eSealInfo      cryptobyte.String    //SES_Seal_Info
	cert           []byte               //cert
	signatureAlgId std.ObjectIdentifier //signatureAlgId
	signature      std.BitString        //signature
}

func New_SESeal(seSeal cryptobyte.String) (*SESeal, error) {
	var eSealInfo cryptobyte.String
	seSeal.ReadASN1Element(&eSealInfo, asn1.SEQUENCE)

	var cert cryptobyte.String
	seSeal.ReadASN1(&cert, asn1.OCTET_STRING)

	var signatureAlgId std.ObjectIdentifier
	seSeal.ReadASN1ObjectIdentifier(&signatureAlgId)

	var signature std.BitString
	seSeal.ReadASN1BitString(&signature)

	return &SESeal{
		eSealInfo:      eSealInfo,
		cert:           cert,
		signatureAlgId: signatureAlgId,
		signature:      signature,
	}, nil
}

func (seal *SESeal) Get_SES_SealInfo() (*SES_Seal_Info, error) {
	eSeal, err := New_SES_Seal_Info(seal.eSealInfo)
	if err != nil {
		return nil, err
	}
	return eSeal, nil
}
func (seal *SESeal) Get_Cert() ([]byte, error) {
	return seal.cert, nil
}

func (seal *SESeal) Get_SignatureAlgId() string {
	return seal.signatureAlgId.String()
}

func (seal *SESeal) Get_Signature() []byte {
	return seal.signature.Bytes

}

func New_SES_Seal_Info(input cryptobyte.String) (*SES_Seal_Info, error) {
	var eSealInfo cryptobyte.String
	input.ReadASN1(&eSealInfo, asn1.SEQUENCE)
	var header cryptobyte.String
	eSealInfo.ReadASN1(&header, asn1.SEQUENCE)
	var esID cryptobyte.String
	eSealInfo.ReadASN1(&esID, asn1.IA5String)

	var property cryptobyte.String
	eSealInfo.ReadASN1(&property, asn1.SEQUENCE)

	var pictrue cryptobyte.String
	eSealInfo.ReadASN1(&pictrue, asn1.SEQUENCE)

	var extDatas cryptobyte.String
	eSealInfo.ReadASN1(&extDatas, asn1.SEQUENCE)

	return &SES_Seal_Info{
		header:   header,
		esID:     string(esID),
		property: property,
		pictrue:  pictrue,
		extDatas: extDatas,
	}, nil
}

type SES_Seal_Info struct {
	header   cryptobyte.String //SES_Header
	esID     string            //esID
	property cryptobyte.String //SES_ESPropertyInfo
	pictrue  cryptobyte.String //SES_ESPictrueInfo
	extDatas cryptobyte.String //ExtensionDatas
}

func (ses_seal_info *SES_Seal_Info) Get_ESID() string {
	return ses_seal_info.esID
}

func (ses_seal_info *SES_Seal_Info) Get_Header() (*SES_Header, error) {
	ses_hender, err := New_SES_Header(ses_seal_info.header)
	if err != nil {
		return nil, err
	}
	return ses_hender, nil
}
func (ses_seal_info *SES_Seal_Info) Get_Property() (*SES_ESPropertyInfo, error) {
	ses_property, err := New_SES_ESPropertyInfo(ses_seal_info.property)
	if err != nil {
		return nil, err
	}
	return ses_property, nil
}
func (ses_seal_info *SES_Seal_Info) Get_Pictrue() (*SES_ESPictrueInfo, error) {
	ses_picture, err := New_SES_ESPictrueInfo(ses_seal_info.pictrue)
	if err != nil {
		return nil, err
	}
	return ses_picture, nil
}
func (ses_seal_info *SES_Seal_Info) Get_ExtDatas() (*ExtensionDatas, error) {
	extDatas, err := New_ExtensionDatas(ses_seal_info.extDatas)
	if err != nil {
		return nil, err
	}
	return extDatas, nil
}

func New_ExtensionDatas(extDatas cryptobyte.String) (*ExtensionDatas, error) {
	var extnId std.ObjectIdentifier
	extDatas.ReadASN1ObjectIdentifier(&extnId)

	var critical bool
	extDatas.ReadASN1Boolean(&critical)

	var extnValue cryptobyte.String

	extDatas.ReadASN1(&extnValue, asn1.OCTET_STRING)
	return &ExtensionDatas{
		ExtnId:    extnId,
		Critical:  critical,
		ExtnValue: extnValue,
	}, nil
}

func New_SES_ESPropertyInfo(ses_property cryptobyte.String) (*SES_ESPropertyInfo, error) {
	var seal_type int64
	ses_property.ReadASN1Integer(&seal_type)

	var name cryptobyte.String
	ses_property.ReadASN1(&name, asn1.UTF8String)

	var certListType int64
	ses_property.ReadASN1Integer(&certListType)

	var certList cryptobyte.String
	ses_property.ReadASN1(&certList, asn1.SEQUENCE)

	var createDate time.Time
	ses_property.ReadASN1GeneralizedTime(&createDate)

	var validStart time.Time
	ses_property.ReadASN1GeneralizedTime(&validStart)

	var validEnd time.Time
	ses_property.ReadASN1GeneralizedTime(&validEnd)

	return &SES_ESPropertyInfo{
		Type:         seal_type,
		Name:         string(name),
		CertListType: certListType,
		CertList:     certList,
		CreateDate:   createDate,
		ValidStart:   validStart,
		ValidEnd:     validEnd,
	}, nil
}
func New_SES_ESPictrueInfo(pictrue cryptobyte.String) (*SES_ESPictrueInfo, error) {
	var pic_type cryptobyte.String
	pictrue.ReadASN1(&pic_type, asn1.IA5String)

	var data cryptobyte.String
	pictrue.ReadASN1(&data, asn1.OCTET_STRING)

	var width int64
	pictrue.ReadASN1Integer(&width)

	var height int64
	pictrue.ReadASN1Integer(&height)

	return &SES_ESPictrueInfo{
		Type:   string(pic_type),
		Data:   data,
		Width:  width,
		Height: height,
	}, nil
}

type SES_Header struct {
	ID      string     //id
	Version int64      //version 
	Vid     string     //vid
}

func New_SES_Header(ses_header cryptobyte.String) (*SES_Header, error) {
	var id cryptobyte.String
	ses_header.ReadASN1(&id, asn1.IA5String)

	var verison int64
	ses_header.ReadASN1Integer(&verison)

	var vID cryptobyte.String
	ses_header.ReadASN1(&vID, asn1.IA5String)
	return &SES_Header{
		ID:      string(id),
		Version: verison,
		Vid:     string(vID),
	}, nil
}

type SES_ESPropertyInfo struct {
	Type         int64
	Name         string
	CertListType int64
	CertList     cryptobyte.String
	CreateDate   time.Time
	ValidStart   time.Time
	ValidEnd     time.Time
}

func (ses_property_info *SES_ESPropertyInfo) Get_CertList() (*SES_CertList, error) {
	ses_certlist, err := New_SES_CertList(ses_property_info.CertListType, ses_property_info.CertList)
	if err != nil {
		return nil, err
	} else {
		return ses_certlist, nil
	}

}

type SES_CertList struct {
	certs          cryptobyte.String
	certDigestList cryptobyte.String
}

func New_SES_CertList(cert_type int64, cert_list cryptobyte.String) (*SES_CertList, error) {
	if cert_type == 1 {
		var certs cryptobyte.String
		cert_list.ReadASN1(&certs, asn1.SEQUENCE)
		return &SES_CertList{
			certs:          certs,
			certDigestList: nil,
		}, nil
	} else if cert_type == 2 {
		var certDigestList cryptobyte.String
		cert_list.ReadASN1(&certDigestList, asn1.SEQUENCE)
		return &SES_CertList{
			certs:          nil,
			certDigestList: certDigestList,
		}, nil
	} else {
		return nil, errors.New("type is not valid")
	}
}

type Cert []byte

type CertInfoList struct {
	Certs []Cert
}

func (ses_certlist *SES_CertList) Get_Certs() []byte {
	return ses_certlist.certs
}

func (ses_certlist *SES_CertList) Get_CertDigestList() []byte {
	return ses_certlist.certDigestList
}

type CertDigestList struct {
	CertDigestObjs []CertDigestObj
}

type ObjType string
type CertDigestValue []byte

type CertDigestObj struct {
	Type  ObjType
	Value CertDigestValue
}

type SES_ESPictrueInfo struct {
	Type   string
	Data   []byte
	Width  int64
	Height int64
}
