package response

type Issue struct {
	Key    string `json:"key"`
	Fields struct {
		IssueType IssueType `json:"issuetype"`
		Project   Project   `json:"project"`
		Priority  Priority  `json:"priority"`
		Assignee  User      `json:"assignee"`
		Status    Status    `json:"status"`
		Summary   string    `json:"summary"`
		Reporter  User      `json:"reporter"`
	} `json:"fields"`
}
