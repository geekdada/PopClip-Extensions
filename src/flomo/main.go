package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gojektech/heimdall/v6/httpclient"
	"github.com/pkg/errors"
)

func main() {
	api := os.Getenv("POPCLIP_OPTION_API")
	tag := os.Getenv("POPCLIP_OPTION_TAG")
	rawContent := os.Getenv("POPCLIP_TEXT")
	browserTitle := os.Getenv("POPCLIP_BROWSER_TITLE")
	browserURL := os.Getenv("POPCLIP_BROWSER_URL")
	content := strings.TrimSpace(rawContent)

	type payload struct {
		Content string `json:"content"`
	}

	if api == "" || content == "" {
		log.Fatalf("缺少必要参数")
	}

	if browserTitle != "" && browserURL != "" {
		content += fmt.Sprintf("\n\n标题：%s", browserTitle)
		content += fmt.Sprintf("\nURL：%s", browserURL)
	}

	if tag != "" {
		content += fmt.Sprintf("\n\n#%s", tag)
	}

	timeout := 3000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	payloadJSON, _ := json.Marshal(payload{
		content,
	})
	body := ioutil.NopCloser(strings.NewReader(string(payloadJSON)))
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")

	log.Printf("Raw content: %s", content)
	log.Printf("Payload JSON: %s", payloadJSON)

	response, err := client.Post(api, body, headers)

	if err != nil {
		errors.Wrap(err, "failed to make a request to server")
		log.Println(err)
		os.Exit(1)
		return
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		errors.Wrap(err, "failed to read response body")
		log.Println(err)
		os.Exit(1)
		return
	}

	var responseData map[string]interface{}

	if err := json.Unmarshal(responseBody, &responseData); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	statusCode := response.StatusCode
	message := responseData["message"].(string)

	if statusCode >= 200 && statusCode < 400 {
		log.Println(message)
		os.Exit(0)
	} else if statusCode >= 400 && statusCode < 500 {
		os.Exit(2)
	} else {
		os.Exit(1)
	}
}
