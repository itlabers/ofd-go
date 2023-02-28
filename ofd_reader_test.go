package ofd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewOFDReader(t *testing.T) {
	pwd, _ := os.Getwd()
    // 替换成目标文件
	file := filepath.Join(pwd, "sample.ofd")

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

	doc_0 := ofd.DocBody[0].DocInfo.DocID.Text
	doc, err := ofdReader.GetDocumentById(doc_0)
	if err != nil {
		t.Logf("%v", err)
	} else {
		t.Logf("%v", doc)
	}

	if attachments, err := doc.GetAttachments(); err != nil {
		t.Logf("%v", err)
	} else {
		t.Logf("%v", attachments)
	}

	signs, err := ofdReader.GetSignaturesById(doc_0)
	if err != nil {
		t.Logf("%v", err)
	} else {
		t.Logf("%v", signs)
	}

	sign_1 := signs.Signature[0].ID
	sign, err := signs.GetSignatureById(sign_1)
	if err != nil {
		t.Logf("%v", err)
	} else {
		t.Logf("%v", sign)
	}
	if result, err := sign.VerifyDigest(); err != nil {
		t.Logf("%v", err)
	} else {
		t.Logf("%v", result)
	}
	if result, err := sign.Verify(); err != nil {
		t.Logf("%v", err)
	} else {
		t.Logf("%v", result)
	}

}

func TestBatch(t *testing.T) {
	pwd, _ := os.Getwd()
	testcases := []struct {
		name   string
		path   string
		wanted bool
	}{
		//  Add test cases.
		{"光大", filepath.Join(pwd, "samples", "2020062290131000005100000000013540122236009.ofd"), true},
		{"工行", filepath.Join(pwd, "samples", "0200216819200056225_001_21253000001_20220101_acc.ofd"), true},
		{"中电财", filepath.Join(pwd, "samples", "DZHD_1605281110201000001_202205250016050102202100000069141950_20220525_000413.ofd"), true},
		{"浙商银行", filepath.Join(pwd, "samples", "bkrs_issuer_20220309_C1030311000455.ofd"), true},
		{"中国铁路-退票", filepath.Join(pwd, "samples", "退票样例.ofd"), true},
		{"中国铁路-售票", filepath.Join(pwd, "samples", "售票换开样例.ofd"), true},
	}
	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			ofdReader, err := NewOFDReader(test.path, WithValidator(&CommonValidator{}))
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

			for _, doc := range ofd.DocBody {
				doc_0 := doc.DocInfo.DocID.Text
				doc, err := ofdReader.GetDocumentById(doc_0)
				if err != nil {
					t.Logf("%v", err)
				} else {
					t.Logf("%v", doc)
				}

				signs, err := ofdReader.GetSignaturesById(doc_0)
				if err != nil {
					t.Logf("%v", err)
				}

				sign_1 := signs.Signature[0].ID
				sign, _ := signs.GetSignatureById(sign_1)
				if err != nil {
					t.Logf("%v", err)
				}
				if result, err := sign.VerifyDigest(); err != nil {
					t.Logf("%v", err)
				} else {
					t.Logf("%v", result)
				}
				if result, err := sign.Verify(); err != nil {
					t.Logf("%v", err)
				} else {
					t.Logf("%v", result)
				}
			}

		})
	}
}
