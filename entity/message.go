package entity

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type SendMessageRequest struct {
	Topic string  `json:"topic"`
	Data  Message `json:"data"`
}

type GrafanaEvent struct {
	EvalMatches []GrafanaEventMatch `json:"evalMatches"`
	ImageUrl    string              `json:"imageUrl"`
	Message     string              `json:"message"`
	RuleId      int                 `json:"ruleId"`
	RuleName    string              `json:"ruleName"`
	RuleUrl     string              `json:"ruleUrl"`
	State       string              `json:"state"`
	Title       string              `json:"title"`
}
type GrafanaEventMatch struct {
	Value  int         `json:"value"`
	Metric string      `json:"metric"`
	Tags   interface{} `json:"tags"`
}

type GoCronEvent struct {
	Status string `json:"status"`
	Name   string `json:"taskName"`
	TaskId int    `json:"taskId"`
	Result string `json:"result"`
}
