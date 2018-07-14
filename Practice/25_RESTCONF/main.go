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
	//Initalize rendering of homepage template
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	///Serve CSS file to client
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//Handlers for URI of APP, calling functions based on site location
	http.HandleFunc("/", index)
	http.HandleFunc("/restconf", restConf)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	//Renders form on root page of site
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func restConf(w http.ResponseWriter, r *http.Request) {
	//Grab the value from the form to put into the logic
	trgt := r.FormValue("dns")

	//Create the URI which uses the RESTCONF entry point
	url := "https://" + trgt + ":9443/restconf/data/ietf-interfaces:interfaces"

	//This is to allow the app to ignore the 'unsafe' trigger from self-signed certs
	ignoreCert := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	//The JSON to interface with RESTCONF to be translated into impacting a YANG data model
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

	//Oour custom client sourced from the APP to the RESTCONF interface on the router
	client := &http.Client{Transport: ignoreCert}

	//Crafting our request body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(content))
	if err != nil {
		fmt.Println(err)
	}

	//Crafting our request headers
	req.SetBasicAuth("root", "D_Vay!_10&")
	req.Header.Add("Content-Type", "application/yang-data+json")
	req.Header.Add("Accept", "application/yang-data+json")

	//Sending our request to the router
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	//Accpeting the return from the router
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//Printing out the routers response body to the page
	fmt.Fprint(w, string(body))
}
