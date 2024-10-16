package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"hw2/internal/model"
	"io"
	"net/http"
	"time"
)

var Addr = "http://localhost:8080/"

func VersionRequest() {
	request, err := http.NewRequest("GET", Addr+"version", nil)
	if err != nil {
		fmt.Printf("Error in creating request: %s", err)
		return
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("Error in doing request: %s", err)
		return
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Printf("Error while closing body of response: %s", err)
			return
		}
	}()
	respBody, _ := io.ReadAll(response.Body)
	fmt.Println(string(respBody))
}

func DecodeBase64Request(base64String string) {
	req := model.DecodeReq{Base64String: base64String}
	jsonBody, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("Error in creating json: %s", err)
		return
	}
	request, err := http.NewRequest("POST", Addr+"decode", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Printf("Error in creating request: %s", err)
		return
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("Error in doing request: %s", err)
		return
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Printf("Error while closing body of response: %s", err)
			return
		}
	}()
	body, _ := io.ReadAll(response.Body)
	var resp model.DecodeResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Printf("Error while parsing json: %s", err)
		return
	}
	fmt.Println(resp.DecodeString)
}

func TimeoutRequest() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", Addr+"hard-op", nil)
	if err != nil {
		fmt.Printf("Error in creating request: %s", err)
		return
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		if ctx.Err() != nil {
			fmt.Println("false")
		} else {
			fmt.Printf("Error in doing request: %s", err)
		}
		return
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Printf("Error while closing body of response: %s", err)
			return
		}
	}()
	fmt.Printf("true, %d", response.StatusCode)

}

func main() {
	VersionRequest()
	DecodeBase64Request("U29tZXRpbWVzIGkgZHJlYW0gYWJvdXQgY2hlZXNl")
	TimeoutRequest()
}
