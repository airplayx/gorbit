package wxpay

import (
	"encoding/xml"
	"github.com/clbanning/mxj"
	"testing"
)

type NotifyResult struct {
	AppId    string `xml:"appid"`
	MchId    string `xml:"mch_id"`
	NonceStr string `xml:"nonce_str"`
	Sign     string `xml:"sign"`

	ReturnCode    string `xml:"return_code"`
	ReturnMsg     string `xml:"return_msg"`
	ResultCode    string `xml:"result_code"`
	OpenId        string `xml:"openid"`
	IsSubscribe   string `xml:"is_subscribe"`
	TradeType     string `xml:"trade_type"`
	BankType      string `xml:"bank_type"`
	TotalFee      int    `xml:"total_fee"`
	FeeType       string `xml:"fee_type"`
	CashFee       int    `xml:"cash_fee"`
	CashFeeType   string `xml:"cash_fee_type"`
	TransactionId string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	Attach        string `xml:"attach"`
	TimeEnd       string `xml:"time_end"`
}

func TestMakeWxPaySign(t *testing.T) {
	t.Parallel()
	body := []byte(`<xml>
  <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
  <attach><![CDATA[支付测试]]></attach>
  <bank_type><![CDATA[CFT]]></bank_type>
  <fee_type><![CDATA[CNY]]></fee_type>
  <is_subscribe><![CDATA[Y]]></is_subscribe>
  <mch_id><![CDATA[10000100]]></mch_id>
  <nonce_str><![CDATA[5d2b6c2a8db53831f7eda20af46e531c]]></nonce_str>
  <openid><![CDATA[oUpF8uMEb4qRXf22hE3X68TekukE]]></openid>
  <out_trade_no><![CDATA[1409811653]]></out_trade_no>
  <result_code><![CDATA[SUCCESS]]></result_code>
  <return_code><![CDATA[SUCCESS]]></return_code>
  <sign><![CDATA[B552ED6B279343CB493C5DD0D78AB241]]></sign>
  <time_end><![CDATA[20140903131540]]></time_end>
  <total_fee>1</total_fee>
  <coupon_fee><![CDATA[10]]></coupon_fee>
  <coupon_count><![CDATA[1]]></coupon_count>
  <coupon_type><![CDATA[CASH]]></coupon_type>
  <coupon_id><![CDATA[10000]]></coupon_id>
  <coupon_fee><![CDATA[100]]></coupon_fee>
  <trade_type><![CDATA[JSAPI]]></trade_type>
  <transaction_id><![CDATA[1004400740201409030005092168]]></transaction_id>
</xml>`)
	var notifyReq NotifyResult
	err := xml.Unmarshal(body, &notifyReq)
	if err != nil {
		t.Error(err)
	}
	mv, err := mxj.NewMapXml(body)
	if err != nil {
		t.Error(err)
	}
	notify := mv["xml"].(map[string]interface{})
	delete(notify, "sign")
	str := make(map[string]string)
	for key, value := range notify {
		switch value := value.(type) {
		case string:
			str[key] = value
		}
	}
	t.Log(PaySign(str, "Wohexiaoyujintianyaogaimimabixu3"))
}

func TestSetNonceStr(t *testing.T) {
	t.Parallel()
	t.Log(NonceStr(0))
	t.Log(NonceStr(20))
}
