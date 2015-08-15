package main

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
)

var (
	ErrIllegalParameter = errors.New("illegal parameter(s)")
	ErrDataToLarge      = errors.New("data is too large (len > 128) ")
	ErrDataLen          = errors.New("data length error")
	ErrDataBroken       = errors.New("data broken, first byte is not zero")
	ErrKeyPairDismatch  = errors.New("data is not encrypted by the private key")
)

const (
	//MaxDataLen max data single decrypt data length.
	maxDataLen = 128
)

// leftPad returns a new slice of length size. The contents of input are right
// aligned in the new slice.
// copy from crypto/rsa/rsa.go.
func leftPad(input []byte, size int) (out []byte) {
	n := len(input)
	if n > size {
		n = size
	}
	out = make([]byte, size)
	copy(out[len(out)-n:], input)
	return
}

func rsaPublicKeyDecryptHelper(pemKey string, data []byte) ([]byte, error) {

	if len(data) > maxDataLen {
		return nil, ErrDataToLarge
	}

	block, _ := pem.Decode([]byte(pemKey))

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pubKey := pub.(*rsa.PublicKey)

	k := (pubKey.N.BitLen() + 7) / 8

	if k != len(data) {
		return nil, ErrDataLen
	}
	m := new(big.Int).SetBytes(data)

	if m.Cmp(pubKey.N) > 0 {
		return nil, ErrDataToLarge
	}

	m.Exp(m, big.NewInt(int64(pubKey.E)), pubKey.N)

	d := leftPad(m.Bytes(), k)

	if d[0] != 0 {
		return nil, ErrDataBroken
	}

	if d[1] != 0 && d[1] != 1 {
		return nil, ErrKeyPairDismatch
	}

	var i = 2
	for ; i < len(d); i++ {
		if d[i] == 0 {
			break
		}
	}
	i++
	if i == len(d) {
		return nil, nil
	}
	return d[i:], nil

}

//RSAPublicKeyDecryptBase64 implement decrypt with base64 first and then RSA Public Key.
func RSAPublicKeyDecryptBase64(pem string, data []byte) ([]byte, error) {

	if pem == "" || data == nil {
		return nil, ErrIllegalParameter
	}

	dst := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	_, err := base64.StdEncoding.Decode(dst, data)
	if err != nil {
		return nil, err
	}

	return RSAPublicKeyDecrypt(pem, dst)
}

//RSAPublicKeyDecrypt implement decrypt with RSA Public Key.
func RSAPublicKeyDecrypt(pem string, data []byte) ([]byte, error) {

	if pem == "" || data == nil {
		return nil, ErrIllegalParameter
	}

	var (
		err       error
		plainText []byte
	)

	buf := &bytes.Buffer{}

	for len(data) >= maxDataLen {

		dataStub := data[:maxDataLen]
		data = data[maxDataLen:]

		plainText, err = rsaPublicKeyDecryptHelper(pem, dataStub)
		if err != nil {
			break
		}
		buf.Write(plainText)
	}

	return buf.Bytes(), nil

}

func main() {

	publicKey := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC17fxyt9Bqbpu/vPsYK8PUCYbX
vqMlyGEx0MYsT2mFtz61+kBe5lKpCKcOA4MsiCn4R+kiOAcLQqHfOU/t94AtMWK9
RNSFOt+s/bTB17N5Plebnj5jhLnzNbmGdK1GLNYvyrR62OUZppg+26no7qjtmK9D
46BPph9r06KXSotM9QIDAQAB
-----END PUBLIC KEY-----
`
	data := []byte(`
Q7hkKkGCjTdcKck1O9M8ieVsYdFmfDAqz/Aj0nqAGUSCD
6L1eSYic/DCk+uep4s7MxqkFgqjNQ2W6CZ/gzVgz93+qZ
ElYcGKXiJwB4Gdb/Lxr/CiIXQTmQ1q35bMaNuIZu8E8IX
AjGMQbY+ML/P3CDSuJPN6yOk4XRmzvMrNFmI=`)

	buf, _ := RSAPublicKeyDecryptBase64(publicKey, data)
	fmt.Println(string(buf))
	base64.StdEncoding.DecodeString()
}
