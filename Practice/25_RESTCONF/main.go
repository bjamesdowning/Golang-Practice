package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	///
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", index)
	http.HandleFunc("/login", restConf)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func restConf(w http.ResponseWriter, r *http.Request) {
	url := "https://172.16.167.150:443/restconf/data/Cisco-IOS-XE-native:native/interface/"

	ignoreCert := &http.Transport{
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

	client := &http.Client{Transport: ignoreCert}
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
	fmt.Fprint(w, string(body))
}
