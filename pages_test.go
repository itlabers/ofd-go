package ofd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPages(t *testing.T) {
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
	allPages, err := pages.GetPages()
	if err != nil {
		t.Logf("%v", err)
	} else {
		t.Logf("%v", allPages)
	}

	page_0 := pages.Page[0].ID
	page, err := pages.GetPageById(page_0)
	if err != nil {
		t.Logf("%v", err)
	} else {
		t.Logf("%v", page)
	}
}
