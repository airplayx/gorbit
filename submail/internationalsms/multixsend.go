package internationalsms

import (
	"encoding/json"
	"github.com/airplayx/gorbit/submail/lib"
)

type multixsend struct {
	appid    string
	appkey   string
	signType string
	project  string
	multi    []map[string]interface{}
}

type multi struct {
	to   string
	vars map[string]string
}

const multixsendURL = lib.Server + "/internationalsms/multixsend"

func CreateMulti() *multi {
	return &multi{"", make(map[string]string)}
}

func (this *multi) SetTo(to string) {
	this.to = to
}

func (this *multi) AddVar(key string, val string) {
	this.vars[key] = val
}

func (this *multi) Get() map[string]interface{} {
	item := make(map[string]interface{})
	item["to"] = this.to
	item["vars"] = this.vars
	return item
}

func CreateMultiXsend(config map[string]string) *multixsend {
	return &multixsend{config["appid"], config["appkey"], config["signType"], "", nil}
}

func (this *multixsend) SetProject(project string) {
	this.project = project
}

func (this *multixsend) AddMulti(multi map[string]interface{}) {
	this.multi = append(this.multi, multi)
}

func (this *multixsend) MultiXsend() string {
	config := make(map[string]string)
	config["appid"] = this.appid
	config["appkey"] = this.appkey
	config["signType"] = this.signType

	request := make(map[string]string)
	request["appid"] = this.appid
	request["project"] = this.project
	if this.signType != "normal" {
		request["sign_type"] = this.signType
		request["timestamp"] = lib.GetTimestamp()
		request["sign_version"] = "2"
	}
	request["signature"] = lib.CreateSignature(request, config)
	//v2 数字签名 multi 不参与计算

	data, err := json.Marshal(this.multi)
	if err == nil {
		request["multi"] = string(data)
	}

	return lib.MultipartPost(multixsendURL, request)
}
