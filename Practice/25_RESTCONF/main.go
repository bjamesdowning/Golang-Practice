package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://ios-xe-mgmt.cisco.com:9443/restconf/data/ietf-interfaces:interfaces/interface=GigabitEthernet1?fields=name;description"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.SetBasicAuth("root", "D_Vay!_10&")
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
