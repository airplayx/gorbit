package sms

import (
	"encoding/json"
	"github.com/airplayx/gorbit/submail/lib"
)

type multisend struct {
	appid    string
	appkey   string
	signType string
	content  string
	multi    []map[string]interface{}
	tag      string
}

const multisendURL = lib.Server + "/message/multisend"

func CreateMultiSend(config map[string]string) *multisend {
	return &multisend{config["appid"], config["appkey"], config["signType"], "", nil, ""}
}

func (this *multisend) SetContent(content string) {
	this.content = content
}

func (this *multisend) AddMulti(multi map[string]interface{}) {
	this.multi = append(this.multi, multi)
}

func (this *multisend) SetTag(tag string) {
	this.tag = tag
}

func (this *multisend) MultiSend() string {
	config := make(map[string]string)
	config["appid"] = this.appid
	config["appkey"] = this.appkey
	config["signType"] = this.signType

	request := make(map[string]string)
	request["appid"] = this.appid
	request["content"] = this.content
	if this.signType != "normal" {
		request["sign_type"] = this.signType
		request["timestamp"] = lib.GetTimestamp()
		request["sign_version"] = "2"
	}
	if this.tag != "" {
		request["tag"] = this.tag
	}
	request["signature"] = lib.CreateSignature(request, config)
	//v2 数字签名 multi 不参与计算

	data, err := json.Marshal(this.multi)
	if err == nil {
		request["multi"] = string(data)
	}

	return lib.Post(multisendURL, request)
}
