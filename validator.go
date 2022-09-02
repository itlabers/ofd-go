package ofd

import (
	std "encoding/asn1"
	"math/big"

	"github.com/itlabers/crypto/sm/sm2"
	"github.com/itlabers/crypto/sm/sm3"
	smx509 "github.com/itlabers/crypto/x509"
)

const (
	SM3_OID        = "1.2.156.10197.1.401"
	SM3WITHSM2_OID = "1.2.156.10197.1.501"
	MAX            = "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
)

type Category int

const (
	SIGN Category = iota
	SEAL
)

type Hash interface {
	Digest([]byte) []byte
}
type Validator interface {
	Hash

	Verify([]byte, []byte, []byte) (bool, error)
}

type CommonValidator struct {
}

func (common *CommonValidator) Digest(msg []byte) []byte {
	h := sm3.New()
	h.Write(msg)
	dataDash := h.Sum(nil)
	return dataDash

}
func (common *CommonValidator) Verify(cert []byte, msg []byte, signature []byte) (bool, error) {
	certificate, err := smx509.ParseCertificate(cert)
	if err != nil {
		return false, err
	}
	pk := certificate.PublicKey.(*sm2.PublicKey)
	hashed := sm3.New()
	if len(signature) == 64 {
		r := new(big.Int).SetBytes(signature[0:32])
		s := new(big.Int).SetBytes(signature[32:64])
		result := sm2.Verify(pk, "", msg, hashed, r, s)
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
			result := sm2.Verify(pk, "", msg, hashed, sign.R, sign.S)
			return result, nil
		}
	}
}
