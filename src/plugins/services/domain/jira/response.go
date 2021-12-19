package jira

type Search struct {
	MaxResults int     `json:"maxResults"`
	Total      int     `json:"total"`
	Issues     []Issue `json:"issues"`
}

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

type IssueType struct {
	Name string `json:"name"`
}

type Project struct {
	Name string `json:"name"`
}

type Priority struct {
	Name string `json:"name"`
}

type User struct {
	DisplayName string `json:"displayName"`
}

type Status struct {
	Name string `json:"name"`
}
