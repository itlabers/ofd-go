## ofd 文档

OFD（Open Fixed-layout Documents的简称，意为开放版式文件）版式文档是版面呈现效果高度精确固定的电子文件，其呈现与设备无关。与pdf文件相仿，具有格式独立、版面固定、固化呈现等特点。OFD也逐渐开始在电子发票、电子公文、电子证照等等的领域中应用。

## ofd 特点

OFD标准有一系列技术优势。

1. 体积精简，格式开放，利于理解，长期可读可用
2. 根据我国各领域特色需要进行特性扩展，更深入地贴合了应用需求
3. 标准可支持国产密码算法，是文档安全性的有力保证，也是文件具有法律效力的基本条件
4. 标准是自主可控的，国家再有需要对OFD做上面提到的扩展时，特别是在我国党政军严肃类文档应用领域，可以不受控于外部的厂商，我们有自主的标准话语权 

## 自定义签名验证器

通过实现 Validator 接口 ,初始化的时候加载验证器WithValidator

```
ofdReader, err :=NewOFDReader(testpath, WithValidator(XXValidator{}))
if err != nil {
	t.Logf("%s", err)
}
defer ofdReader.Close()
```

// 文本摘要
Digest([]byte) []byte

// 签名验证
Verify([]byte, []byte, []byte) (bool, error)

**完整例子**
以github.com/emmansun/gmsm  国密算法库为例

```
package ofd

import (
	"crypto/ecdsa"
	std "encoding/asn1"
	"math/big"

	"github.com/emmansun/gmsm/sm2"
	"github.com/emmansun/gmsm/sm3"
	"github.com/emmansun/gmsm/smx509"
)

type OtherValidator struct {
}


func (gm *OtherValidator) Digest(msg []byte) []byte {
	h := sm3.New()
	h.Write(msg)
	dataDash := h.Sum(nil)
	return dataDash

}
func (gm *OtherValidator) Verify(cert []byte, msg []byte, signature []byte) (bool, error) {
	certificate, err := smx509.ParseCertificate(cert)
	if err != nil {
		return false, err
	}
	pk := certificate.PublicKey.(*ecdsa.PublicKey)
	if len(signature) == 64 {
		r := new(big.Int).SetBytes(signature[0:32])
		s := new(big.Int).SetBytes(signature[32:64])

		result := sm2.VerifyWithSM2(pk, nil, msg, r, s)
		return result, nil
	} else {
		type Sign struct {
			R *big.Int
			S *big.Int
		}
		var sign Sign
		if _, err := std.Unmarshal(signature, &sign); err != nil {
			return false, err
		} else {
			ff, _ := new(big.Int).SetString(MAX, 16)
			if sign.R.Sign() == -1 {
				sign.R.And(sign.R, ff)
			}
			if sign.S.Sign() == -1 {
				sign.S.And(sign.S, ff)
			}
			result := sm2.VerifyWithSM2(pk, nil, msg, sign.R, sign.S)
			return result, nil
		}
	}
}

```

## 小通知
  ofd-go项目发起有段时间,由于工作中涉及ofd的相关开发,我们团队曾经使用ofd-fw这个java库来完成我们的开发任务，使用过程中发现不少问题,期间我们也fork该库,形成自己的分支,问题的根源在于尽管ofd标准是统一的，但是各家软件集成商对于标准的理解不太一样,特别是关键字段的赋值上差别很大，所以需要针对不同的ofd软件集成商做适配,才能实现通用，由于ofd-rw的解决方案足有优秀,我们决定开辟新天地,ofd-go正是在这个背景下产生的，依托于go语言优秀的跨平台能力,成熟的生态,可以高度定制,开发HTTP服务,SDK,动态链接库，都是不错的选择.

  目前的ofd-go每天的独立clone在50左右，反映出在市面上有一定的需求，但是除了项目初期有热心人士关注，提出issue,提交一些bug,后续的clone者反馈问题较少,作为发起人，我意识到ofd-go本身还不完善，为了持续推进开源项目，本着"我为人人，人人为我"的精神，现阶段希望征集大家对于ofd-go的需求提案,无论是提交bug、提供ofd测试样本或者是提供产品思路,都希望广泛交流，为此我建了讨论组，感兴趣的可以加群
  
  QQ:628529123
![QQ群](https://github.com/itlabers/ofd-go-reference/blob/master/images/qq_group.jpg)

  