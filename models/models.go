package models



type InternDetails struct {
	SlackUsername string `json:"slackusername"`
	Backend       bool   `json:"backend"`
	Age           int    `json:"age"`
	Bio           string `json:"bio"`
}
type Req_body struct {
	Operation_type string `json:"operation_type"`
	X             int `json:"x"`
	Y            int `json:"y"`
}
type Response_Body struct {
	Username  string    `json:"slackusername"`
	Operation_type string `json:"operation"`
	Result         int       `json:"result"`
	
}
