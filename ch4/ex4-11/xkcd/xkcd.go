// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 110.
//!+

// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package xkcd

import "time"

const IssuesURL = "http://www.xkcd.com"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type Comic struct {
	Month      string `json: "month"`
	Num        int    `json: "num"`
	Link       string `json: "link"`
	Year       string `json: "year"`
	News       string `json: "news"`
	Safe_title string `json: "safe_title"`
	Transcript string `json: "transcript"`
	Alt        string `json: "alt"`
	Img        string `json: "img"`
	Title      string `json: "title"`
	Day        string `json: "day"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

//!-
