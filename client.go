package ecpay

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
)

type ECPayClient struct {
	HashKey string
	HashIV  string
}

func NewECPayClient(hashIV string, hashKey string) *ECPayClient {
	return &ECPayClient{HashKey: hashIV, HashIV: hashKey}
}

func NewStageECPayClient() *ECPayClient {
	return &ECPayClient{HashKey: "5294y06JbISpM5x9", HashIV: "v77hoKGq4kWxNNIS"}
}

func (c ECPayClient) GenerateCheckMacValue(params url.Values) string {
	encodedParams := ECPayValues{params}.Encode()
	encodedParams = fmt.Sprintf("HashKey=%s&%s&HashIV=%s", c.HashKey, encodedParams, c.HashIV)
	encodedParams = url.QueryEscape(encodedParams)
	encodedParams = strings.ToLower(encodedParams)
	sum := sha256.Sum256([]byte(encodedParams))
	checkMac := strings.ToUpper(hex.EncodeToString(sum[:]))
	return checkMac
}

func GenerateAutoSubmitHtmlForm(params url.Values) {

}
