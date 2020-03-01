// rsa签名（使用客户提供的证书所解析的密钥）
package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"

	// "github.com/golang/crypto/pkcs12"
	// "github.com/golang/crypto/pkcs12"
	"golang.org/x/crypto/pkcs12"
)

var signKey_private = []byte(`-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCbIYzaa88j6vGs
xHHFawn1JNie7bjs0CMUcEdxIwGKCPA99KagliNBCOOoQ4ntZpgb7SL+yYsQ6RKN
LrCWg4DbKYJbN0h+mIlcVSxYZoputVXzRGaqzq7VcanihHblq0wKFt54gwrWx4Cf
mUXm+Av9sEAvlNtFgTG2q33UibSTev9MDg3JFFtKTT0VrG07etBZXvHOh3OoXIX7
xzbsGHIlj2gWLnibsvTpfInBuYCkf2U8oPN8dhnG1SFyg0s5/1AdtHvK2AW+RKzy
mQKSIAUzltRvfT15RWAZCDFH9aiwlTNSeIv+VZf8yjwDI88iXAvK9926K799wrrN
iICjW/xZAgMBAAECggEAEkClUQHEUoaPPvOyKzdH6acXMcPosTxgFESNs1saEEKy
+acyf8yKg1lnZ+HidaNPnGGPb4kJKskxnJj0wBDi9WZx5wefKP4Lk2iWGiSefx4o
6cGPjuNYoNaJ1ZQeZTF9deZw8J8ChHjewmiHWACxcHzvYGnUTt4/mi2dVYeKTpvI
SiyqdUAS/6IfPQB3u9K0FrydHTEx34Zf0uek/OfxjWJPCYCEkL5hviX3jbSRx8Ze
xK+lYiCXE/x7JZB6UDMz32ksIkciGs1vxSQC/v3wgH783lk7lCvX1EdGEON9C6Tg
Uo8gEYgOXYWNzzc44qId/KIpJEhHkc79AJm3sq8w3QKBgQDBz8xF1TKBtiUeuxhM
qmi3iUTPCB9q+f8iWdgE8HVsxTbaS5tnuqwxQRi6fhAu4sf9ezEm5J8sRhYfP86F
Jx9fuWJ9ETCIZHdxyviqSBU/N6iAn36ykTsmnZylw3VqQB6MM4eUhfaIfxVI0/0T
q/eSuMLZs46crvbabkOefrzh+wKBgQDM6Gv9Vwfr2RuS0B0uzK6iedOxG2LrRpvT
VsnqfI1TncTviVHCuBfphUZOKbqIn66fZVP50XD3Po+x9CvMIG+oEDmTT3lUc4vg
M+JXMYsKBzxVbZaQkfCpOh0VY5T+5DLtY66llDAvvdnByo5pEQjr+3XtiQG8Ho5M
SJptmZ6euwKBgQC/X/ncJtX1gS8MNkCbS1WVjaT1LZcgYfUNLVwCHJJGskIksa2l
co2fISFS3TILqgFCigeR2QRXimlDuBSsJGJ1ejAHQRcERvgu0/ZQ5lacks//Gf7Y
QlH/InjQsb1aCjBfzm4dOV6jj6Erxa/LYe5X4br9jrDOhHHaZDC+63SFSwKBgFXS
bxB/uDCfET5vbhmg3u/sKK8AMnEqdK+SpSVSH3tLFCXjlNwCcPl8uZKZRahdgSz/
qqlygGb9y/+Trhhj6YAxs3uwiLfHYWl5Ma6lcH+1wh5htPM+PNcUIQRJXilX+MCW
xJpTM897K85QAU+eM+dqvdzMSm65OY5Xxl5rPhVrAoGAATM8OCyyyAhZIzQOuLiz
wCKYsUshKYV4OkyIj5zwP7gY1szlPSq0/X4KdhxqCQOYTILkjzyaU/fGULNOb9MH
4hQM9m0AnzU4v2lECgzXfor45B3dquVh0f+RqJbwYKCiS82GVQaZ71GOnNofxBIY
v+eJtVS6pH2ssGTQIC/D0ZQ=
-----END PRIVATE KEY-----`)

func main() {
	pfxFile, err := ioutil.ReadFile("./test/zhengshu/yqh.pfx")
	if err != nil {
		fmt.Println(err)
	}
	// blocks, err := pkcs12.ToPEM(pfxFile, "111111")
	// if err != nil {
	// 	panic(err)
	// }
	// var pemData []byte
	// for _, b := range blocks {
	// 	pemData = append(pemData, pem.EncodeToMemory(b)...)
	// }
	// fmt.Println("qqqqqq", string(pemData))

	private, _, err := pkcs12.Decode(pfxFile, "111111")
	if err != nil {
		fmt.Println("2222222")
	}
	prive := private.(*rsa.PrivateKey)
	var hash crypto.Hash = crypto.MD5
	h := hash.New()
	h.Write([]byte("&111&"))
	digest := h.Sum(nil)
	data, err := rsa.SignPKCS1v15(nil, prive, hash, digest)
	if err != nil {
		fmt.Errorf("rsaSign SignPKCS1v15 error")
	}
	fmt.Println("eeeee", base64.StdEncoding.EncodeToString(data))

	rsa_key, _ := SignPKCS1v15([]byte("&111&"), signKey_private, hash)
	fmt.Println("qqqqqq", rsa_key)

	rsa_key1, err := SignPKCS1v15xxx([]byte("&111&"), signKey_private, hash)
	if err != nil {
		fmt.Println("err===", err.Error())
	}
	fmt.Println("zzzzzzzz", rsa_key1)
}

func SignPKCS1v15(origData, privateKey []byte, hash crypto.Hash) (string, error) {
	h := hash.New()
	h.Write(origData)
	digest := h.Sum(nil)

	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("privateKey key error")
	}
	var pri *rsa.PrivateKey
	var err error
	if hash == crypto.MD5 {
		fmt.Println("==========")
		pubInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return "", err
		}
		pri = pubInterface.(*rsa.PrivateKey)
	} else {
		fmt.Println("************")

		pri, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return "", err
		}
	}
	data, err := rsa.SignPKCS1v15(nil, pri, hash, digest)
	if err != nil {
		fmt.Errorf("rsaSign SignPKCS1v15 error")
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func SignPKCS1v15xxx(origData, privateKey []byte, hash crypto.Hash) (string, error) {
	h := hash.New()
	h.Write(origData)
	digest := h.Sum(nil)

	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("privateKey key error")
	}
	var pri *rsa.PrivateKey
	var err error

	fmt.Println("************")

	pri, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	data, err := rsa.SignPKCS1v15(nil, pri, hash, digest)
	if err != nil {
		fmt.Errorf("rsaSign SignPKCS1v15 error")
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}
