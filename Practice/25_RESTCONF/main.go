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
	http.HandleFunc("/restconf", restConf)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func restConf(w http.ResponseWriter, r *http.Request) {
	trgt := r.FormValue("dns")
	url := "https://" + trgt + ":9443/restconf/data/ietf-interfaces:interfaces"

	ignoreCert := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	content := []byte(`{
		"ietf-interfaces:interface": {
			"name": "Loopback102",
			"description": "Configured by RESTCONF",
			"type": "iana-if-type:softwareLoopback",
			"enabled": true,
			"ietf-ip:ipv4": {
				"address": [
					{
						"ip": "101.101.102.103",
						"netmask": "255.255.255.255"
					}
				]
			}
		}
	}`)

	client := &http.Client{Transport: ignoreCert}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(content))
	if err != nil {
		fmt.Println(err)
	}
	req.SetBasicAuth("root", "D_Vay!_10&")
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
