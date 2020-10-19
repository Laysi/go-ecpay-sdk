package ecpayGin

import (
	"bytes"
	"fmt"
	"github.com/Laysi/go-ecpay-sdk"
	"io/ioutil"
)
import "github.com/gin-gonic/gin"

// ResponseBodyDateTimePatchHelper is a help handler to help binding the model that contain ecpay time format with the issue: https://github.com/gin-gonic/gin/issues/2510
func ResponseBodyDateTimePatchHelper(c *gin.Context) error {
	body, err := c.GetRawData()
	if err != nil {
		return err
	}

	data := string(body)
	data = ecpay.QuoteDatetime(data)
	fmt.Println(data)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(data)))
	return nil
}
