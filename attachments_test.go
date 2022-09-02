package ofd

import (
	"encoding/xml"
	"testing"
)

const attachment_content = `
<?xml version="1.0" encoding="UTF-8"?>
<ofd:Attachments
	xmlns:ofd="http://www.ofdspec.org/2016">
	<ofd:Attachment ID="66" Name="bker_issuer_20191231_C10303110004552019030390296600243000000000019444" Format="xml" Size="3.103516" Visible="true" CreationDate="2022-03-28T07:13:54">bker_issuer_20191231_C10303110004552019030390296600243000000000019444.xml</ofd:Attachment>
</ofd:Attachments>
`

func Test_Attachments(t *testing.T) {
	var res Attachments
	if err := xml.Unmarshal([]byte(attachment_content), &res); err != nil {
		t.Logf("%s", err)
	} else {
		t.Logf("%v", res)
	}
}
