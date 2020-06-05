package sms

import (
	"github.com/airplayx/gorbit/submail/lib"
)

type send struct {
	appid    string
	appkey   string
	signType string
	to       string
	content  string
	tag      string
}

const sendURL = lib.Server + "/message/send"

func CreateSend(config map[string]string) *send {
	return &send{config["appid"], config["appkey"], config["signType"], "", "", ""}
}

func (this *send) SetTo(to string) {
	this.to = to
}

func (this *send) SetContent(content string) {
	this.content = content
}

func (this *send) SetTag(tag string) {
	this.tag = tag
}

func (this *send) Send() string {
	config := make(map[string]string)
	config["appid"] = this.appid
	config["appkey"] = this.appkey
	config["signType"] = this.signType

	request := make(map[string]string)
	request["appid"] = this.appid
	request["to"] = this.to
	if this.signType != "normal" {
		request["sign_type"] = this.signType
		request["timestamp"] = lib.GetTimestamp()
		request["sign_version"] = "2"
	}
	if this.tag != "" {
		request["tag"] = this.tag
	}
	request["signature"] = lib.CreateSignature(request, config)
	//v2 数字签名 content 不参与计算
	request["content"] = this.content
	return lib.Post(sendURL, request)
}
