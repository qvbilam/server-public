package sms

import (
	"encoding/json"
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Client struct {
	AccessKeyId     string
	AccessKeySecret string
	SignName        string
	TemplateCode    string
	*dysmsapi20170525.Client
}

var successCode = "OK"
var url = "dysmsapi.aliyuncs.com"

func (c *Client) Create() (*dysmsapi20170525.Client, error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: &c.AccessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: &c.AccessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String(url)
	res := &dysmsapi20170525.Client{}
	res, err := dysmsapi20170525.NewClient(config)
	return res, err
}

func (c *Client) SendCode(mobile, code string) (*string, error) {
	type param struct {
		Code string `json:"code"`
	}
	p := param{Code: code}
	templateParam, _ := json.Marshal(p)

	// 工程代码泄露可能会导致AccessKey泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	client, err := c.Create()
	if err != nil {
		return nil, err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  &mobile,
		SignName:      &c.SignName,
		TemplateCode:  tea.String(c.TemplateCode),
		TemplateParam: tea.String(string(templateParam)),
	}
	res, tryErr := func() (res *dysmsapi20170525.SendSmsResponse, e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		res, err := client.SendSmsWithOptions(sendSmsRequest, &util.RuntimeOptions{})
		if err != nil {
			return res, err
		}

		return res, nil
	}()

	if tryErr != nil {
		var teaError = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			teaError = _t
		} else {
			teaError.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, err := util.AssertAsString(teaError.Message)
		if err != nil {
			return nil, err
		}
	}

	if *res.Body.Code != successCode {
		return nil, errors.New(*res.Body.Message)
	}

	return res.Body.RequestId, nil
}
