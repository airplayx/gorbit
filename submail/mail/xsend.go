package mail

import (
	"encoding/json"
	"fmt"
	"github.com/airplayx/gorbit/submail/lib"
	"strings"
)

type mailxsend struct {
	appid        string
	appkey       string
	signType     string
	to           []map[string]string
	from         string
	fromName     string
	reply        string
	cc           []string
	bcc          []string
	subject      string
	project      string
	vars         map[string]string
	links        map[string]string
	headers      map[string]string
	asynchronous string
	tag          string
}

const xsendURL = lib.Server + "/mail/xsend"

func CreateXsend(config map[string]string) *mailxsend {
	return &mailxsend{config["appid"], config["appkey"], config["signType"], nil, "", "", "", nil, nil, "", "", make(map[string]string), make(map[string]string), make(map[string]string), "", ""}
}

func (this *mailxsend) AddTo(address string, name string) {
	item := make(map[string]string)
	item["address"] = address
	item["name"] = name
	this.to = append(this.to, item)
}

func (this *mailxsend) SetSender(address, name string) {
	this.from = address
}

func (this *mailxsend) SetReply(address string) {
	this.reply = address
}

func (this *mailxsend) AddCc(address string) {
	this.cc = append(this.cc, address)
}

func (this *mailxsend) AddBcc(address string) {
	this.bcc = append(this.bcc, address)
}

func (this *mailxsend) SetSubject(subject string) {
	this.subject = subject
}

func (this *mailxsend) SetProject(project string) {
	this.project = project
}

func (this *mailxsend) AddVar(key string, val string) {
	this.vars[key] = val
}

func (this *mailxsend) AddLink(key string, val string) {
	this.links[key] = val
}

func (this *mailxsend) AddHeaders(key string, val string) {
	this.headers[key] = val
}

func (this *mailxsend) SetAsynchronous(status bool) {
	if status {
		this.asynchronous = "true"
	} else {
		this.asynchronous = "false"
	}
}

func (this *mailxsend) SetTag(tag string) {
	this.tag = tag
}

func (this *mailxsend) Xsend() string {
	config := make(map[string]string)
	config["appid"] = this.appid
	config["appkey"] = this.appkey
	config["signType"] = this.signType

	request := make(map[string]string)
	request["appid"] = this.appid
	if len(this.to) > 0 {
		to_list := make([]string, 0, 32)
		for _, item := range this.to {
			to_list = append(to_list, fmt.Sprintf("%s<%s>", item["name"], item["address"]))
		}
		request["to"] = strings.Join(to_list, ",")
	}
	if this.from != "" {
		request["from"] = this.from
	}
	if this.fromName != "" {
		request["from_name"] = this.fromName
	}
	if this.reply != "" {
		request["reply"] = this.reply
	}
	if len(this.cc) > 0 {
		request["cc"] = strings.Join(this.cc, ",")
	}
	if len(this.bcc) > 0 {
		request["bcc"] = strings.Join(this.bcc, ",")
	}
	if this.subject != "" {
		request["subject"] = this.subject
	}
	request["project"] = this.project
	if this.asynchronous != "" {
		request["asynchronous"] = this.asynchronous
	}
	if this.tag != "" {
		request["tag"] = this.tag
	}
	if this.signType != "normal" {
		request["sign_type"] = this.signType
		request["timestamp"] = lib.GetTimestamp()
		request["sign_version"] = "2"
	}
	request["signature"] = lib.CreateSignature(request, config)
	if len(this.vars) > 0 {
		data, err := json.Marshal(this.vars)
		if err == nil {
			request["vars"] = string(data)
		}
	}

	if len(this.links) > 0 {
		data, err := json.Marshal(this.links)
		if err == nil {
			request["links"] = string(data)
		}
	}

	if len(this.headers) > 0 {
		data, err := json.Marshal(this.headers)
		if err == nil {
			request["headers"] = string(data)
		}
	}
	return lib.Post(xsendURL, request)
}
