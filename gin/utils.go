package gin

import (
	"bytes"
	"github.com/Laysi/go-ecpay-sdk/base"
	"github.com/gin-gonic/gin/binding"
	"io/ioutil"
	"net/url"
	"strconv"
)
import "github.com/gin-gonic/gin"

func BindingPeriodReturnData(data *base.PeriodReturnData, c *gin.Context) error {
	return QuotedAndBindingData(data, []string{"ProcessDate", "TradeDate"}, c)
}

func BindingReturnData(data *base.ReturnData, c *gin.Context) error {
	return QuotedAndBindingData(data, []string{"PaymentDate", "TradeDate"}, c)
}

func QuotedAndBindingData(data interface{}, quotedfields []string, c *gin.Context) error {
	body, err := c.GetRawData()
	if err != nil {
		return err
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	err = c.Request.ParseForm()
	if err != nil {
		return err
	}

	quotedValues(c.Request.PostForm, quotedfields)

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(c.Request.PostForm.Encode())))
	err = c.MustBindWith(data, binding.FormPost)
	if err != nil {
		return err
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return nil
}

func quotedValues(values url.Values, fields []string) {
	for _, val := range fields {
		values.Set(val, strconv.Quote(values.Get(val)))
	}
}
