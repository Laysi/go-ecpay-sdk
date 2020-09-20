package ecpay

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"html/template"
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

var OrderTemplateText = `<form id="order_form" action="{{.Action}}" method="post">
 {{range $key,$element := .Values -}} 
	<input type="hidden" name="{{$key}}" id="{{$key}}" value="{{index $element 0}}" /> 
{{end -}}
</form>
<script>$("#order_form").submit();</script>`

type OrderTmplArgs struct {
	Values url.Values
	Action string
}

var OrderTmpl = template.Must(template.New("AutoPostOrder").Parse(OrderTemplateText))

func (c ECPayClient) GenerateAutoSubmitHtmlForm(params url.Values, targetUrl string) string {

	var result bytes.Buffer
	err := OrderTmpl.Execute(&result, OrderTmplArgs{
		Values: params,
		Action: targetUrl,
	})
	if err != nil {
		panic(err)
	}
	return result.String()
}
