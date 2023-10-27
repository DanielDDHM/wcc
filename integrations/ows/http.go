package ows

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/DanielDDHM/world-coin-converter/config"
)

func OwsRequest(method string, path string, body string) (map[string]interface{}, error) {

	req, err := http.NewRequest(method, config.GetConfig().OwsUrlBase+path, bytes.NewBuffer([]byte(body)))

	if err != nil {
		return nil, err
	}

	nonce := strconv.Itoa(int(time.Now().UnixNano()))

	fmt.Println("Body", body)
	signature := SignRequest("/v1/"+path, nonce, body)
	req.Header.Add("X-API-Key", config.GetConfig().OwsKey)
	req.Header.Add("X-Nonce", nonce)
	req.Header.Add("X-Signature", signature)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	fmt.Println("Response => ", resp)
	fmt.Println("Request => ", req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(string(responseBody)), &jsonMap)

	return jsonMap, nil
}
