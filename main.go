package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}
	jdcookie := os.Getenv("JDCOOKIE")
	req, _ := http.NewRequest("GET", "https://api.m.jd.com/client.action?functionId=signBeanAct&appid=ld&client=android", nil)
	req.Header.Set("Cookie", jdcookie)
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

}
