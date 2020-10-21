package ecpay

import "regexp"

var regexFormUrlEncodedDatetime = regexp.MustCompile(`=(\d{4}\/\d{2}\/\d{2} \d{2}:\d{2}:\d{2})`)

func QuoteDatetime(data string) string {
	return regexFormUrlEncodedDatetime.ReplaceAllString(data, `="$1"`)
}

func PtrNilString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func PtrNilInt(s int) *int {
	if s == 0 {
		return nil
	}
	return &s
}
