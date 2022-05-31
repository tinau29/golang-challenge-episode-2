package main

import (
	"encoding/json"
	"encoding/xml"
	"episode-2/app"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {

	id := 1
	name := "Test"
	responseJsonXml := "xml"
	data := new(app.Base)
	data.ID = id
	data.Name = strings.ToLower(name)
	responseJsonXml = strings.ToLower(responseJsonXml)
	output, err := json.Marshal(data)
	switch responseJsonXml {
	case "json":
		output, err = json.Marshal(data)
		break
	case "xml":
		output, err = xml.Marshal(data)
		break
	default:
		output, err = json.Marshal(data)
		break
	}

	if nil != err {
		panic(err)
	}

	fmt.Println(string(output))

	product := []app.Product{}

	productOne := new(app.Product)
	categoryOne := new(app.Category)
	categoryOne.SetCategory(2, "Ecommerce")
	productOne.SetProduct(1, "Lakugan", 50000000)
	productOne.Category = categoryOne
	product = append(product, *productOne)

	productTwo := new(app.Product)
	categoryTwo := new(app.Category)
	categoryTwo.SetCategory(1, "Project management")
	productTwo.SetProduct(2, "My Score Card", 60000000)
	productTwo.Category = categoryTwo
	product = append(product, *productTwo)

	http.HandleFunc("/api/v1/products.json", func(res http.ResponseWriter, req *http.Request) {
		output, _ := json.MarshalIndent(product, "", " ")
		res.Header().Set("content-type", "application/json")
		res.WriteHeader((200))
		res.Write(output)
	})

	http.HandleFunc("/api/v1/products.xml", func(res http.ResponseWriter, req *http.Request) {
		output, _ := xml.MarshalIndent(product, "", " ")
		res.Header().Set("content-type", "text/xml")
		res.WriteHeader(200)
		res.Write(output)

	})

	http.HandleFunc("/api/v1/add-products.json", func(res http.ResponseWriter, req *http.Request) {
		payload, _ := ioutil.ReadAll(req.Body)
		pr := new(app.Product)
		json.Unmarshal(payload, pr)

		output, _ := json.MarshalIndent(pr, "", " ")
		res.Header().Set("content-type", "application/json")
		res.WriteHeader(201)
		res.Write(output)

	})

	http.HandleFunc("/api/v1/add-products.xml", func(res http.ResponseWriter, req *http.Request) {
		payload, _ := ioutil.ReadAll(req.Body)
		pr := new(app.Product)

		xml.Unmarshal(payload, pr)

		output, _ := xml.MarshalIndent(pr, "", " ")
		res.Header().Set("content-type", "application/xml")
		res.WriteHeader(201)
		res.Write(output)

	})

	log.Println("Go server running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))

}
