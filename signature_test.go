package ofd

import (
	"encoding/xml"
	"testing"
)

const signature_content = `
<?xml version="1.0" encoding="UTF-8"?>
<ofd:Signature
	xmlns:ofd="http://www.ofdspec.org/2016">
	<ofd:SignedInfo>
		<ofd:Provider ProviderName="BaiWang_OES" Version="1.0" Company=""/>
		<ofd:SignatureMethod>1.2.156.10197.1.501</ofd:SignatureMethod>
		<ofd:SignatureDateTime>20220328071354Z</ofd:SignatureDateTime>
		<ofd:Parameters>
			<ofd:Parameter Name="Protect_OFD">none</ofd:Parameter>
		</ofd:Parameters>
		<ofd:References CheckMethod="1.2.156.10197.1.401">
			<ofd:Reference FileRef="/Doc_0/Pages/Page_0/Content.xml">
				<ofd:CheckValue>e70+E5QyBH4wRV4xET9joIjuIIFQUGE35sF1y8a3mSA=</ofd:CheckValue>
			</ofd:Reference>
			<ofd:Reference FileRef="/Doc_0/DocumentRes_0.xml">
				<ofd:CheckValue>Xkbf6Pfn02s8RdWAfYiE9VmVpw3+MgFuEEEEt5EE+Vs=</ofd:CheckValue>
			</ofd:Reference>
			<ofd:Reference FileRef="/Doc_0/PublicRes_0.xml">
				<ofd:CheckValue>Sds4LG5RrkFWRR9909cHfWF+mAiyQclIGAjJUgHOu24=</ofd:CheckValue>
			</ofd:Reference>
			<ofd:Reference FileRef="/Doc_0/Attachs/Attachments.xml">
				<ofd:CheckValue>KXDYVC+KVyIOox46/9u174/Y0a1dHEnqWDIqxIDn7cI=</ofd:CheckValue>
			</ofd:Reference>
			<ofd:Reference FileRef="/Doc_0/Attachs/bker_issuer_20191231_C10303110004552019030390296600243000000000019444.xml">
				<ofd:CheckValue>VrM59lM2wAZ5KYm3sgvjo3wQCZ3UdE99Spj2MHDIM0w=</ofd:CheckValue>
			</ofd:Reference>
			<ofd:Reference FileRef="/Doc_0/Tags/CustomTags.xml">
				<ofd:CheckValue>x9wTJfjt1nnSZGKtNOKF8ipYFiI+YR2jQ0DM29axKv4=</ofd:CheckValue>
			</ofd:Reference>
			<ofd:Reference FileRef="/Doc_0/Tags/Tag_kj.xml">
				<ofd:CheckValue>qEGum3RvTwYlzJPGFkkolHEsV6FsNUjb2dLAtyeEkr0=</ofd:CheckValue>
			</ofd:Reference>
			<ofd:Reference FileRef="/Doc_0/Document.xml">
				<ofd:CheckValue>AROs4wgy8qPc/RYElgh6q9wC74PZx+eeTlHhqTztch8=</ofd:CheckValue>
			</ofd:Reference>
		</ofd:References>
		<ofd:StampAnnot ID="1" PageRef="6" Boundary="153.00 50.00 40.00 33.00"/>
	</ofd:SignedInfo>
	<ofd:SignedValue>/Doc_0/Signs/Sign_0/SignedValue.dat</ofd:SignedValue>
</ofd:Signature>
`

func TestSignature(t *testing.T) {
	var signature Signature
	if err := xml.Unmarshal([]byte(signature_content), &signature); err != nil {
		t.Logf("%v", err)
	} else {
		t.Logf("%v", signature.SignedInfo.References)
	}
}
