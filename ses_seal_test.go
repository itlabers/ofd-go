package ofd

import (
	"archive/zip"
	"bytes"
	_ "encoding/hex"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_New(t *testing.T) {
	pwd, _ := os.Getwd()
	testcases := []struct {
		name   string // 来源
		path   string // 文件路径
		wanted bool   // 用bool方便判断是否返回error，如果类型改为error反而不好判断
	}{
		//  Add test cases.
		{"光大", filepath.Join(pwd, "samples", "2020062290131000005100000000013540122236009.ofd"), true},
		{"工行", filepath.Join(pwd, "samples", "0200216819200056225_001_21253000001_20220101_acc.ofd"), true},
		{"工行1", filepath.Join(pwd, "samples", "gonghang_1.ofd"), true},

		{"工行2", filepath.Join(pwd, "samples", "0402021509300142766_001_22206000039_20220725_acc.ofd"), true},
		{"中电财", filepath.Join(pwd, "samples", "DZHD_1605281110201000001_202205250016050102202100000069141950_20220525_000413.ofd"), true},
		{"浙商银行", filepath.Join(pwd, "samples", "bkrs_issuer_20220309_C1030311000455.ofd"), true},
	}
	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			file := test.path
			r, err := zip.OpenReader(file)
			if err != nil {
				t.Logf("%v", err.Error())
			}
			defer r.Close()
			for _, f := range r.File {
				t.Logf("文件名: %s\n", f.Name)
				if strings.HasSuffix(f.Name, ".dat") {
					rc, err := f.Open()
					if err != nil {
						t.Logf("%v", err.Error())
					}
					defer rc.Close()
					var buf bytes.Buffer
					_, err = io.CopyN(&buf, rc, int64(f.UncompressedSize64))
					if err != nil {
						t.Logf("%v", err.Error())
					}
					content := buf.Bytes()

					if ses, err := New_SES_Signature(content); err != nil {
						t.Logf("%v", err.Error())
					} else {
						tbs_sign, err := ses.Get_TBS_Sign()
						if err != nil {
							t.Logf("%v", err.Error())
						} else {
							seal, err := tbs_sign.Get_Seal()
							if err != nil {
								t.Logf("%v", err.Error())
							} else {
								t.Logf("%v", seal)
							}
						}
					}
				}

			}
		})
	}
} 
