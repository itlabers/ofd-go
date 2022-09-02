package ofd

import (
	"encoding/xml"
	"testing"
)

const public_res_content = `
<?xml version="1.0" encoding="UTF-8"?>
<ofd:Res
	xmlns:ofd="http://www.ofdspec.org/2016" BaseLoc="Res">
	<ofd:Fonts>
		<ofd:Font ID="1" FontName="宋体"/>
		<ofd:Font ID="2" FontName="楷体"/>
		<ofd:Font ID="3" FontName="Courier New"/>
		<ofd:Font ID="8" FontName="楷体" Bold="true"/>
	</ofd:Fonts>
</ofd:Res>`

func Test_Res(t *testing.T) {
	var res PublicRes
	if err := xml.Unmarshal([]byte(public_res_content), &res); err != nil {
		t.Logf("%s", err)
	} else {
		t.Logf("%v", res)
	}
}
