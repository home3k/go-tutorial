package github

import (
	"time"
	"net/url"
	"strings"
	"net/http"
	"fmt"
	"encoding/json"
)

const IssueSearchUrl  = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json: "total_count"`
	Items []*Issue
}

type Issue struct {
	Number int
	HtmlUrl string `json: "html_url"`
	Title string
	State string
	User *User
	CreateAt time.Time `json:"create_at"`
	Body string
}

type User struct {
	Login string
	HtmlUrl string `json:"html_url"`
}

func SearchIssue(items []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(items, " "))
	resp, err := http.Get(IssueSearchUrl + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err!=nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
