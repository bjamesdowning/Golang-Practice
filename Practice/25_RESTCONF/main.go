package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://172.16.167.150:443/restconf/data/Cisco-IOS-XE-native:native/interface/"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	content := []byte(`{
		"Cisco-IOS-XE-native:Port-channel": [
		  {
			"name": "1", 
			"description": "This is a port-channel interace",
			"delay": 22222, 
			"load-interval": 30, 
			"mtu": 1501
		  }
		]
	 }`)

	client := &http.Client{
		Transport: tr,
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(content))
	if err != nil {
		fmt.Println(err)
	}
	req.SetBasicAuth("cisco", "cisco")
	req.Header.Add("Content-Type", "application/yang-data+json")
	req.Header.Add("Accept", "application/yang-data+json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(string(body))

}
