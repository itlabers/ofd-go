package ofd

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"
)

const document_content = `
<?xml version="1.0" encoding="UTF-8"?>
<ofd:Document
	xmlns:ofd="http://www.ofdspec.org/2016">
	<ofd:CommonData>
		<ofd:MaxUnitID>66</ofd:MaxUnitID>
		<ofd:PageArea>
			<ofd:PhysicalBox>0 0 210 297</ofd:PhysicalBox>
		</ofd:PageArea>
		<ofd:PublicRes>PublicRes_0.xml</ofd:PublicRes>
		<ofd:DocumentRes>DocumentRes_0.xml</ofd:DocumentRes>
		<ofd:TemplatePage ID="4" BaseLoc="Temps/Temp_0/Content.xml"/>
	</ofd:CommonData>
	<ofd:Pages>
		<ofd:Page ID="6" BaseLoc="Pages/Page_0/Content.xml"/>
	</ofd:Pages>
	<ofd:Attachments>Attachs/Attachments.xml</ofd:Attachments>
	<ofd:CustomTags>Tags/CustomTags.xml</ofd:CustomTags>
</ofd:Document>
`

func TestDocumentXml(t *testing.T) {
	var doc Document
	if err := xml.Unmarshal([]byte(document_content), &doc); err != nil {
		t.Logf("%v", err)
	} else {
		t.Logf("%v", doc)
	}
}
func TestDocument(t *testing.T) {
	pwd, _ := os.Getwd()
	file := filepath.Join(pwd, "samples", "DZHD_1605281110201000001_202205250016050102202100000069141950_20220525_000413.ofd")
	ofdReader, err := NewOFDReader(file)
	if err != nil {
		t.Logf("%s", err)
	}
	defer ofdReader.Close()

	ofd, err := ofdReader.OFD()
	if err != nil {
		t.Logf("%v", err)
	} else {
		t.Logf("%v", ofd)
	}

	index_0 := ofd.DocBody[0].DocInfo.DocID.Text
	doc, err := ofdReader.GetDocumentById(index_0)
	if err != nil {
		t.Logf("%v", err)
	} else {
		t.Logf("%v", doc)
	}
	commonData, err := doc.GetCommonData()
	if err != nil {
		t.Logf("%v", err)
	} else {
		t.Logf("%v", commonData)
	}

	pages, err := doc.GetPages()
	if err != nil {
		t.Logf("%v", err)
	} else {
		t.Logf("%v", pages)
	}

	attachments, err := doc.GetAttachments()
	if err != nil {
		t.Logf("%v", err)
	} else {
		t.Logf("%v", attachments)
	}

	ct, err := doc.GetCustomTags()
	if err != nil {
		t.Logf("%v", err)
	} else {
		t.Logf("%v", ct)
	}
}
