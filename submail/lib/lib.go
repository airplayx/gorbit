package lib

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const Server = "https://api.mysubmail.com"

func Get(requesturl string) string {
	u, _ := url.Parse(requesturl)
	retstr, err := http.Get(u.String())
	if err != nil {
		return err.Error()
	}
	result, err := ioutil.ReadAll(retstr.Body)
	retstr.Body.Close()
	if err != nil {
		return err.Error()
	}
	return string(result)
}

func Post(requesturl string, postdata map[string]string) string {
	var r http.Request
	r.ParseForm()

	for key, val := range postdata {
		r.Form.Add(key, val)
	}

	body := strings.NewReader(r.Form.Encode())

	//打印请求体
	//fmt.Println("request:", r.Form.Encode())

	retstr, err := http.Post(requesturl, "application/x-www-form-urlencoded;charset=utf-8", body)

	if err != nil {
		return err.Error()
	}
	result, err := ioutil.ReadAll(retstr.Body)
	retstr.Body.Close()
	if err != nil {
		return err.Error()
	}
	return string(result)
}

func MultipartPost(requesturl string, postdata map[string]string) string {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, val := range postdata {
		if key == "attachments" {
			attachments := strings.Split(val, ",")
			if len(attachments) > 0 {
				for _, filename := range attachments {
					//fmt.Println("file:", filename)
					file, err := os.Open(filename)
					if err != nil {
						return err.Error()
					}
					defer file.Close()
					part, err := writer.CreateFormFile("attachments[]", filepath.Base(filename))
					if err != nil {
						return err.Error()
					}
					_, err = io.Copy(part, file)
				}
			}
		} else {
			_ = writer.WriteField(key, val)
		}
	}
	err := writer.Close()
	if err != nil {
		return err.Error()
	}
	contentType := writer.FormDataContentType()
	writer.Close()
	//打印请求体
	//fmt.Println("request:", string(body.Bytes()))

	resp, err := http.Post(requesturl, contentType, body)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return string(result)
}

func openfile(filename string) *os.File {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	return file
}

func GetTimestamp() string {
	resp := Get(Server + "/service/timestamp")
	var dict map[string]interface{}
	err := json.Unmarshal([]byte(resp), &dict)
	if err != nil {
		return err.Error()
	}
	return strconv.Itoa(int(dict["timestamp"].(float64)))
}

func CreateSignature(request map[string]string, config map[string]string) string {
	appkey := config["appkey"]
	appid := config["appid"]
	signtype := config["signType"]

	keys := make([]string, 0, 32)
	for key, _ := range request {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	str_list := make([]string, 0, 32)
	for _, key := range keys {
		str_list = append(str_list, key+"="+request[key])
	}
	sigstr := strings.Join(str_list, "&")
	sigstr = appid + appkey + sigstr + appid + appkey
	if signtype == "md5" {
		mymd5 := md5.New()
		io.WriteString(mymd5, sigstr)
		return hex.EncodeToString(mymd5.Sum(nil))
	} else if signtype == "sha1" {
		mysha1 := sha1.New()
		io.WriteString(mysha1, sigstr)
		return hex.EncodeToString(mysha1.Sum(nil))
	} else {
		return appkey
	}
}
