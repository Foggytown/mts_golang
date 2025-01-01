package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hw2/internal/model"
	"math/rand"
	"net/http"
	"time"
)

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		_, err := w.Write([]byte("version: 1.0.0-alpha+fyodor-173a"))
		if err != nil {
			fmt.Printf("Error while writing response: %s", err)
			http.Error(w, "Error while writing response", http.StatusInternalServerError)
			return
		}
	default:
		fmt.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func DecodeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var req model.DecodeReq
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Printf("Error while decoding body %s\n", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		data, err := base64.StdEncoding.DecodeString(req.Base64String)
		if err != nil {
			fmt.Printf("Error while decoding base64: %s\n", err)
			http.Error(w, "Invalid base64 string", http.StatusBadRequest)
			return
		}

		res := model.DecodeResp{DecodeString: string(data)}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			fmt.Printf("Error while writing response: %s\n", err)
			http.Error(w, "Error while writing response", http.StatusInternalServerError)
			return
		}

	default:
		fmt.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func HardOpHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		dur := time.Duration(10 + rand.Intn(11))
		time.Sleep(dur * time.Second)

		if rand.Intn(2) == 0 {
			return
		} else {
			code := 500 + rand.Intn(12)
			http.Error(w, "Random 500 code", code)
			return
		}

	default:
		fmt.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
