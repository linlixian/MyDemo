//BASE64：编码方式（8位字节代码），二进制（[]byte）与字符串(string)相互转换
//（1字节（byte）=8位，A=65=01000001，a=97=01100001  字符 0=48  9=57）
// UTF-8编码有一个额外的好处，就是ASCII编码实际上可以被看成是UTF-8编码的一部分，
// s所以，大量只支持ASCII编码的历史遗留软件可以在UTF-8编码下继续工作。
package main

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
)

//"zcm_crm_zj/models/crm"

// -------------------
//      Interface
// -------------------
const (
	key = `h*.d;cy7x_12dkx?#j39fdl!`
)

func main() {
	data := []byte("6uEjUa3s3ZiDHU59d474CI9P6YfVWjcYGBJZRARymyTv2kEtTx6ofTwhO0yrIBXlS2DHJ+GCzRSWNB+s6rNF/vKHPQ/c5Dbw47hQhNGIab+6IrhMsl/kPQl/4DP2TY0+2NtlpysdEWAnycDXFOXJL+WLcSlJWHe2iHu7CwHqDniuC4zV587zwZumXkRKNSEZCyjO5ZtYjIE=")
	fmt.Println(data)
	// panic(errors.New("User stop run"))
	respData := DesBase64Decrypt(data)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	fmt.Println(string(respData))

	s := []byte("6")
	fmt.Println(s)

}

func DesBase64Decrypt(crypted []byte) []byte {
	result, _ := base64.StdEncoding.DecodeString(string(crypted))
	remain := len(result) % 8
	if remain > 0 {
		mod := 8 - remain
		for i := 0; i < mod; i++ {
			result = append(result, 0)
		}
	}
	origData, err := TripleDesDecrypt(result, []byte(key))
	if err != nil {
		panic(err)
	}
	return origData
}

// 3DES解密
func TripleDesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
