package ecpay

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/Laysi/go-ecpay-sdk/base"
	"html/template"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

const AioCheckOutPath = "/Cashier/AioCheckOut/V5"

type ECPayMode int

const (
	PRODUCTION_MODE ECPayMode = iota
	STAGE_MODE
)

func (e ECPayMode) Index() int {
	return int(e)
}

type ECPayClient struct {
	MerchantID      string
	AioCheckOutPath string
	ReturnURL       string
	HashKey         string
	HashIV          string
	mode            ECPayMode
	//ctx             context.Context
	apiClient *base.APIClient
}

func NewECPayClient(merchantID string, returnURL string, hashIV string, hashKey string) *ECPayClient {

	return &ECPayClient{
		MerchantID:      merchantID,
		AioCheckOutPath: AioCheckOutPath,
		ReturnURL:       returnURL,
		HashKey:         hashIV,
		HashIV:          hashKey,
		mode:            PRODUCTION_MODE,
		//ctx:             context.WithValue(context.Background(), base.ContextServerIndex, PRODUCTION_MODE),
		apiClient: base.NewAPIClient(base.NewConfiguration()),
	}
}

func NewStageECPayClient() *ECPayClient {
	return &ECPayClient{
		MerchantID:      "2000132",
		AioCheckOutPath: AioCheckOutPath,
		ReturnURL:       "https://example.com/return",
		HashKey:         "5294y06JbISpM5x9",
		HashIV:          "v77hoKGq4kWxNNIS",
		mode:            STAGE_MODE,
		//ctx:             context.WithValue(context.Background(), base.ContextServerIndex, STAGE_MODE),
		apiClient: base.NewAPIClient(base.NewConfiguration()),
	}
}

func (e ECPayClient) GetCurrentServer() string {
	serverUrl, err := e.apiClient.GetConfig().ServerURL(e.mode.Index(), map[string]string{})
	if err != nil {
		panic(err)
	}
	return serverUrl
}

func (e ECPayClient) DefaultContext() context.Context {
	return context.WithValue(context.Background(), base.ContextServerIndex, e.mode)
}

func (e ECPayClient) WithContext(c context.Context) context.Context {
	return context.WithValue(c, base.ContextServerIndex, e.mode)
}

func FormUrlEncode(s string) string {
	s = url.QueryEscape(s)
	s = strings.ReplaceAll(s, "%2d", "-")
	s = strings.ReplaceAll(s, "%5f", "_")
	s = strings.ReplaceAll(s, "%2e", ".")
	s = strings.ReplaceAll(s, "%21", "!")
	s = strings.ReplaceAll(s, "%2a", "*")
	s = strings.ReplaceAll(s, "%28", "(")
	s = strings.ReplaceAll(s, "%29", ")")
	return s
}

func (e ECPayClient) GenerateCheckMacValue(params map[string]string) string {
	delete(params, "CheckMacValue")
	delete(params, "HashKey")
	delete(params, "HashIV")
	encodedParams := NewECPayValuesFromMap(params).Encode()
	encodedParams = fmt.Sprintf("HashKey=%s&%s&HashIV=%s", e.HashKey, encodedParams, e.HashIV)
	encodedParams = FormUrlEncode(encodedParams)
	encodedParams = strings.ToLower(encodedParams)
	//fmt.Println(encodedParams)
	sum := sha256.Sum256([]byte(encodedParams))
	checkMac := strings.ToUpper(hex.EncodeToString(sum[:]))
	return checkMac
}

var OrderTemplateText = `<form id="order_form" action="{{.Action}}" method="post">
 {{range $key,$element := .Values -}} 
	<input type="hidden" name="{{$key}}" id="{{$key}}" value="{{$element}}" /> 
{{end -}}
</form>
<script>document.querySelector("#order_form").submit();</script>`

type OrderTmplArgs struct {
	Values map[string]string
	Action string
}

var OrderTmpl = template.Must(template.New("AutoPostOrder").Parse(OrderTemplateText))

func (e ECPayClient) GenerateAutoSubmitHtmlForm(params map[string]string, targetUrl string) string {

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

func StructToParamsMap(data interface{}) map[string]string {
	params := map[string]string{}
	iVal := reflect.ValueOf(data)
	iTyp := reflect.TypeOf(data)
	if iVal.Kind() == reflect.Ptr {
		iVal = iVal.Elem()
	}
	if iTyp.Kind() == reflect.Ptr {
		iTyp = iTyp.Elem()
	}

	//stringerInterface := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	marshalerInterface := reflect.TypeOf((*json.Marshaler)(nil)).Elem()
	//typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		ft := iTyp.Field(i)

		if f.Kind() == reflect.Ptr && f.IsNil() {
			continue
		}
		if f.Kind() == reflect.Ptr {
			f = f.Elem()

		}

		if ft.Anonymous {
			nestedParams := StructToParamsMap(f.Interface())
			for key, val := range nestedParams {
				params[key] = val
			}
			continue
		}

		var v string
		//f.Anonymous
		switch realVal := f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(realVal)
		case string:
			v = realVal
		//case time.Time:
		//	v = realVal.Format(time.RFC3339)
		//case base.ECPayDateTime:
		//	v = realVal.String()
		default:
			switch {
			case f.Type().Implements(marshalerInterface):
				data, err := f.Interface().(json.Marshaler).MarshalJSON()
				if err != nil {
					panic(err)
				}
				unquoteData, err := strconv.Unquote(string(data))
				if err != nil {
					v = string(data)
				}
				v = unquoteData
			//case f.Type().Implements(stringerInterface):
			//	v = f.Interface().(fmt.Stringer).String()
			default:
				switch f.Kind() {
				case reflect.String:
					v = f.String()
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					v = strconv.FormatInt(f.Int(), 10)
				default:
					panic(fmt.Sprintf("Unknown type %T during transfer struct to map!\n", f.Interface()))

				}
			}
		}
		params[ft.Name] = v
	}
	return params
}
