package main

//测试用客户端
import (
	"bytes"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type hust struct {
	s string
}

func main() {
	test := []string{"test1", "test2", "test3", "test4", "test5"}
	server := "http://127.0.0.1:9090/?action=log"
	var _test []byte
	for _, v := range test {
		_test = append(_test, []byte(v)...)
	}

	body := bytes.NewReader(_test)
	resp, err := http.Post(server, "", body)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	txt, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(txt))

}
