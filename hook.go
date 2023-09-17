package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"
)

type HookRequest struct {
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
	Hash      string `json:"hash"`
}

func generateHash(secret string, timestamp int64, message string) string {
	// Generate hash logic here
	// For example, use sha256
	data := fmt.Sprintf("%s%d%s", secret, timestamp, message)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func handleHook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request")
	// Parse JSON from request body
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println("Request body:", string(body))
	var hookRequest HookRequest
	json.Unmarshal(body, &hookRequest)
	fmt.Println("Parsed JSON:", hookRequest)

	// Validation
	secret := "your_secret_here"
	expectedHash := generateHash(secret, hookRequest.Timestamp, hookRequest.Message)

	if expectedHash == hookRequest.Hash && time.Now().Unix()-hookRequest.Timestamp < 60 {
		// Execute the hook
		exec.Command("./hook.sh").Run()
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func main() {
	http.HandleFunc("/hook", handleHook)
	http.ListenAndServe(":1113", nil)
}
