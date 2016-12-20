package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	fmt.Println(hello())
	fmt.Println(testJson())
}

func hello() string {
	return "Hello, playground"
}

func testJson() string {
	resp, err := http.Get("https://raw.githubusercontent.com/mostlybob/LearnAboutTheUS/master/data.json")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if (err != nil){
		fmt.Println(err)
	}
	return string(body)

}
