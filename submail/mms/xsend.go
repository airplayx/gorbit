package mms

import (
	"encoding/json"
	"github.com/airplayx/gorbit/submail/lib"
)

type xsend struct {
	appid    string
	appkey   string
	signType string
	to       string
	project  string
	vars     map[string]string
	tag      string
}

const xsendURL = lib.Server + "/mms/xsend"

func CreateXsend(config map[string]string) *xsend {
	return &xsend{config["appid"], config["appkey"], config["signType"], "", "", make(map[string]string), ""}
}

func (this *xsend) SetTo(to string) {
	this.to = to
}

func (this *xsend) SetProject(project string) {
	this.project = project
}

func (this *xsend) AddVar(key string, val string) {
	this.vars[key] = val
}

func (this *xsend) SetTag(tag string) {
	this.tag = tag
}

func (this *xsend) Xsend() string {
	config := make(map[string]string)
	config["appid"] = this.appid
	config["appkey"] = this.appkey
	config["signType"] = this.signType

	request := make(map[string]string)
	request["appid"] = this.appid
	request["to"] = this.to
	request["project"] = this.project
	if this.signType != "normal" {
		request["sign_type"] = this.signType
		request["timestamp"] = lib.GetTimestamp()
		request["sign_version"] = "2"
	}
	if this.tag != "" {
		request["tag"] = this.tag
	}
	request["signature"] = lib.CreateSignature(request, config)
	//v2 数字签名 vars 不参与计算
	if len(this.vars) > 0 {
		data, err := json.Marshal(this.vars)
		if err == nil {
			request["vars"] = string(data)
		}
	}
	return lib.Post(xsendURL, request)
}
