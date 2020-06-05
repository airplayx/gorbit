package voice

import (
	"encoding/json"
	"github.com/airplayx/gorbit/submail/lib"
)

type Multixsend struct {
	appID    string
	appKey   string
	signType string
	project  string
	multi    []map[string]interface{}
}

type Multi struct {
	to   string
	vars map[string]string
}

const multixsendURL = lib.Server + "/voice/multixsend"

func CreateMulti() *Multi {
	return &Multi{"", make(map[string]string)}
}

func (m *Multi) SetTo(to string) {
	m.to = to
}

func (m *Multi) AddVar(key string, val string) {
	m.vars[key] = val
}

func (m *Multi) Get() map[string]interface{} {
	item := make(map[string]interface{})
	item["to"] = m.to
	item["vars"] = m.vars
	return item
}

func CreateMultiXsend(config map[string]string) *Multixsend {
	return &Multixsend{config["appid"], config["appkey"], config["signType"], "", nil}
}

func (m *Multixsend) SetProject(project string) {
	m.project = project
}

func (m *Multixsend) AddMulti(multi map[string]interface{}) {
	m.multi = append(m.multi, multi)
}

func (m *Multixsend) MultiXsend() string {
	config := make(map[string]string)
	config["appid"] = m.appID
	config["appkey"] = m.appKey
	config["signType"] = m.signType

	request := make(map[string]string)
	request["appid"] = m.appID
	request["project"] = m.project
	if m.signType != "normal" {
		request["sign_type"] = m.signType
		request["timestamp"] = lib.GetTimestamp()
		request["sign_version"] = "2"
	}
	request["signature"] = lib.CreateSignature(request, config)
	//v2 数字签名 multi 不参与计算

	data, err := json.Marshal(m.multi)
	if err == nil {
		request["multi"] = string(data)
	}

	return lib.Post(multixsendURL, request)
}
