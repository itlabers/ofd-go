package ofd

import (
	"encoding/xml"
	"testing"
)

const doc_res_content = `
<?xml version="1.0" encoding="UTF-8"?>
<ofd:Res
	xmlns:ofd="http://www.ofdspec.org/2016" BaseLoc="Res">
	<ofd:MultiMedias>
		<ofd:MultiMedia ID="63" Type="Image">
			<ofd:MediaFile>Image_0.png</ofd:MediaFile>
		</ofd:MultiMedia>
		<ofd:MultiMedia ID="65" Type="Image">
			<ofd:MediaFile>Image_1.png</ofd:MediaFile>
		</ofd:MultiMedia>
	</ofd:MultiMedias>
</ofd:Res>`

func Test_DocumentRes(t *testing.T) {
	var res DocumentRes
	if err := xml.Unmarshal([]byte(doc_res_content), &res); err != nil {
		t.Logf("%s", err)
	} else {
		t.Logf("%v", res)
	}
}
