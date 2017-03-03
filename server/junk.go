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
	Number string `json:"name"`
	State string `json:"state"`
	Title string `json:"title"`
	Body string `json:"body"`
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
			slackString := "*" + pullRequest.GetTitle() + "*\n" + pullRequest.GetBody() + " _" + pullRequest.GetState() + "_"
			url := "https://hooks.slack.com/services/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"
			sendSlackMessage(url, slackString)
		}
		time.Sleep(time.Second * 5)
	}
}

func sendSlackMessage(url string, text string) string {
	payload := strings.NewReader("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"payload\"\r\n\r\n{\"channel\": \"kill-all-humans\",\"text\": \""+text+"\",\"icon_emoji\": \":ghost:\" }\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--")
	fmt.Println(string(payload))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
	res, err := http.DefaultClient.Do(req) //ToDo: Store if message was already sent in the last x amount of hours
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
	fmt.Println(string(body))

	return
}
