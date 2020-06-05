package mail

import (
	"encoding/json"
	"fmt"
	"github.com/airplayx/gorbit/submail/lib"
	"strings"
)

type mailsend struct {
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
	html         string
	text         string
	vars         map[string]string
	links        map[string]string
	headers      map[string]string
	asynchronous string
	attachments  []string
	tag          string
}

const sendURL = lib.Server + "/mail/send"

func CreateSend(config map[string]string) *mailsend {
	mail := new(mailsend)
	mail.appid = config["appid"]
	mail.appkey = config["appkey"]
	mail.signType = config["signType"]
	mail.vars = make(map[string]string)
	mail.links = make(map[string]string)
	mail.headers = make(map[string]string)
	return mail
}

func (this *mailsend) AddTo(address string, name string) {
	item := make(map[string]string)
	item["address"] = address
	item["name"] = name
	this.to = append(this.to, item)
}

func (this *mailsend) SetSender(address string, name string) {
	this.from = address
	this.fromName = name
}

func (this *mailsend) SetReply(address string) {
	this.reply = address
}

func (this *mailsend) AddCc(address string) {
	this.cc = append(this.cc, address)
}

func (this *mailsend) AddBcc(address string) {
	this.bcc = append(this.bcc, address)
}

func (this *mailsend) SetSubject(subject string) {
	this.subject = subject
}

func (this *mailsend) SetHtml(html string) {
	this.html = html
}

func (this *mailsend) SetText(text string) {
	this.text = text
}

func (this *mailsend) AddVar(key string, val string) {
	this.vars[key] = val
}

func (this *mailsend) AddLink(key string, val string) {
	this.links[key] = val
}

func (this *mailsend) AddHeaders(key string, val string) {
	this.headers[key] = val
}

func (this *mailsend) SetAsynchronous(status bool) {
	if status {
		this.asynchronous = "true"
	} else {
		this.asynchronous = "false"
	}
}

func (this *mailsend) AddAttachments(file string) {
	this.attachments = append(this.attachments, file)
}

func (this *mailsend) SetTag(tag string) {
	this.tag = tag
}

func (this *mailsend) Send() string {
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

	//V2 版数字签名 html / text / vars / links / headers / attachments 不参与数字签名计算
	if this.html != "" {
		request["html"] = this.html
	}
	if this.text != "" {
		request["text"] = this.text
	}

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
	if len(this.attachments) > 0 {
		request["attachments"] = strings.Join(this.attachments, ",")
	}
	return lib.MultipartPost(sendURL, request)
}
