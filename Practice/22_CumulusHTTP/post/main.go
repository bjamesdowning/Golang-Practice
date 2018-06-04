package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Basic grab configs from RestAPI on Cumulus Linux
func main() {
	//Ignores invalid cert from Linux
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//Url of host
	url := "https://192.168.198.128:8080/nclu/v1/rpc/"
	//Json to be sent to the host, commands to nclu. This is what ansible also uses.
	jsonStr := []byte(`{"cmd": "show counters"}`)
	//Altering the client to take use non default transport and ignore invalid cert
	client := &http.Client{Transport: tr}
	//creating a new request
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	//Placing default credentials in headers for request
	req.SetBasicAuth("cumulus", "CumulusLinux!")
	//Making call
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	//grabbing body from *response and printing
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
