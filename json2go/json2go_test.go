package json2go

import (
	"fmt"
	"log"
	"testing"
)

func Test_json2struct(t *testing.T) {
	test, err := New([]byte(`[[{"id":123},[{"test":false,"hello":"abc","msg":{}}]]]`), "")
	if err != nil {
		log.Fatalln(err.Error())
	}
	if bytes, err := test.WriteGo(); err != nil {
		log.Fatalln(err.Error())
	} else {
		fmt.Println(string(bytes))
	}
	test, err = New([]byte(`[{"abc":123},{"a":[9,8,7,6,5,4,3,2,1.1]}]`), "")
	if err != nil {
		log.Fatalln(err.Error())
	}
	if bytes, err := test.WriteGo(); err != nil {
		log.Fatalln(err.Error())
	} else {
		fmt.Println(string(bytes))
	}
	test, err = New([]byte(`[[[[[[[]]]]]]]`), "")
	if err != nil {
		log.Fatalln(err.Error())
	}
	if bytes, err := test.WriteGo(); err != nil {
		log.Fatalln(err.Error())
	} else {
		fmt.Println(string(bytes))
	}
	test, err = New([]byte(`[{},[],[{"a":"b"}],{"c":"d"},{}]`), "")
	if err != nil {
		log.Fatalln(err.Error())
	}
	if bytes, err := test.WriteGo(); err != nil {
		log.Fatalln(err.Error())
	} else {
		fmt.Println(string(bytes))
	}
}
