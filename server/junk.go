package main

import (
	"fmt"
	"context"
	"golang.org/x/oauth2"
	"github.com/google/go-github/github"
	"time"
	"net/http"
	"strings"
	"io/ioutil"
)

var (
	githubToken = "f39dc652fef49c0be580a614ed64e74411ad638c"
	client *github.Client
)

type PullRequestResponse struct {
	Number string `fmt:"name"`
	State string `fmt:"state"`
	Title string `fmt:"title"`
	Body string `fmt:"body"`
}

func main() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client = github.NewClient(tc)

	ctx := context.Background()

	for {
		pullRequests, _, err := client.PullRequests.List(ctx, "tsergium", "reactjsGolangRethinkdb", nil)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, pullRequest := range pullRequests {
			fmt.Println(pullRequest.GetTitle())
			fmt.Println(pullRequest.GetBody())
			fmt.Println(pullRequest.GetState())
			fmt.Println(pullRequest.GetCreatedAt())
			fmt.Println("========================")

			payload := strings.NewReader("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"payload\"\r\n\r\n{\"channel\": \"kill-all-humans\",\"text\": \"*"+pullRequest.GetTitle()+"*\\n"+pullRequest.GetBody()+" | _"+pullRequest.GetState()+"_\",\"icon_emoji\": \":ghost:\" }\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--")
			url := "https://hooks.slack.com/services/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"
			req, _ := http.NewRequest("POST", url, payload)
			req.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")

			res, _ := http.DefaultClient.Do(req)

			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)

			fmt.Println(res)
			fmt.Println(string(body))
		}
		time.Sleep(time.Second * 5)
	}
}
