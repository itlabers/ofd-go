package ofd

import (
	"encoding/xml"

	"testing"
)

const ofd_content = `
<?xml version="1.0" encoding="UTF-8"?>
<ofd:OFD
	xmlns:ofd="http://www.ofdspec.org/2016" Version="1.1">
	<ofd:DocBody>
		<ofd:DocInfo>
			<ofd:DocID>8132e5b821b249d8bde7a18642cca902</ofd:DocID>
		</ofd:DocInfo>
		<ofd:DocRoot>Doc_0/Document.xml</ofd:DocRoot>
		<ofd:Signatures>Doc_0/Signs/Signatures.xml</ofd:Signatures>
	</ofd:DocBody>
</ofd:OFD>
`

func TestOFD(t *testing.T) {
	var ofd OFD
	if err := xml.Unmarshal([]byte(ofd_content), &ofd); err != nil {
		t.Logf("%v", err)
	} else {
		t.Logf("%v", ofd)
	}
}
