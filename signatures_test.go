package ofd

import (
	"encoding/xml"
	"testing"
)

const signatures_content = `
<?xml version="1.0" encoding="UTF-8"?>
<ofd:Signatures
	xmlns:ofd="http://www.ofdspec.org/2016">
	<ofd:MaxSignId>2</ofd:MaxSignId>
	<ofd:Signature ID="2" BaseLoc="/Doc_0/Signs/Sign_0/Signature.xml"/>
</ofd:Signatures>
`

func TestSignatures(t *testing.T) {
	var signature Signatures
	if err := xml.Unmarshal([]byte(signatures_content), &signature); err != nil {
		t.Logf("%v", err)
	} else {
		t.Logf("%v", signature.Signature)
	}
}
