package aes

import (
	"testing"
)

var (
	testEncode = []string{"", "", "hello world", "123", "测试", "~!@#$%^&*()_+:,./|}{][*-+"}
	testDecode = []string{
		"",
		"E5XWKFZkFGrwiNDajYPNKEnDbr4YsUwo8T5pwpRVVAI",
		"vkE1fw1DusD9o_Rc_g3THm0Y6eo7dsMtm9UhhiZSoCI",
		"eeKdAcLPgSdIwi8Mbg48UV7zdH8FCUrtl7tkI7VCSG0",
		"aqxhq36jLzKxTRJFcxoyga_dOwgvwuc1nC0HyjcNF-E",
		"4reZrTn-hNeukDUZpFD3qQ1mE59JJthp5sXB4Wn-aNJWeSYNxvXeHyNU2ZRVlwax",
	}
)

func TestAesEncrypt(t *testing.T) {
	t.Parallel()
	for _, v := range testEncode {
		str, err := Encrypt(v)
		if str == "" {
			t.Errorf("AesDecrypt Empty: %s", v)
		}
		if err != nil {
			t.Error(err.Error())
		}
		t.Logf("AesEncrypt ok: %s => %s", v, str)
	}
}

func TestAesDecrypt(t *testing.T) {
	t.Parallel()
	for k, v := range testDecode {
		str, err := Decrypt(v)
		if err != nil {
			t.Error(err.Error())
		}
		if result := testEncode[k]; str != testEncode[k] {
			t.Errorf("AesDecrypt fail: %s => %s ,not %s", v, str, result)
		}
		t.Logf("AesDecrypt ok: %s => %s", v, str)
	}
}
