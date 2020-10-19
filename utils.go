package ecpay

import "regexp"

var regexFormUrlEncodedDatetime = regexp.MustCompile(`=(\d{4}\/\d{2}\/\d{2} \d{2}:\d{2}:\d{2})`)

func QuoteDatetime(data string) string {
	return regexFormUrlEncodedDatetime.ReplaceAllString(data, `="$1"`)
}
