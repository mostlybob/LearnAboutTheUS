package main

// playing with the bits found on https://stackoverflow.com/questions/17156371/how-to-get-json-response-in-golang/31129967#31129967

import (
  "fmt"
  "time"
  "net/http"
  "encoding/json"
)

type Foo struct {
    Bar string
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
    r, err := myClient.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

func main() {
    foo1 := new(Foo) // or &Foo{}
    getJson("https://raw.githubusercontent.com/mostlybob/LearnAboutTheUS/master/data.json", foo1)
    fmt.Println(foo1.Bar)

    // alternately:

    //foo2 := Foo{}
    //getJson("http://example.com", &foo2)
    //println(foo2.Bar)
}




//https://play.golang.org/p/pbE4ySX0jP
