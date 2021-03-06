package urlscan_test

import (
	"context"
	"fmt"
	"log"

	"github.com/vertoforce/urlscan-go/urlscan"
)

func ExampleClient_Submit() {
	client := urlscan.NewClient("YOUR-API-KEY")
	task, err := client.Submit(context.Background(), urlscan.SubmitArguments{URL: "https://golang.org"})
	if err != nil {
		log.Fatal(err)
	}

	err = task.WaitForReport(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, cookie := range task.Result.Data.Cookies {
		fmt.Printf("Cookie: %s = %s\n", cookie.Name, cookie.Value)
	}
}

func ExampleClient_Search() {
	resp, err := urlscan.Search(context.Background(), urlscan.SearchArguments{
		Query:  "ip:1.2.3.x",
		Size:   1,
		Offset: 0,
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, result := range resp.Results {
		fmt.Printf("Related URL: %s\n", result.Page.URL)
	}
}
