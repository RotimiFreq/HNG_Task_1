package models

type InternDetails struct {
	SlackUsername string `json:"slackusername"`
	Backend       bool   `json:"backend"`
	Age           int    `json:"age"`
	Bio           string `json:"bio"`
}
