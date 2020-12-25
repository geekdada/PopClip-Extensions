package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/geekdada/flomo-cli/client"
)

func main() {
	api := os.Getenv("POPCLIP_OPTION_API")
	tag := os.Getenv("POPCLIP_OPTION_TAG")
	rawContent := os.Getenv("POPCLIP_TEXT")
	browserTitle := os.Getenv("POPCLIP_BROWSER_TITLE")
	browserURL := os.Getenv("POPCLIP_BROWSER_URL")
	content := strings.TrimSpace(rawContent)

	if api == "" || content == "" {
		log.Fatalf("缺少必要参数")
	}

	if browserTitle != "" && browserURL != "" {
		content += fmt.Sprintf("\n\n标题：%s", browserTitle)
		content += fmt.Sprintf("\nURL：%s", browserURL)
	}

	memo := client.Memo{
		Api:     api,
		Content: content,
		Tag:     tag,
	}

	responseMessage, err := memo.Submit(true)

	if err != nil {
		switch err.(type) {
		case *client.ResponseError:
			re, _ := err.(*client.ResponseError)

			log.Println(err)

			if re.StatusCode >= 400 && re.StatusCode < 500 {
				os.Exit(2)
			} else {
				os.Exit(1)
			}
		default:
			log.Println(err)
			os.Exit(1)
		}
	}

	log.Println(responseMessage)
	os.Exit(0)
}
