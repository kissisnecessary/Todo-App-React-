
package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Hyperledger-TWGC/fabric-gm-plugins/workshop"
)

var priFile = "priv.pem"
var pubFile = "pub.pem"
var path = "./"

type Encrypt struct {
	Msg     string `json:"msg"`
	Encrypt string `json:"encrypt"`
}

func main() {
	// args 1 as key store path
	// args 2 as method
	// args 3 as server endpoint(optional)
	path = os.Args[1]
	if os.Args[2] == "generate" {
		fmt.Println("generate key pair at " + path)
		source, _ := workshop.GenerateSM2Instance(workshop.TJ)
		source.SaveFile(path+priFile, path+pubFile)
	}
	if os.Args[2] == "decrypt" {
		fmt.Println("decrypt")
		httpRequest, _ := http.NewRequest("GET", "http://"+os.Args[3]+"/encrypt", nil)
		httpRequest.Header.Set("Content-Type", "application/json")
		client := http.Client{}
		response, _ := client.Do(httpRequest)
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(body))
		var data Encrypt
		err := json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println(err)
		}
		priv, _ := workshop.LoadFromPriPem(path + priFile)

		test, err := hex.DecodeString(data.Encrypt)