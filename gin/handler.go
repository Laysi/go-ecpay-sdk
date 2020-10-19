package ecpayGin

import (
	"bytes"
	"github.com/Laysi/go-ecpay-sdk"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func ECPayCheckMacValueHandler(client *ecpay.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := c.GetRawData()
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		err = c.Request.ParseForm()
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		params := ecpay.ECPayValues{c.Request.PostForm}.ToMap()
		c.Request.Form = nil
		c.Request.PostForm = nil

		senderMac := params["CheckMacValue"]
		delete(params, "CheckMacValue")
		mac := client.GenerateCheckMacValue(params)
		if mac != senderMac {
			c.String(http.StatusBadRequest, "0|Error")
			c.Abort()
		}
	}

}
