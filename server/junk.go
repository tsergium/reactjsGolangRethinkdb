package main

import (
	"github.com/octokit/go-octokit/octokit"
	"fmt"
)

func main() {
	client := octokit.NewClient(nil)
	//octokit.pull_requests('rails/rails', :state => 'closed')

	url, err := octokit.UserURL.Expand(octokit.M{"user": "tsergium"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, result := client.Users(url).One()
	if result.HasError() {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(user.ReposURL)
}