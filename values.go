package ecpay

import (
	"net/url"
	"sort"
	"strings"
)

type LowerStringSlice []string

func (p LowerStringSlice) Len() int           { return len(p) }
func (p LowerStringSlice) Less(i, j int) bool { return strings.ToLower(p[i]) < strings.ToLower(p[j]) }
func (p LowerStringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type ECPayValues struct {
	url.Values
}

func NewECPayValuesFromMap(values map[string]string) *ECPayValues {
	v := ECPayValues{Values: url.Values{}}
	for key, value := range values {
		v.Set(key, value)
	}
	return &v
}

func (v ECPayValues) Encode() string {
	if v.Values == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v.Values))
	for k := range v.Values {
		keys = append(keys, k)
	}
	sort.Sort(LowerStringSlice(keys))
	for _, k := range keys {
		vs := v.Values[k]
		//keyEscaped := url.QueryEscape(k)
		keyEscaped := k
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}
	return buf.String()
}
